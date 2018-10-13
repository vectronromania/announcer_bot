package commands

import (
	"time"
)

type Task struct {
	id          int
	body        string
	assignedto  []string
	duedate     time.Time
	attachments []Attachment
}

type Attachment struct {
	description string
	link        string
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

func NewAttachment(description, link string) Attachment {
	return Attachment{
		description: description,
		link:        link,
	}
}

func (att Attachment) Export() string {
	return att.description + " : " + att.link
}

// Task handler
func task(command command) {

}
