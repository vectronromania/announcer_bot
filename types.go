package types

import (
	"fmt"
	"time"
)

const (
	timezone                = "Europe/Bucharest"
	Location *time.Location = time.LoadLocation(timezone)
)

type Attachment struct {
	description string
	link        string
}

type Announcement struct {
	id       int
	body     string
	datetime time.Time
	channels []string
}

type Task struct {
	id          int
	body        string
	assignedto  []string
	duedate     time.Time
	attachments []attachment
}

type Reminder struct {
	id       int
	body     string
	datetime time.Time
	channels []string
}

func ExportAttachment(att Attachment) string {
	return att.description + " : " + att.link
}
