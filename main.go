package main

import (
	"fmt"
	"os"

	"github.com/Trey2k/gostocks-discord/td"
	"github.com/Trey2k/gostocks-discord/utils"
	"github.com/Trey2k/gostocks-discord/webapp"
	"github.com/bwmarrin/discordgo"
)

var channel = make(chan Commands)

func init() {
	var err error
	utils.Config, err = utils.GetConfig()
	td.Init()

	utils.ErrCheck("Error getting config", err)

	if utils.IsStructEmpty(utils.Config.Discord) {
		println("A value in config.Discord is empty")
		os.Exit(1)
	}
	if utils.IsStructEmpty(utils.Config.TD) {
		println("A value in config.TD is empty")
		os.Exit(1)
	}
}

func main() {
	go webapp.Start(td.CallbackAddress, td.AuthURL)

	td.Auth() //Holding call untill authed

	var response td.GetAccountResponses
	err := td.GetAccounts(&response)
	utils.ErrCheck("Error getting accounts", err)

	fmt.Println("Cash for trading: " + fmt.Sprint(response[0].SecuritiesAccount.CurrentBalances.CashAvailableForTrading))

	discord, err := discordgo.New("")
	utils.ErrCheck("error creating discord session", err)
	utils.ErrCheck("error creating discord session", discord.Login(utils.Config.Discord.Username, utils.Config.Discord.Password))

	discord.AddHandler(chatListener)
	discord.AddHandler(discordStatus)

	err = discord.Open()
	utils.ErrCheck("Error opening connection to Discord", err)
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
