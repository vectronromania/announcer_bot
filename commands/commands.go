package commands

import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

const (
	// Put your timezone here
	timezone string = "Europe/Bucharest"

	// The prefix for this bot. Change it according to your needs
	prefix string = "_"

	// Channel IDs for the channels where this bot is going to send messages
	id_announcements string = "485516381014458389"
	id_tasks         string = "499632407251517451"
	id_reminders     string = "499632385000734743"
)

var (
	// Loading the location from the timezone
	// Make sure your timezone is according to the standards, otherwise the bot will crash
	Location, err = time.LoadLocation(timezone)
)

// A type that acts as a container for a command
type command struct {
	Name        string
	Params      []string
	Paramcount  int
	ChanelID    string
	SessionInfo discordgo.Session
}

// Function that checks if a message is a command
func checkcommand(session *discordgo.Session, message *discordgo.MessageCreate) bool {

	if message.Author.ID == session.State.User.ID {
		return false
	}

	if string(message.Content[0]) != prefix {
		return false
	}
	return true
}

// Function that packages the input message into a usable command type
func makecommand(session *discordgo.Session, message *discordgo.MessageCreate) command {

	input := string(message.Content[1:])

	fields := strings.Fields(input)

	var thisisacommand command

	thisisacommand.Name = fields[0]

	if len(fields) > 1 {
		params := make([]string, len(fields)-1)
		for i := 0; i < len(params); i++ {
			params[i] = fields[i+1]
		}
		thisisacommand.Params = params
		thisisacommand.Paramcount = len(params)
	} else {
		thisisacommand.Paramcount = 0
	}

	thisisacommand.ChanelID = message.ChannelID
	thisisacommand.SessionInfo = *session

	return thisisacommand

}

// Function that handles incoming commands
func OnCommand(session *discordgo.Session, message *discordgo.MessageCreate) {

	if !checkcommand(session, message) {
		return
	}

	command := makecommand(session, message)

	fmt.Println("Received command " + command.Name)

	// Check for the command and call it's handler
	switch command.Name {

	case "announcement":
		// go announcement(command)

	case "task":
		// go task(co?mand)

	case "reminder":
		// go reminder(command)

	case "help":
		go help(command)

	case "test":
		go embed(session, message.ChannelID)

	default:
		go unknown(command)

	}
}

// Function that sends useful information about the bot
func help(payload command) {

}

// Function that handles unknown commands
func unknown(payload command) {

}
