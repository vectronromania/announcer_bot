package main

import (
  "fmt"
  "os"

  "github.com/bwmarrin/discordgo"
)

var (
  token string
)

func init() {
  string = os.Getenv("DISCORD_BOT_TOKEN")
}

func main() {

  init()

  discord, err := discordgo.New("Bot " + token)
  if err != nil {
    fmt.Println("Error creating a Discord session,", err.Error())
    return
  }

  err = discord.Open()
  if err != nil {
    fmt.Println("Error opening connection,", err.Error())
  }

}
