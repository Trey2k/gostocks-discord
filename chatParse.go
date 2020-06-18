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
type Order struct {
	buy        bool
	ticker     string
	expDate    time.Time
	strikPrice string
	price      float64
	risky      bool
	stopLoss   float64
	sender     discordgo.User
	messageID  string
}

//ChatParse : Parse a chat message and build a array of commands
func ChatParse(msg string, sender discordgo.User, messageID string) {
	var err error
	var order Order

	order.sender = sender
	order.messageID = messageID

	msg = strings.ToLower(strings.Split(msg, "\n")[0])

	order.risky = false
	if strings.Contains(msg, "risky") || strings.Contains(msg, "lotto") {
		order.risky = true
		msg = strings.Replace(msg, "riksy", "", 1)
		msg = strings.Replace(msg, "lotto", "", 1)
	}

	if strings.Contains(msg, "bto") {
		order.buy = true
		msg = strings.Replace(msg, "bto", "", 1)
	} else {
		order.buy = false
		msg = strings.Replace(msg, "stc", "", 1)
	}

	msgs := strings.Split(msg, " ")

	for i := 0; i < len(msgs); i++ {
		cmd := msgs[i]
		if i <= 5 && td.IsValidTicker(cmd) {
			order.ticker = cmd
		} else {
			if strings.Contains(cmd, "/") && utils.IsNumericIgnore(cmd, "/", 2) {

				dates := strings.Split(cmd, "/")
				if len(dates) == 2 {
					date, err := time.Parse("1/2/2006", cmd+"/"+fmt.Sprint(time.Now().Year()))
					if err != nil {
						println("Error converting string date '" + cmd + "' to date: " + err.Error())
						date = time.Now()
					}
					order.expDate = date
				} else if len(dates) == 3 {
					date, err := time.Parse("1/2/2006", cmd)
					if err != nil {
						println("Error converting string date '" + cmd + "' to date: " + err.Error())
						date = time.Now()
					}
					order.expDate = date
				} else {
					println("Error converting string date '" + cmd + "' to date: Unknown format. Settig exp date for today")
					order.expDate = time.Now()
				}
				if order.expDate.Year() <= time.Now().Year() && order.expDate.YearDay() <= time.Now().YearDay() {
					order.risky = true
				}

			} else if strings.Contains(cmd, "p") && utils.IsNumericIgnore(cmd, "p", 1) || strings.Contains(cmd, "c") && utils.IsNumericIgnore(cmd, "c", 1) {
				order.strikPrice = cmd
			} else if strings.Contains(cmd, ".") && utils.IsNumericIgnore(cmd, "@", 1) {
				if order.price == 0 {
					order.price, err = utils.ToNumericIgnore(cmd, "@", 1)
					if err != nil {
						println("Error converting price '" + cmd + "' to float64: " + err.Error())
					}
				} else if order.stopLoss == 0 {
					order.stopLoss, err = utils.ToNumeric(cmd)
					if err != nil {
						println("Error converting stop loss '" + cmd + "' to float64: " + err.Error())
					}
				}
			}
		}
	}
	var emptyDate time.Time
	if order.expDate == emptyDate {
		order.expDate = time.Now()
		order.risky = true
	}
	if order.stopLoss == 0 {
		order.stopLoss = order.price * utils.Config.Settings.Trade.StopLossPercentage
	}

	ordersChannel <- order
}
