package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func chatListener(discord *discordgo.Session, message *discordgo.MessageCreate) {
	serverID := message.GuildID
	channelID := message.ChannelID

	if serverID == config.GuildID && channelID == config.ChannelID {
		var commands [7]string
		commands = ChatParse(message.Content)
		fmt.Println("-----------------------------------------")
		fmt.Print("Commands: ")
		for i := 0; i < len(commands); i++ {
			fmt.Print(commands[i] + ", ")
			commands[i] = ""
		}
		fmt.Println("\n-----------------------------------------")
	}
}
