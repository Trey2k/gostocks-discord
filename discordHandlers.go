package main

import (
	"fmt"

	"github.com/Trey2k/gostocks-discord/utils"
	"github.com/bwmarrin/discordgo"
)

func chatListener(discord *discordgo.Session, message *discordgo.MessageCreate) {
	serverID := message.GuildID
	channelID := message.ChannelID

	if serverID == utils.Config.Discord.GuildID && channelID == utils.Config.Discord.ChannelID {
		ChatParse(message.Content)
	}
}

func discordStatus(discord *discordgo.Session, ready *discordgo.Ready) {
	err := discord.UpdateStatus(1, utils.Config.Discord.GameStatus)
	utils.ErrCheck("Error attempting to set my status", err)
	fmt.Println("Started discord client.")
}
