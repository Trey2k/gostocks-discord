package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func chatListener(discord *discordgo.Session, message *discordgo.MessageCreate) {
	serverID := message.GuildID
	channelID := message.ChannelID

	if serverID == config.Discord.GuildID && channelID == config.Discord.ChannelID {
		ChatParse(message.Content)
	}
}

func discordStatus(discord *discordgo.Session, ready *discordgo.Ready) {
	err := discord.UpdateStatus(1, config.Discord.GameStatus)
	errCheck("Error attempting to set my status", err)
	fmt.Println("Started discord client.")
}
