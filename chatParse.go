package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/Trey2k/gostocks-discord/utils"
	"github.com/bwmarrin/discordgo"
	"github.com/pkg/errors"
)

//Order Order built from Discord command

//ChatParse : Parse a chat message and build a array of commands
func ChatParse(msg string, sender discordgo.User, messageID string) utils.OrderStruct {
	var err error
	var order utils.OrderStruct

	order.Sender = sender
	order.MessageID = messageID
	order.Message = msg

	msg = strings.ToUpper(strings.Split(msg, "\n")[0])

	order.Risky = false
	if strings.Contains(msg, "RISK") || strings.Contains(msg, "LOTTO") {
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

		if !strings.Contains(cmd, "http") {

			if i <= 3 && isValidTicker(cmd) {
				order.Ticker = cmd
			} else {
				if strings.Contains(cmd, "/") && utils.IsNumericIgnore(cmd, "/", 2) && order.ExpDate.IsZero() {

					dates := strings.Split(cmd, "/")
					if len(dates) == 2 {
						date, err := time.Parse("1/2/2006", cmd+"/"+fmt.Sprint(time.Now().Year()))
						if err != nil {
							utils.Log("Error converting string date '"+cmd+"' to date: "+errors.WithStack(err).Error(), utils.LogError)
						}
						order.ExpDate = date
					} else if len(dates) == 3 {
						date, err := time.Parse("1/2/2006", cmd)
						if err != nil {
							utils.Log("Error converting string date '"+cmd+"' to date: "+errors.WithStack(err).Error(), utils.LogError)
						}
						order.ExpDate = date
					} else {
						utils.Log("Error converting string date '"+cmd+"' to date: "+errors.WithStack(err).Error(), utils.LogError)
					}
					future := time.Now().Add(time.Hour * 24)
					exp := order.ExpDate
					if exp.Before(future) {
						order.Risky = true
					}

				} else if strings.Contains(cmd, "P") && utils.IsNumericIgnore(cmd, "P", 1) || strings.Contains(cmd, "C") && utils.IsNumericIgnore(cmd, "C", 1) {

					if strings.Contains(cmd, "C") {
						x, err := utils.ToNumericIgnore(cmd, "C", 1)
						if err != nil {
							utils.Log("Error converting strike price '"+cmd+"' to number: "+errors.WithStack(err).Error(), utils.LogError)
						}
						order.StrikPrice = x
						order.ContractType = "CALL"
					} else {
						x, err := utils.ToNumericIgnore(cmd, "P", 1)
						if err != nil {
							utils.Log("Error converting strike price '"+cmd+"' to number: "+errors.WithStack(err).Error(), utils.LogError)
						}
						order.StrikPrice = x
						order.ContractType = "PUT"
					}

				} else if strings.Contains(cmd, ".") && utils.IsNumericIgnore(cmd, "@", 1) {
					if order.Price == 0 {
						order.Price, err = utils.ToNumericIgnore(cmd, "@", 1)
						if err != nil {
							utils.Log("Error converting price '"+cmd+"' to number: "+errors.WithStack(err).Error(), utils.LogError)
						}
					} else if order.StopLoss == 0 {
						order.StopLoss, err = utils.ToNumeric(cmd)
						if err != nil {
							utils.Log("Error converting price '"+cmd+"' to number: "+errors.WithStack(err).Error(), utils.LogError)
						}
					}
				}
			}
		}
	}
	if order.StopLoss == 0 {
		if order.Risky {
			order.StopLoss = float64(int((order.Price*utils.Config.Settings.Trade.RiskyStopLossPercent)*100)) / 100
		} else {
			order.StopLoss = float64(int((order.Price*utils.Config.Settings.Trade.SafeStopLossPercent)*100)) / 100
		}
	}

	return order
}
