package commands

import (
	// "errors"
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
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
func (queue AnnouncementQueue) Create(payload command) {
	_, exists := queue.Checkexistence(payload)

	if exists != -1 {
		sendsimple(payload.SessionInfo, payload.ChannelID, "There is an existing entry for your user id.")
		return
	}

	item := AnnouncementQueueItem{
		UserID:  payload.UserID,
		Payload: Announcement{},
	}

	defer sendsimple(payload.SessionInfo, payload.ChannelID, "Created an announcement.")

	queue.Items = append(queue.Items, item)

	// fmt.Println(queue.Items)

}

// Function that deletes a user's entry
func (queue AnnouncementQueue) Delete(payload command) {
	_, index := queue.Checkexistence(payload)

	if index == -1 {
		sendsimple(payload.SessionInfo, payload.ChannelID, "There is no entry for your user id.")
		return
	}

	defer sendsimple(payload.SessionInfo, payload.ChannelID, "Deleted an announcement.")

	queue.Items = append(queue.Items[:index], queue.Items[index+1:]...)
}

// Functions that sends the user's entry and deletes it
func (queue AnnouncementQueue) Send(payload command) {
	item, exists := queue.Checkexistence(payload)

	if exists == -1 {
		sendsimple(payload.SessionInfo, payload.ChannelID, "There is no entry for your user id.")
		return
	}

	datefield := makefield("Date&time", item.Payload.Datetime.String(), false)

	fields := []*discordgo.MessageEmbedField{&datefield}

	message := &discordgo.MessageEmbed{
		Title:       "Announcement",
		Description: item.Payload.Body,
		Color:       0xffff00,
		Fields:      fields,
	}

	defer queue.Delete(payload)

	sendembed(payload.SessionInfo, id_announcements, message)

}

// Functions that sets a user's entry a certain attribute
func (queue AnnouncementQueue) Setattribute(payload command) {
	item, exists := queue.Checkexistence(payload)

	if exists == -1 {
		sendsimple(payload.SessionInfo, payload.ChannelID, "There is no entry for your user id.")
		return
	}

	var attr_name string = payload.Params[0]

	var attr_value string = ""
	attr_value += payload.Params[1]
	for i := 2; i < payload.Paramcount; i++ {
		attr_value += " "
		attr_value += payload.Params[i]
	}

	defer sendsimple(payload.SessionInfo, payload.ChannelID, "Set the attribute "+attr_name+" the value: "+attr_value)

	switch attr_name {
	case "body":
		item.Payload.Body = attr_value
	case "datetime":
		item.Payload.Datetime, err = time.ParseInLocation(TimeLayout, attr_value, Location)
	}

}

// Function that checks whether a specific user already has an entry or not
func (queue AnnouncementQueue) Checkexistence(payload command) (AnnouncementQueueItem, int) {
	for i := 0; i < len(queue.Items); i++ {
		fmt.Println(i, queue.Items[i])
		if payload.UserID == queue.Items[i].UserID {
			return queue.Items[i], i
		}
	}
	return AnnouncementQueueItem{}, -1
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
func announcements_error(payload command) {

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
		go announcementQueue.Create(payload)
	case "delete":
		go announcementQueue.Delete(payload)
	case "send":
		go announcementQueue.Send(payload)
	default:
		if announcements_checkifattr(payload.Params[0]) {
			go announcementQueue.Setattribute(payload)
		}
		go announcements_error(payload)
	}
}
