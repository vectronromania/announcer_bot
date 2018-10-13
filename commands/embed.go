package commands

import (
	"github.com/bwmarrin/discordgo"
)

// Function that creates an embedded message field
func makefield(name string, value string, inline bool) discordgo.MessageEmbedField {
	return discordgo.MessageEmbedField{
		Name:   name,
		Value:  value,
		Inline: inline,
	}
}

// Function that creates an embedded message footer(text only)
func makefooter(text string) discordgo.MessageEmbedFooter {
	return discordgo.MessageEmbedFooter{
		Text: text,
	}
}

// func embed(s *discordgo.Session, chanelID string) {

// 	field1 := discordgo.MessageEmbedField{
// 		Name:   "Test 1",
// 		Value:  "This is a test field.",
// 		Inline: true,
// 	}

// 	field2 := discordgo.MessageEmbedField{
// 		Name:   "Test 2",
// 		Value:  "This is a test field.",
// 		Inline: true,
// 	}

// 	fields := []*discordgo.MessageEmbedField{&field1, &field2, field3}

// 	footer := &discordgo.MessageEmbedFooter{
// 		Text: "Test footer",
// 	}

// 	embedded := &discordgo.MessageEmbed{
// 		Title:       "Test title",
// 		Description: "Test description",
// 		Color:       0x000000,
// 		Fields:      fields,
// 		Footer:      footer,
// 	}

// 	sendembed(s, chanelID, embedded)
// }
