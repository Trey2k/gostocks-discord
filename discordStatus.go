package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func discordStatus(discord *discordgo.Session, ready *discordgo.Ready) {
	err := discord.UpdateStatus(1, config.GameStatus)
	errCheck("Error attempting to set my status", err)
	fmt.Println("Started discord client.")
}
