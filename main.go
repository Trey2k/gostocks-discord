package main

import (
	"os"

	"github.com/bwmarrin/discordgo"
)

var botID string

func init() {
	var err error
	config, err = getConfig()
	errCheck("Error getting config", err)

	if config.Token == "" || config.GuildID == "" || config.ChannelID == "" || config.GameStatus == "" {
		println("A value in config.json is empty")
		os.Exit(1)
	}
}

func main() {

	token := config.Token

	discord, err := discordgo.New(token)
	errCheck("error creating discord session", err)

	user, err := discord.User("@me")
	errCheck("error retrieving account", err)

	discord.AddHandler(chatListener)
	discord.AddHandler(discordStatus)
	botID = user.ID

	err = discord.Open()
	errCheck("Error opening connection to Discord", err)
	defer discord.Close()

	<-make(chan struct{})
}
