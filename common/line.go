// このファイルには、LINE Messaging API との通信を行う部分が記述されています。
package common

import (
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
	"github.com/walnuts1018/2024-golang-linebot/common/config"
)

func NewRouter(cfg config.Config) (*gin.Engine, error) {
	if cfg.LogLevel != slog.LevelDebug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	api, err := messaging_api.NewMessagingApiAPI(cfg.LineChannelToken)
	if err != nil {
		return nil, err
	}

	// 今回は、データをファイルに保存します。
	// ./tmp/db.jsonに保存されていきます。
	// dbClient.AddSubject()でデータを追加し、dbClient.GetSubjects()でデータを取得します。
	dbClient, err := NewFileDB("./tmp/db.json")
	if err != nil {
		return nil, fmt.Errorf("failed to create db client: %v", err)
	}

	// https://URL/proxy/8080 にアクセスされたときの処理を記述します。
	proxy := r.Group("/proxy/8080")
	{
		// /proxy/8080/healthz にアクセスされたときに、"OK"と返します。
		// これによって、サーバーが正しく起動できているか確認できます。
		proxy.GET("/healthz", func(c *gin.Context) {
			c.String(200, "OK")
		})

		// /proxy/8080/callback にPOSTリクエストが送られたときの処理を記述します。
		// これは、LINEにメッセージが送られたときに呼び出されます。
		proxy.POST("/callback", func(c *gin.Context) {
			// 送られてきたデータをパースして、メッセージを取り出します。
			// パースとは、データを解析して、プログラムが扱いやすい形に変換することです。
			req, err := webhook.ParseRequest(cfg.LineChannelSecret, c.Request)
			// エラーが発生した場合は、エラーメッセージを返信します。
			if err != nil {
				slog.Error(fmt.Sprintf("Failed to parse request: %v", err))
				c.JSON(500, gin.H{"error": "Failed to parse request"})
				return
			}

			// 送られてきたメッセージには、「イベント」というものが含まれています。
			// イベントとは「○○さんがメッセージを送った」「スタンプを送った」など、LINE上で起こった出来事のことです。
			// for でイベントを一つずつ取り出し、メッセージの内容によって処理を分岐します。
			for _, event := range req.Events {
				slog.Info(fmt.Sprintf("Event: %v", event))

				// switchを使って、イベントの種類によって処理を分岐します。
				switch e := event.(type) {
				// メッセージイベントの場合
				case webhook.MessageEvent:
					switch message := e.Message.(type) {
					// テキストメッセージの場合
					case webhook.TextMessageContent:
						// 「カレンダー」というメッセージが送られた場合は、カレンダーを表示します。
						if message.Text == "カレンダー" {
							showCalendar(c, api, e, dbClient)
						} else {
							// それ以外の場合は、授業を追加します。
							addSubject(c, message, api, e, dbClient)
						}
					}
				}
			}

			c.JSON(200, gin.H{"message": "OK"})
		})
	}
	return r, nil
}

// 「月1 講義名 場所」の形式のメッセージが送られてきた時、授業を登録します。
func addSubject(c *gin.Context, message webhook.TextMessageContent, api *messaging_api.MessagingApiAPI, e webhook.MessageEvent, dbClient Storage) {
	// 受け取ったメッセージから授業を正しい形式で取り出します。
	subject, err := ParseFromString(message.Text)

	// エラーが発生した場合はエラーメッセージを返信する
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to parse subject: %v", err))
		sendTextMessage(c, fmt.Sprintf("授業の取得に失敗しました: %v", err), api, e)
		// 終了
		return
	}

	// メッセージから、ユーザーIDを取得します。
	// ユーザーIDは、LINEのユーザーを識別するためのIDです。
	// ユーザーIDごとに授業を保存しておかないと、いろんな人の授業が混ざってしまいます。
	userID, err := getUserID(e)
	// エラーが発生した場合はエラーメッセージを返信する
	if err != nil {
		slog.Error("Failed to get user id")
		sendTextMessage(c, "ユーザーIDの取得に失敗しました", api, e)
		return
	}

	// エラーが発生せず、授業が正しく取り出せた場合は、授業を保存しておきます。
	if err := dbClient.AddSubject(subject, userID); err != nil {
		slog.Error(fmt.Sprintf("Failed to parse subject: %v", err))
		sendTextMessage(c, fmt.Sprintf("授業の保存に失敗しました: %v", err), api, e)
		// 終了
		return
	}

	// エラーが発生せず、授業が正しく取り出せ、保存できた場合は、保存した授業を返信してあげます。
	sendTextMessage(c, fmt.Sprintf("授業を保存しました:\n %v", subject), api, e)

	slog.Info(fmt.Sprintf("Replied message: %v", message.Text))
}

// 「カレンダー」というメッセージが送られた場合、カレンダーを表示します。
func showCalendar(c *gin.Context, api *messaging_api.MessagingApiAPI, e webhook.MessageEvent, dbClient Storage) {
	// UserIDを取得します。
	userID, err := getUserID(e)
	if err != nil {
		slog.Error("Failed to get user id")
		sendTextMessage(c, "ユーザーIDの取得に失敗しました", api, e)
		return
	}

	// UserIDを元に、保存されている授業を取得します。
	subjects, err := dbClient.GetSubjects(userID)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to get subjects: %v", err))
		sendTextMessage(c, fmt.Sprintf("授業の取得に失敗しました: %v", err), api, e)
		return
	}

	// 取得した授業を元に、時間割の形に整形したメッセージを作成します。
	flexMessage, err := CreateCalenderJson(subjects)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to create flex message: %v", err))
		sendTextMessage(c, fmt.Sprintf("カレンダーの作成に失敗しました: %v", err), api, e)
		return
	}

	// 作成したメッセージを返信します。
	if _, err := api.ReplyMessage(
		&messaging_api.ReplyMessageRequest{
			ReplyToken: e.ReplyToken,
			Messages: []messaging_api.MessageInterface{
				flexMessage,
			},
		},
	); err != nil {
		slog.Error(fmt.Sprintf("Failed to reply message: %v", err))
		c.JSON(500, gin.H{"error": "Failed to reply message"})
		return
	}

	slog.Info("Replied message: カレンダー")
}

// テキストメッセージを返信します。
func sendTextMessage(c *gin.Context, text string, api *messaging_api.MessagingApiAPI, e webhook.MessageEvent) {
	if _, err := api.ReplyMessage(
		&messaging_api.ReplyMessageRequest{
			ReplyToken: e.ReplyToken,
			Messages: []messaging_api.MessageInterface{
				messaging_api.TextMessage{
					Text: text,
				},
			},
		},
	); err != nil {
		slog.Error(fmt.Sprintf("Failed to reply message: %v", err))
		c.JSON(500, gin.H{"error": "Failed to reply message"})
		return
	}

	slog.Info(fmt.Sprintf("Replied message: %v", text))
}

// ユーザーIDを取得します。
func getUserID(e webhook.MessageEvent) (string, error) {
	userSource, ok := e.Source.(webhook.UserSource)
	if !ok {
		return "", fmt.Errorf("failed to get user id")
	}
	return userSource.UserId, nil
}
