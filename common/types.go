package common

import "time"

type Subject struct {
	Name    string
	Weekday time.Weekday
	Period  int
	Room    string
}
