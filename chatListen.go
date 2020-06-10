package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func chatListener(discord *discordgo.Session, message *discordgo.MessageCreate) {
	user := message.Author
	serverID := message.GuildID
	channelID := message.ChannelID

	if serverID == config.GuildID && channelID == config.ChannelID {
		fmt.Println(user.Username, ": ", message.Content)
	}
}
