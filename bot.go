package main

import (
  "fmt"
  "os"
  "os/signal"
  "syscall"

  "github.com/cezarmathe/announcer_bot/commands"

  "github.com/bwmarrin/discordgo"
)

var (
// token string = os.Getenv("DISCORD_BOT_TOKEN")
)

func main() {

  discord, err := discordgo.New("Bot " + token)
  if err != nil {
    fmt.Println("Error creating a Discord session,", err.Error())
    return
  }

  discord.AddHandler(commands.OnCommand)

  err = discord.Open()
  if err != nil {
    fmt.Println("Error opening connection,", err.Error())
  }

  fmt.Println("Bot is now running.  Press CTRL-C to exit.")
  sc := make(chan os.Signal, 1)
  signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
  <-sc

  // Cleanly close down the Discord session.
  discord.Close()

}
