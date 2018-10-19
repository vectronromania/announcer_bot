package commands

import (
	"time"
)

// Container for an announcement
type Announcement struct {
	ID       int
	Body     string
	Datetime time.Time
	// Channels in which the bot would specify to answer to the announcement
	Channels []string
}

type AnnouncementQueueItem struct {
	UserID  string
	Payload Announcement
}

// Queue of announcements that are currently constructed
type AnnouncementQueue struct {
	Items []AnnouncementQueueItem
}

var (
	announcementQueue AnnouncementQueue = AnnouncementQueue{make([]AnnouncementQueueItem, 0)}
	attr_list         []string          = []string{"body", "tag", "datetime", "channel"}
)

// Function that creates a new announcement for the current user
func (queue AnnouncementQueue) Create(userID string) {
	_, exists := queue.Checkexistence(userID)

	if !exists {
		return
	}

	item := AnnouncementQueueItem{
		UserID:  userID,
		Payload: Announcement{},
	}

	queue.Items = append(queue.Items, item)

}

// Function that deletes a user's entry
func (queue AnnouncementQueue) Delete(userID string) {

}

// Functions that sends the user's entry and deletes it
func (queue AnnouncementQueue) Send(userID string) {

}

// Functions that sets a user's entry a certain attribute
func (queue AnnouncementQueue) Setattribute(userID string, attr_name string, attr_value string) {

}

// Function that checks whether a specific user already has an entry or not
func (queue AnnouncementQueue) Checkexistence(userID string) (int, bool) {
	for index, q_userID := range queue.Items {
		if userID == q_userID.UserID {
			return index, true
		}
	}
	return -1, false
}

// Function that checks whether a command parameter is an announcement attribute
// in order to set it accordingly
func announcements_checkifattr(input string) bool {
	for _, attr := range attr_list {
		if input == attr {
			return true
		}
	}
	return false
}

// Function that signals an error in the syntax of the announcement command
func announcements_error() {

}

func NewAnnouncement(id int, body string, datetime time.Time, channels []string) Announcement {
	return Announcement{
		ID:       id,
		Body:     body,
		Datetime: datetime,
		Channels: channels,
	}
}

// Announcement handler
func announcement(payload command) {
	switch payload.Params[0] {
	case "create":
		go announcementQueue.Create(payload.UserID)
	case "delete":
		go announcementQueue.Delete(payload.UserID)
	case "send":
		go announcementQueue.Send(payload.UserID)
	default:
		if announcements_checkifattr(payload.Params[0]) {
			var attr_value string = ""
			attr_value += payload.Params[1]
			for i := 2; i < payload.Paramcount; i++ {
				attr_value += " "
				attr_value += payload.Params[i]
			}
			go announcementQueue.Setattribute(payload.UserID, payload.Params[0], attr_value)
		}
		go announcements_error()
	}
}
