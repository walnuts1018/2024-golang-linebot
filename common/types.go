package common

import (
	"fmt"
	"time"
)

type Subject struct {
	Name    string
	Weekday time.Weekday
	Period  int
	Room    string
}

func (s Subject) String() string {
	return fmt.Sprintf("%s曜日%d限\n教室: %s\n授業名: %s", WeekToString(s.Weekday), s.Period, s.Room, s.Name)
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

func ParseWeekday(s string) time.Weekday {
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
