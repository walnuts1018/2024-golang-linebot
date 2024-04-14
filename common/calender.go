package common

import (
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

func CreateCalenderJson(subjects []Subject) (messaging_api.FlexMessage, error) {
	const json = `{
		"type": "bubble",
		"body": {
		  "type": "box",
		  "layout": "vertical",
		  "contents": [
			{
			  "type": "box",
			  "layout": "horizontal",
			  "contents": [
				{
				  "type": "text",
				  "text": "　",
				  "flex": 1
				},
				{
				  "type": "text",
				  "text": "月",
				  "flex": 3,
				  "align": "center"
				},
				{
				  "type": "text",
				  "text": "火",
				  "flex": 3,
				  "align": "center"
				},
				{
				  "type": "text",
				  "text": "水",
				  "flex": 3,
				  "align": "center"
				},
				{
				  "type": "text",
				  "text": "木",
				  "flex": 3,
				  "align": "center"
				},
				{
				  "type": "text",
				  "text": "金",
				  "flex": 3,
				  "align": "center"
				},
				{
				  "type": "text",
				  "text": "土",
				  "flex": 3,
				  "align": "center"
				},
				{
				  "type": "text",
				  "text": "日",
				  "flex": 3,
				  "align": "center"
				}
			  ],
			  "cornerRadius": "none",
			  "justifyContent": "center",
			  "alignItems": "center",
			  "spacing": "none"
			},
			{
			  "type": "separator"
			},
			{
			  "type": "box",
			  "layout": "horizontal",
			  "contents": [
				{
				  "type": "text",
				  "text": "1",
				  "flex": 1
				},
				{
				  "type": "text",
				  "text": "パワーエレクトロニクス",
				  "flex": 3,
				  "scaling": true,
				  "wrap": true,
				  "size": "10px"
				},
				{
				  "type": "text",
				  "text": "パワーエレクトロニクス",
				  "flex": 3,
				  "scaling": true,
				  "wrap": true,
				  "size": "10px"
				},
				{
				  "type": "text",
				  "text": "パワーエレクトロニクス",
				  "flex": 3,
				  "scaling": true,
				  "wrap": true,
				  "size": "10px"
				},
				{
				  "type": "text",
				  "text": "パワーエレクトロニクス",
				  "flex": 3,
				  "scaling": true,
				  "wrap": true,
				  "size": "10px"
				},
				{
				  "type": "text",
				  "text": "パワーエレクトロニクス",
				  "flex": 3,
				  "scaling": true,
				  "wrap": true,
				  "size": "10px"
				},
				{
				  "type": "text",
				  "text": "パワーエレクトロニクス",
				  "flex": 3,
				  "scaling": true,
				  "wrap": true,
				  "size": "10px"
				},
				{
				  "type": "text",
				  "text": "パワーエレクトロニクス",
				  "flex": 3,
				  "scaling": true,
				  "wrap": true,
				  "size": "10px"
				}
			  ],
			  "spacing": "none"
			},
			{
			  "type": "box",
			  "layout": "horizontal",
			  "contents": [
				{
				  "type": "text",
				  "text": "2",
				  "flex": 1,
				  "align": "center"
				},
				{
				  "type": "text",
				  "text": "パワーエレクトロニクス",
				  "flex": 3,
				  "scaling": true,
				  "wrap": true,
				  "size": "10px"
				},
				{
				  "type": "text",
				  "text": "パワーエレクトロニクス",
				  "flex": 3,
				  "scaling": true,
				  "wrap": true,
				  "size": "10px"
				},
				{
				  "type": "text",
				  "text": "パワーエレクトロニクス",
				  "flex": 3,
				  "scaling": true,
				  "wrap": true,
				  "size": "10px"
				},
				{
				  "type": "text",
				  "text": "パワーエレクトロニクス",
				  "flex": 3,
				  "scaling": true,
				  "wrap": true,
				  "size": "10px"
				},
				{
				  "type": "text",
				  "text": "パワーエレクトロニクス",
				  "flex": 3,
				  "scaling": true,
				  "wrap": true,
				  "size": "10px"
				},
				{
				  "type": "text",
				  "text": "パワーエレクトロニクス",
				  "flex": 3,
				  "scaling": true,
				  "wrap": true,
				  "size": "10px"
				},
				{
				  "type": "text",
				  "text": "パワーエレクトロニクス",
				  "flex": 3,
				  "scaling": true,
				  "wrap": true,
				  "size": "10px"
				}
			  ],
			  "spacing": "none"
			},
			{
			  "type": "box",
			  "layout": "horizontal",
			  "contents": [
				{
				  "type": "text",
				  "text": "3",
				  "flex": 1,
				  "align": "center"
				},
				{
				  "type": "text",
				  "text": "パワーエレクトロニクス",
				  "flex": 3,
				  "scaling": true,
				  "wrap": true,
				  "size": "10px"
				},
				{
				  "type": "text",
				  "text": "パワーエレクトロニクス",
				  "flex": 3,
				  "scaling": true,
				  "wrap": true,
				  "size": "10px"
				},
				{
				  "type": "text",
				  "text": "パワーエレクトロニクス",
				  "flex": 3,
				  "scaling": true,
				  "wrap": true,
				  "size": "10px"
				},
				{
				  "type": "text",
				  "text": "パワーエレクトロニクス",
				  "flex": 3,
				  "scaling": true,
				  "wrap": true,
				  "size": "10px"
				},
				{
				  "type": "text",
				  "text": "パワーエレクトロニクス",
				  "flex": 3,
				  "scaling": true,
				  "wrap": true,
				  "size": "10px"
				},
				{
				  "type": "text",
				  "text": "パワーエレクトロニクス",
				  "flex": 3,
				  "scaling": true,
				  "wrap": true,
				  "size": "10px"
				},
				{
				  "type": "text",
				  "text": "パワーエレクトロニクス",
				  "flex": 3,
				  "scaling": true,
				  "wrap": true,
				  "size": "10px"
				}
			  ],
			  "spacing": "none"
			},
			{
			  "type": "box",
			  "layout": "horizontal",
			  "contents": [
				{
				  "type": "text",
				  "text": "4",
				  "flex": 1,
				  "align": "center"
				},
				{
				  "type": "text",
				  "text": "パワーエレクトロニクス",
				  "flex": 3,
				  "scaling": true,
				  "wrap": true,
				  "size": "10px"
				},
				{
				  "type": "text",
				  "text": "パワーエレクトロニクス",
				  "flex": 3,
				  "scaling": true,
				  "wrap": true,
				  "size": "10px"
				},
				{
				  "type": "text",
				  "text": "パワーエレクトロニクス",
				  "flex": 3,
				  "scaling": true,
				  "wrap": true,
				  "size": "10px"
				},
				{
				  "type": "text",
				  "text": "パワーエレクトロニクス",
				  "flex": 3,
				  "scaling": true,
				  "wrap": true,
				  "size": "10px"
				},
				{
				  "type": "text",
				  "text": "パワーエレクトロニクス",
				  "flex": 3,
				  "scaling": true,
				  "wrap": true,
				  "size": "10px"
				},
				{
				  "type": "text",
				  "text": "パワーエレクトロニクス",
				  "flex": 3,
				  "scaling": true,
				  "wrap": true,
				  "size": "10px"
				},
				{
				  "type": "text",
				  "text": "パワーエレクトロニクス",
				  "flex": 3,
				  "scaling": true,
				  "wrap": true,
				  "size": "10px"
				}
			  ],
			  "spacing": "none"
			},
			{
			  "type": "box",
			  "layout": "horizontal",
			  "contents": [
				{
				  "type": "text",
				  "text": "5",
				  "flex": 1,
				  "align": "center"
				},
				{
				  "type": "text",
				  "text": "パワーエレクトロニクス",
				  "flex": 3,
				  "scaling": true,
				  "wrap": true,
				  "size": "10px"
				},
				{
				  "type": "text",
				  "text": "パワーエレクトロニクス",
				  "flex": 3,
				  "scaling": true,
				  "wrap": true,
				  "size": "10px"
				},
				{
				  "type": "text",
				  "text": "パワーエレクトロニクス",
				  "flex": 3,
				  "scaling": true,
				  "wrap": true,
				  "size": "10px"
				},
				{
				  "type": "text",
				  "text": "パワーエレクトロニクス",
				  "flex": 3,
				  "scaling": true,
				  "wrap": true,
				  "size": "10px"
				},
				{
				  "type": "text",
				  "text": "パワーエレクトロニクス",
				  "flex": 3,
				  "scaling": true,
				  "wrap": true,
				  "size": "10px"
				},
				{
				  "type": "text",
				  "text": "パワーエレクトロニクス",
				  "flex": 3,
				  "scaling": true,
				  "wrap": true,
				  "size": "10px"
				},
				{
				  "type": "text",
				  "text": "パワーエレクトロニクス",
				  "flex": 3,
				  "scaling": true,
				  "wrap": true,
				  "size": "10px"
				}
			  ],
			  "spacing": "none"
			}
		  ],
		  "spacing": "lg"
		}
	  }`

	message := messaging_api.FlexMessage{}

	err := message.UnmarshalJSON([]byte(json))
	if err != nil {
		return messaging_api.FlexMessage{}, err
	}

	return message, nil
}
