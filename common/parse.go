package common

import (
	"fmt"
	"regexp"
	"strconv"
)

var regex = regexp.MustCompile(`^(?P<weekday>[月火水木金土日])(?P<period>\d)\s(?P<name>.+)\s(?P<room>.+)\s?$`)

func ParseFromString(s string) (Subject, error) {
	params := regex.FindStringSubmatch(s)

	if len(params) != regex.NumSubexp()+1 {
		return Subject{}, fmt.Errorf("形式が間違っています。「月1 パワーエレクトロニクス 電総大」などと入力してください")
	}

	weekday := params[regex.SubexpIndex("weekday")]
	name := params[regex.SubexpIndex("name")]
	room := params[regex.SubexpIndex("room")]

	period, err := strconv.Atoi(params[regex.SubexpIndex("period")])
	if err != nil {
		return Subject{}, fmt.Errorf("時間の取得に失敗しました: %v", err)
	}

	return Subject{
		Name:    name,
		Weekday: ParseWeekday(weekday),
		Period:  period,
		Room:    room,
	}, nil
}
