package main

import (
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
		}
	}(ordersChannel)
	<-make(chan struct{})
}
