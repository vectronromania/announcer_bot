package commands

// Function that sends useful information about the bot
func help(payload command) {
	sendsimple(payload.SessionInfo, payload.ChannelID, "The help command is not available.")
}
