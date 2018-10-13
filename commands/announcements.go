package commands

import (
	"time"
)

// Container for an announcement
type Announcement struct {
	id       int
	body     string
	datetime time.Time
	// Channels in which the bot would specify to answer to the announcement
	channels []string
}

func NewAnnouncement(id int, body string, datetime time.Time, channels []string) Announcement {
	return Announcement{
		id:       id,
		body:     body,
		datetime: datetime,
		channels: channels,
	}
}

// Announcement handler
func announcement(command command) {

}
