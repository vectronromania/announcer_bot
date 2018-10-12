package commands

import (
	"github.com/bwmarrin/discordgo"
)

func embed(s *discordgo.Session, chanelID string) {

	field1 := discordgo.MessageEmbedField{
		Name:   "Test 1",
		Value:  "This is a test field.",
		Inline: true,
	}

	field2 := discordgo.MessageEmbedField{
		Name:   "Test 2",
		Value:  "This is a test field.",
		Inline: true,
	}

	fields := []*discordgo.MessageEmbedField{&field1, &field2}

	footer := &discordgo.MessageEmbedFooter{
		Text: "Test footer",
	}

	embedded := &discordgo.MessageEmbed{
		Title:       "Test title",
		Description: "Test description",
		Color:       0x000000,
		Fields:      fields,
		Footer:      footer,
	}

	s.ChannelMessageSendEmbed(chanelID, embedded)
}
