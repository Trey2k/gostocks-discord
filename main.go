package main

import (
	"fmt"
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
	if isStructEmpty(config.TD) {
		println("A value in config.TD is empty")
		os.Exit(1)
	}
}

func main() {
	token := config.Discord.Token

	discord, err := discordgo.New(token)
	errCheck("error creating discord session", err)

	discord.AddHandler(chatListener)
	discord.AddHandler(discordStatus)

	err = discord.Open()
	errCheck("Error opening connection to Discord", err)
	defer discord.Close()

	go tdauth()
	go func(cmdChan chan Commands) {
		for {
			printCommands(<-cmdChan)
		}
	}(channel)

	<-make(chan struct{})
}

func printCommands(commands Commands) {
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	fmt.Println("Buy/Sell: " + commands.buysell + ", Ticker: " + commands.ticker + ", ExpDate: " + commands.expDate + ", StrikerPrice: " + commands.strikPrice + ", Buy Price: " + fmt.Sprint(commands.price) + ", Danger: " + commands.danger + ", Stop Loss: " + fmt.Sprint(commands.stopLoss))
	fmt.Println("------------------------------------------------------------------------------------------------------------")
}
