package common

import (
	"fmt"
	"regexp"
	"time"
)

var regex = regexp.MustCompile(`^(?P<weekday>[月火水木金土日])?(?P<period>\d) (?P<name>.+) (?P<room>.+)$`)

func ParseFromString(s string) (Subject, error) {
	params := regex.FindStringSubmatch(s)

	if len(params) != regex.NumSubexp() {
		return Subject{}, fmt.Errorf("形式が間違っています。「月1 パワーエレクトロニクス 電総大」などと入力してください")
	}

	weekday := params[regex.SubexpIndex("weekday")]
	period := params[regex.SubexpIndex("period")]
	name := params[regex.SubexpIndex("name")]
	room := params[regex.SubexpIndex("room")]

	return Subject{
		Name:    name,
		Weekday: parseWeekday(weekday),
		Period:  parsePeriod(period),
		Room:    room,
	}, nil

}

func (s Subject) String() string {
	return fmt.Sprintf("%s曜日%d限\n教室: %s\n授業名: %s", s.Weekday, s.Period, s.Room, s.Name)
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
		return time.Monday
	}
}

func parsePeriod(s string) int {
	switch s {
	case "1":
		return 1
	case "2":
		return 2
	case "3":
		return 3
	case "4":
		return 4
	case "5":
		return 5
	case "6":
		return 6
	default:
		return 1
	}
}
