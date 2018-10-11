package commands

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

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
	for i := 1; i < len(fields); i++ {
		thisisacommand.Params[i-1] = fields[i]
	}
	thisisacommand.Paramcount = len(thisisacommand.Params)

	return thisisacommand

}

// function that handles incoming commands
func OnCommand(session *discordgo.Session, message *discordgo.MessageCreate) {

	if !checkcommand(session, message) {
		return
	}

	command := makecommand(message.Content)

	fmt.Println("Received command " + command.Name)

	session.ChannelMessageSend(message.ChannelID, "This was a command.")
}
