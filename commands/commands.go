package commands

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const (
	prefix string = "!"

	id_announcements string = "485516381014458389"
	id_tasks         string = "499632407251517451"
	id_reminders     string = "499632385000734743"
)

type command struct {
	Name       string
	Params     []string
	Paramcount int
}

// function that checks if a message is a command
func checkcommand(session *discordgo.Session, message *discordgo.MessageCreate) bool {

	if message.Author.ID == session.State.User.ID {
		return false
	}

	if string(message.Content[0]) != prefix {
		return false
	}
	return true
}

// function that packages the input message into a usable command type
func makecommand(input string) command {

	input = string(input[1:])

	fields := strings.Fields(input)

	var thisisacommand command

	thisisacommand.Name = fields[0]

	if len(fields) > 1 {
		params = make([]string, len(fields)-1)
		for i := 0; i < len(params); i++ {
			params[i] = fields[i+1]
		}
		thisisacommand.Params = params
		thisisacommand.Paramcount = len(params)
	} else {
		thisisacommand.Paramcount = 0
	}

	return thisisacommand

}

// function that handles incoming commands
func OnCommand(session *discordgo.Session, message *discordgo.MessageCreate) {

	if !checkcommand(session, message) {
		return
	}

	command := makecommand(message.Content)

	fmt.Println("Received command " + command.Name)

	switch command.Name {

	case "announcement":
		go announcement(command)

	case "task":
		go task(command)

	case "reminder":
		go reminder(command)

	case "help":
		go help(command)

	default:
		go unknown(command)

	}

	// 	session.ChannelMessageSend(message.ChannelID, "This was a command.")
}

// function that sends useful information about the bot
func help(payload command) {

}

// function that handles unknown commands
func unknown(payload command) {

}
