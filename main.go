package main

import (
	"strings"

	"github.com/Trey2k/gostocks-discord/mysql"
	"github.com/Trey2k/gostocks-discord/td"
	"github.com/Trey2k/gostocks-discord/utils"
	"github.com/Trey2k/gostocks-discord/webapp"
	"github.com/bwmarrin/discordgo"
)

var ordersChannel = make(chan utils.OrderStruct)

func init() {
	//Initializing packages
	utils.Init()
	mysql.Init()
	td.Init()
}

func main() {
	go webapp.Start(td.CallbackAddress, td.AuthURL)
	td.Auth() //Holding call untill authed
	updateMarketHours()

	discord, err := discordgo.New(utils.Config.Discord.Token)
	utils.ErrCheck("error creating discord session", err)

	discord.AddHandler(chatListener)
	discord.AddHandler(editListener)
	discord.AddHandler(discordStatus)

	err = discord.Open()
	utils.ErrCheck("Error opening connection to Discord", err)
	defer discord.Close()

	go update(utils.Config.Settings.Trade.UpdateInterval)
	procOrder(ordersChannel)
}

//IsValidTicker test if string is a valid ticker
func isValidTicker(ticker string) bool {
	if len(ticker) <= 5 && utils.NoNumbers(ticker) && ticker != "BTO" && ticker != "STC" {
		var quoteResponse td.GetQuoteResponse

		err := td.GetQuote(ticker, &quoteResponse)
		utils.ErrCheck("Error testing is valid ticker for "+ticker, err)

		if quoteResponse.Symbol == strings.ToUpper(ticker) {
			return true
		}
	}
	return false
}
