package main

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func chatListener(discord *discordgo.Session, message *discordgo.MessageCreate) {
	serverID := message.GuildID
	channelID := message.ChannelID

	if serverID == config.GuildID && channelID == config.ChannelID {
		var commands Commands
		ChatParse(message.Content, &commands)
		fmt.Println("------------------------------------------------------------------------------------------------------------")
		fmt.Print("Buy/Sell: " + commands.buysell + ", Ticker: " + commands.ticker + ", ExpDate: " + commands.expDate + ", StrikerPrice: " + commands.strikPrice + ", Buy Price: " + fmt.Sprint(commands.price) + ", Danger: " + commands.danger + ", Stop Loss: " + fmt.Sprint(commands.stopLoss))
		fmt.Println("\n------------------------------------------------------------------------------------------------------------")
	}
}
