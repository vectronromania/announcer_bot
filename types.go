package announcer_bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

const prefix string = "!"

type command struct {
	name        string
	params      string
	param_count int
}
