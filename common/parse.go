package common

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
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
		Weekday: parseWeekday(weekday),
		Period:  period,
		Room:    room,
	}, nil

}

func (s Subject) String() string {
	return fmt.Sprintf("%s曜日%d限\n教室: %s\n授業名: %s", WeekToString(s.Weekday), s.Period, s.Room, s.Name)
}

func parseWeekday(s string) time.Weekday {
	switch s {
	case "月":
		return time.Monday
	case "火":
		return time.Tuesday
	case "水":
		return time.Wednesday
	case "木":
		return time.Thursday
	case "金":
		return time.Friday
	case "土":
		return time.Saturday
	case "日":
		return time.Sunday
	default:
		return -1
	}
}

func WeekToString(w time.Weekday) string {
	switch w {
	case time.Monday:
		return "月"
	case time.Tuesday:
		return "火"
	case time.Wednesday:
		return "水"
	case time.Thursday:
		return "木"
	case time.Friday:
		return "金"
	case time.Saturday:
		return "土"
	case time.Sunday:
		return "日"
	default:
		return ""
	}
}
