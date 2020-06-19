package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/Trey2k/gostocks-discord/td"
	"github.com/Trey2k/gostocks-discord/utils"
	"github.com/bwmarrin/discordgo"
)

//Order Order built from Discord command

//ChatParse : Parse a chat message and build a array of commands
func ChatParse(msg string, sender discordgo.User, messageID string) {
	var err error
	var order utils.OrderStruct

	order.Sender = sender
	order.MessageID = messageID

	msg = strings.ToUpper(strings.Split(msg, "\n")[0])

	order.Risky = false
	if strings.Contains(msg, "RISKY") || strings.Contains(msg, "LOTTO") {
		order.Risky = true
		msg = strings.Replace(msg, "RISKY", "", 1)
		msg = strings.Replace(msg, "LOTTO", "", 1)
	}

	if strings.Contains(msg, "BTO") {
		order.Buy = true
		msg = strings.Replace(msg, "BTO", "", 1)
	} else {
		order.Buy = false
		msg = strings.Replace(msg, "STC", "", 1)
	}

	msgs := strings.Split(msg, " ")

	for i := 0; i < len(msgs); i++ {
		cmd := msgs[i]
		if i <= 5 && td.IsValidTicker(cmd) {
			order.Ticker = cmd
		} else {
			if strings.Contains(cmd, "/") && utils.IsNumericIgnore(cmd, "/", 2) {

				dates := strings.Split(cmd, "/")
				if len(dates) == 2 {
					date, err := time.Parse("1/2/2006", cmd+"/"+fmt.Sprint(time.Now().Year()))
					if err != nil {
						println("Error converting string date '" + cmd + "' to date: " + err.Error())
						date = time.Now()
					}
					order.ExpDate = date
				} else if len(dates) == 3 {
					date, err := time.Parse("1/2/2006", cmd)
					if err != nil {
						println("Error converting string date '" + cmd + "' to date: " + err.Error())
						date = time.Now()
					}
					order.ExpDate = date
				} else {
					println("Error converting string date '" + cmd + "' to date: Unknown format. Settig exp date for today")
					order.ExpDate = time.Now()
				}
				if order.ExpDate.Year() <= time.Now().Year() && order.ExpDate.YearDay() <= time.Now().YearDay()+1 && order.Buy == true {
					order.Risky = true
				}

			} else if strings.Contains(cmd, "P") && utils.IsNumericIgnore(cmd, "P", 1) || strings.Contains(cmd, "C") && utils.IsNumericIgnore(cmd, "C", 1) {

				if strings.Contains(cmd, "C") {
					x, err := utils.ToNumericIgnore(cmd, "C", 1)
					if err != nil {
						println("Error converting strike price '" + cmd + "' to int64: " + err.Error())
					}
					order.StrikPrice = x
					order.ContractType = "CALL"
				} else {
					x, err := utils.ToNumericIgnore(cmd, "P", 1)
					if err != nil {
						println("Error converting strike price '" + cmd + "' to int64: " + err.Error())
					}
					order.StrikPrice = x
					order.ContractType = "PUT"
				}

			} else if strings.Contains(cmd, ".") && utils.IsNumericIgnore(cmd, "@", 1) {
				if order.Price == 0 {
					order.Price, err = utils.ToNumericIgnore(cmd, "@", 1)
					if err != nil {
						println("Error converting price '" + cmd + "' to float64: " + err.Error())
					}
				} else if order.StopLoss == 0 {
					order.StopLoss, err = utils.ToNumeric(cmd)
					if err != nil {
						println("Error converting stop loss '" + cmd + "' to float64: " + err.Error())
					}
				}
			}
		}
	}
	if order.StopLoss == 0 {
		if order.Risky {
			order.StopLoss = order.Price * utils.Config.Settings.Trade.RiskyStopLossPercent
		} else {
			order.StopLoss = order.Price * utils.Config.Settings.Trade.SafeStopLossPercent
		}
	}

	ordersChannel <- order
}
