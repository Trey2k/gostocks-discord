package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var commands [5]string

func chatListener(discord *discordgo.Session, message *discordgo.MessageCreate) {
	serverID := message.GuildID
	channelID := message.ChannelID

	if serverID == config.GuildID && channelID == config.ChannelID {
		chatParse(message.Content, &commands)
		fmt.Println("-----------------------------------------")
		fmt.Print("Commands: ")
		for i := 0; i < 5; i++ {
			fmt.Print(commands[i] + ", ")
			commands[i] = ""
		}
		fmt.Println("\n-----------------------------------------")
	}
}
