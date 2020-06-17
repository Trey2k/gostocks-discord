package main

import (
	"fmt"
	"os"

	"github.com/Trey2k/gostocks-discord/webapp"
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

	genAuthURL()
	go webapp.Start(callbackAddress, authURL)

	tdauth() //Holding call untill authed

	var response GetAccountResponses
	err := getAccounts(accessToken, &response)
	errCheck("Error getting accounts", err)

	fmt.Println("Cash for trading: " + fmt.Sprint(response[0].SecuritiesAccount.CurrentBalances.CashAvailableForTrading))

	token := config.Discord.Token

	discord, err := discordgo.New(token)
	errCheck("error creating discord session", err)

	discord.AddHandler(chatListener)
	discord.AddHandler(discordStatus)

	err = discord.Open()
	errCheck("Error opening connection to Discord", err)
	defer discord.Close()

	go func(cmdChan chan Commands) {
		for {
			cmd := <-cmdChan
			placeOrder(cmd)
			printCommands(cmd)
		}
	}(channel)

	<-make(chan struct{})

}

func printCommands(commands Commands) {
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	fmt.Println("Buy: " + fmt.Sprint(commands.buy) + ", Ticker: " + commands.ticker + ", ExpDate: " + commands.expDate + ", StrikerPrice: " + commands.strikPrice + ", Buy Price: " + fmt.Sprint(commands.price) + ", Risky: " + fmt.Sprint(commands.risky) + ", Stop Loss: " + fmt.Sprint(commands.stopLoss))
	fmt.Println("------------------------------------------------------------------------------------------------------------")
}
