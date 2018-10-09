package types

import (
	// "fmt"
	"time"
)

const (
	timezone string = "Europe/Bucharest"
)

var (
	Location, err = time.LoadLocation(timezone)
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
	attachments []Attachment
}

type Reminder struct {
	id       int
	body     string
	datetime time.Time
	channels []string
}

func NewAttachment(description, link string) Attachment {
	return Attachment{
		description: description,
		link:        link,
	}
}

func NewAnnouncement(id int, body string, datetime time.Time, channels []string) Announcement {
	return Announcement{
		id:       id,
		body:     body,
		datetime: datetime,
		channels: channels,
	}
}

func NewTask(id int, body string, assignedto []string, duedate time.Time, attachments []Attachment) Task {
	return Task{
		id:          id,
		body:        body,
		assignedto:  assignedto,
		duedate:     duedate,
		attachments: attachments,
	}
}

func NewReminder(id int, body string, datetime time.Time, channels []string) Reminder {
	return Reminder{
		id:       id,
		body:     body,
		datetime: datetime,
		channels: channels,
	}
}

func (att Attachment) Export() string {
	return att.description + " : " + att.link
}
