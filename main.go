package main

import (
	"fmt"

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

	discord, err := discordgo.New("")
	utils.ErrCheck("error creating discord session", err)

	err = discord.Login(utils.Config.Discord.Username, utils.Config.Discord.Password)
	utils.ErrCheck("error creating discord session", err)

	discord.AddHandler(chatListener)
	discord.AddHandler(discordStatus)

	err = discord.Open()
	utils.ErrCheck("Error opening connection to Discord", err)
	defer discord.Close()

	go func(cmdChan chan utils.OrderStruct) {
		for {
			cmd := <-cmdChan
			placeOrder(cmd)
			printCommands(cmd)
		}
	}(ordersChannel)
	<-make(chan struct{})
}

func printCommands(order utils.OrderStruct) {
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	fmt.Println("Buy: " + fmt.Sprint(order.Buy) + ", Ticker: " + order.Ticker + ", Date: " + order.ExpDate.Format("1/2/2006") + ", StrikerPrice: " + fmt.Sprint(order.StrikPrice) + ", ContractType: " + order.ContractType + ", Buy Price: " + fmt.Sprint(order.Price) + ", Risky: " + fmt.Sprint(order.Risky) + ", Stop Loss: " + fmt.Sprint(order.StopLoss))
	fmt.Println("------------------------------------------------------------------------------------------------------------")
}
