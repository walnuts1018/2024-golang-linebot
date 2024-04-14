package common

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

func CreateCalenderJson(subjects []Subject) (*messaging_api.FlexMessage, error) {

	Frames := make([]messaging_api.FlexComponentInterface, 5)

	for i := 0; i < 5; i++ {
		classes := make([]messaging_api.FlexComponentInterface, 6)
		classes[0] = &messaging_api.FlexText{
			Text:  fmt.Sprintf("%d", i+1),
			Flex:  1,
			Align: messaging_api.FlexTextALIGN_CENTER,
		}

		for j := 1; j < 6; j++ {
			classes[j] = &messaging_api.FlexText{
				Text:    " ",
				Flex:    3,
				Scaling: true,
				Wrap:    true,
				Size:    "13px",
				Align:   messaging_api.FlexTextALIGN_CENTER,
			}
		}

		Frame := &messaging_api.FlexBox{
			Layout:         messaging_api.FlexBoxLAYOUT_HORIZONTAL,
			Contents:       classes,
			Spacing:        string(messaging_api.FlexBoxSpacing_NONE),
			AlignItems:     messaging_api.FlexBoxALIGN_ITEMS_CENTER,
			JustifyContent: messaging_api.FlexBoxJUSTIFY_CONTENT_CENTER,
		}
		Frames[i] = Frame
	}

	for _, subject := range subjects {
		if subject.Period > 5 || subject.Period < 1 {
			slog.Error(fmt.Sprintf("Invalid period: %d", subject.Period))
			continue
		}

		if subject.Weekday == time.Saturday || subject.Weekday == time.Sunday {
			slog.Error(fmt.Sprintf("Invalid weekday: %s", subject.Weekday))
			continue
		}

		frames, ok := Frames[subject.Period-1].(*messaging_api.FlexBox)
		if !ok {
			slog.Error(fmt.Sprintf("Invalid frame: %v", Frames[subject.Period-1]))
			continue
		}

		text, ok := frames.Contents[subject.Weekday-1].(*messaging_api.FlexText)
		if !ok {
			slog.Error(fmt.Sprintf("Invalid text: %v", frames.Contents[subject.Weekday]))
			continue
		}

		text.Text = subject.String()
	}

	bodyContents := []messaging_api.FlexComponentInterface{
		&messaging_api.FlexBox{
			Layout: "horizontal",
			Contents: []messaging_api.FlexComponentInterface{
				&messaging_api.FlexText{
					Text: "　",
					Flex: 1,
				},
				&messaging_api.FlexText{
					Text:  "月",
					Flex:  3,
					Align: messaging_api.FlexTextALIGN_CENTER,
				},
				&messaging_api.FlexText{
					Text:  "火",
					Flex:  3,
					Align: messaging_api.FlexTextALIGN_CENTER,
				},
				&messaging_api.FlexText{
					Text:  "水",
					Flex:  3,
					Align: messaging_api.FlexTextALIGN_CENTER,
				},
				&messaging_api.FlexText{
					Text:  "木",
					Flex:  3,
					Align: messaging_api.FlexTextALIGN_CENTER,
				},
				&messaging_api.FlexText{
					Text:  "金",
					Flex:  3,
					Align: messaging_api.FlexTextALIGN_CENTER,
				},
			},
			CornerRadius:   string(messaging_api.FlexBoxCornerRadius_NONE),
			JustifyContent: messaging_api.FlexBoxJUSTIFY_CONTENT_CENTER,
			AlignItems:     messaging_api.FlexBoxALIGN_ITEMS_CENTER,
			Spacing:        string(messaging_api.FlexBoxSpacing_NONE),
		},
		messaging_api.FlexSeparator{},
	}

	bodyContents = append(bodyContents, Frames...)

	message := messaging_api.FlexMessage{
		AltText: "カレンダー",
		Contents: messaging_api.FlexBubble{
			Body: &messaging_api.FlexBox{
				Layout:          messaging_api.FlexBoxLAYOUT_VERTICAL,
				Contents:        bodyContents,
				Spacing:         string(messaging_api.FlexBoxSpacing_LG),
				BackgroundColor: "#f2f5f7",
			},
		},
	}

	return &message, nil
}
