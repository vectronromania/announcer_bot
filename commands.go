package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

const (
	prefix string = "!"
)

type command struct {
	name        string
	params      string
	param_count int
}

func OnCommand(session *discordgo.Session, message *discordgo.MessageCreate) {

	if message.Author.ID == session.State.User.ID {
		return
	}
}

func CheckIfCommand()
