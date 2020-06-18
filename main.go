package main

import (
	"fmt"
	"os"

	"github.com/Trey2k/gostocks-discord/td"
	"github.com/Trey2k/gostocks-discord/utils"
	"github.com/Trey2k/gostocks-discord/webapp"
	"github.com/bwmarrin/discordgo"
)

var cmdsChannel = make(chan Commands)

func init() {
	var err error
	utils.Config, err = utils.GetConfig()
	utils.ErrCheck("Error getting config", err)

	td.Init()

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

	var response td.GetAccountResponse
	err := td.GetAccount(utils.Config.TD.AccountID, &response)
	utils.ErrCheck("Error getting accounts", err)

	fmt.Println("Cash aval for trading: " + fmt.Sprint(response.SecuritiesAccount.CurrentBalances.CashAvailableForTrading))
	fmt.Println("Total ball: " + fmt.Sprint(response.SecuritiesAccount.CurrentBalances.CashBalance))
	fmt.Println("Inital ball: " + fmt.Sprint(response.SecuritiesAccount.InitialBalances.CashBalance))

	discord, err := discordgo.New("")
	utils.ErrCheck("error creating discord session", err)

	err = discord.Login(utils.Config.Discord.Username, utils.Config.Discord.Password)
	utils.ErrCheck("error creating discord session", err)

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
	}(cmdsChannel)
	<-make(chan struct{})
}

func printCommands(commands Commands) {
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	fmt.Println("Buy: " + fmt.Sprint(commands.buy) + ", Ticker: " + commands.ticker + ", Date: " + commands.expDate.Format("1/2/2006") + ", StrikerPrice: " + commands.strikPrice + ", Buy Price: " + fmt.Sprint(commands.price) + ", Risky: " + fmt.Sprint(commands.risky) + ", Stop Loss: " + fmt.Sprint(commands.stopLoss))
	fmt.Println("------------------------------------------------------------------------------------------------------------")
}
