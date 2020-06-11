package main

import (
	"os"

	"github.com/bwmarrin/discordgo"
)

var channel = make(chan Commands)

func init() {
	var err error
	config, err = getConfig()

	errCheck("Error getting config", err)

	if isStructEmpty(config.Discord) {
		println("A value in config.Discord is empty")
		os.Exit(1)
	}
	if isStructEmpty(config.IB) {
		println("A value in config.IB is empty")
		os.Exit(1)
	}
}

func main() {
	go connectIB(channel)

	token := config.Discord.Token

	discord, err := discordgo.New(token)
	errCheck("error creating discord session", err)

	discord.AddHandler(chatListener)
	discord.AddHandler(discordStatus)

	err = discord.Open()
	errCheck("Error opening connection to Discord", err)
	defer discord.Close()

	<-make(chan struct{})
}
