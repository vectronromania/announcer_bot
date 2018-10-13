package commands

import (
	"time"
)

type Reminder struct {
	id       int
	body     string
	datetime time.Time
	channels []string
}

func NewReminder(id int, body string, datetime time.Time, channels []string) Reminder {
	return Reminder{
		id:       id,
		body:     body,
		datetime: datetime,
		channels: channels,
	}
}

// Reminder handler
func reminder(command command) {

}
