package main

import (
	"fmt"

	"github.com/Trey2k/gostocks-discord/utils"
	"github.com/bwmarrin/discordgo"
)

func chatListener(discord *discordgo.Session, message *discordgo.MessageCreate) {
	serverID := message.GuildID
	channelID := message.ChannelID

	for i := 0; i < len(utils.Config.Discord.GuildIDs); i++ {
		if serverID == utils.Config.Discord.GuildIDs[i] && channelID == utils.Config.Discord.ChannelIDs[i] {
			order := ChatParse(message.Content, *message.Author, message.ID)
			ordersChannel <- order
		}
	}
}

func discordStatus(discord *discordgo.Session, ready *discordgo.Ready) {
	err := discord.UpdateStatus(1, utils.Config.Discord.GameStatus)
	utils.ErrCheck("Error attempting to set my status", err)
	fmt.Println("Started discord client.")
}
