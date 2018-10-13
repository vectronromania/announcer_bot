package commands

import (
	"github.com/bwmarrin/discordgo"
)

// Function that sends a simple text message
func sendsimple(session *discordgo.Session, channelid string, content string) {
	session.ChannelMessageSend(channelid, content)
}

// Function that sends a complex message
func sendcomplex(session *discordgo.Session, chanelid string, content *discordgo.MessageSend) {
	session.ChannelMessageSendComplex(chanelid, content)
}

// Function that sends an embedded message
func sendembed(session *discordgo.Session, chanelid string, embed *discordgo.MessageEmbed) {
	session.ChannelMessageSendEmbed(chanelid, embed)
}
