package common

import (
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
	"github.com/walnuts1018/2024-golang-linebot/helper/config"
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
						// 受け取ったメッセージから授業を正しい形式で取り出す
						subject, err := ParseFromString(message.Text)

						// エラーが発生した場合はエラーメッセージを返信する
						if err != nil {
							slog.Error(fmt.Sprintf("Failed to parse subject: %v", err))
							if _, err := api.ReplyMessage(
								&messaging_api.ReplyMessageRequest{
									ReplyToken: e.ReplyToken,
									Messages: []messaging_api.MessageInterface{
										messaging_api.TextMessage{
											Text: fmt.Sprintf("授業の抽出に失敗しました: %v", err),
										},
									},
								},
							); err != nil {
								slog.Error(fmt.Sprintf("Failed to reply message: %v", err))
								c.JSON(500, gin.H{"error": "Failed to reply message"})
							}

							return
						}

						if _, err := api.ReplyMessage(
							&messaging_api.ReplyMessageRequest{
								ReplyToken: e.ReplyToken,
								Messages: []messaging_api.MessageInterface{
									messaging_api.TextMessage{
										Text: subject.String(),
									},
								},
							},
						); err != nil {
							slog.Error(fmt.Sprintf("Failed to reply message: %v", err))
							c.JSON(500, gin.H{"error": "Failed to reply message"})
							return
						}

						slog.Info(fmt.Sprintf("Replied message: %v", message.Text))
					}
				}

			}

			c.JSON(200, gin.H{"message": "OK"})
		})
	}
	return r, nil
}
