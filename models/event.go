package models

import "time"

type Event struct {
	Id int
	Name string
	Description string
	Location string
	DateTime time.Time
	UserId int
}
