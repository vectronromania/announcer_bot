package announcer_bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func onCommand(session *discordgo.Session, message *discordgo.MessageCreate) {

	if message.Author.ID == session.State.User.ID {
		return
	}
}
