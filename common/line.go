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

	// dbClient, err := NewDBClient(cfg)
	// if err != nil {
	// 	return nil, err
	// }

	dbClient, err := NewFileDB("./tmp/db.json")
	if err != nil {
		return nil, err
	}

	proxy := r.Group("/proxy/8080")
	{
		proxy.GET("/healthz", func(c *gin.Context) {
			c.String(200, "OK")
		})
		proxy.POST("/callback", func(c *gin.Context) {
			req, err := webhook.ParseRequest(cfg.LineChannelSecret, c.Request)
			if err != nil {
				slog.Error(fmt.Sprintf("Failed to parse request: %v", err))
				c.JSON(500, gin.H{"error": "Failed to parse request"})
				return
			}

			for _, event := range req.Events {
				slog.Info(fmt.Sprintf("Event: %v", event))

				switch e := event.(type) {
				case webhook.MessageEvent:
					switch message := e.Message.(type) {
					case webhook.TextMessageContent:
						if message.Text == "カレンダー" {
							showCalendar(c, api, e, dbClient)
						} else {
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

func addSubject(c *gin.Context, message webhook.TextMessageContent, api *messaging_api.MessagingApiAPI, e webhook.MessageEvent, dbClient Storage) {
	// 受け取ったメッセージから授業を正しい形式で取り出す
	subject, err := ParseFromString(message.Text)

	// エラーが発生した場合はエラーメッセージを返信する
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to parse subject: %v", err))
		sendMessage(c, fmt.Sprintf("授業の取得に失敗しました: %v", err), api, e)
		// 終了
		return
	}

	// エラーが発生せず、授業が正しく取り出せた場合は、保存しておく
	if err := dbClient.AddSubject(subject); err != nil {
		slog.Error(fmt.Sprintf("Failed to parse subject: %v", err))
		sendMessage(c, fmt.Sprintf("授業の保存に失敗しました: %v", err), api, e)
		// 終了
		return
	}

	// エラーが発生せず、授業が正しく取り出せ、保存できた場合は、保存した授業を返信する
	sendMessage(c, fmt.Sprintf("授業を保存しました:\n %v", subject), api, e)

	slog.Info(fmt.Sprintf("Replied message: %v", message.Text))
}

func showCalendar(c *gin.Context, api *messaging_api.MessagingApiAPI, e webhook.MessageEvent, dbClient Storage) {

	subjects, err := dbClient.GetSubjects()
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to get subjects: %v", err))
		sendMessage(c, fmt.Sprintf("授業の取得に失敗しました: %v", err), api, e)
		return
	}

	flexMessage, err := CreateCalenderJson(subjects)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to create flex message: %v", err))
		sendMessage(c, fmt.Sprintf("カレンダーの作成に失敗しました: %v", err), api, e)
		return
	}

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

func sendMessage(c *gin.Context, text string, api *messaging_api.MessagingApiAPI, e webhook.MessageEvent) {
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
