package main

import (
	"strings"

	"github.com/Trey2k/gostocks-discord/td"
	"github.com/Trey2k/gostocks-discord/utils"
)

//Commands list of commands built from messages
type Commands struct {
	buy        bool
	ticker     string
	expDate    string
	strikPrice string
	price      float64
	risky      bool
	stopLoss   float64
}

//ChatParse : Parse a chat message and build a array of commands
func ChatParse(msg string) {
	var err error
	var cmds Commands

	msg = strings.ToLower(strings.Split(msg, "\n")[0])

	cmds.risky = false
	if strings.Contains(msg, "risky") || strings.Contains(msg, "lotto") {
		cmds.risky = true
		msg = strings.Replace(msg, "riksy", "", 1)
		msg = strings.Replace(msg, "lotto", "", 1)
	}

	if strings.Contains(msg, "bto") {
		cmds.buy = true
		msg = strings.Replace(msg, "bto", "", 1)
	} else {
		cmds.buy = false
		msg = strings.Replace(msg, "stc", "", 1)
	}

	msgs := strings.Split(msg, " ")

	for i := 0; i < len(msgs); i++ {
		cmd := msgs[i]
		if i <= 5 && td.IsValidTicker(cmd) {
			cmds.ticker = cmd
		} else {
			if strings.Contains(cmd, "/") && utils.IsNumericIgnore(cmd, "/", 2) {
				cmds.expDate = cmd
			} else if strings.Contains(cmd, "p") && utils.IsNumericIgnore(cmd, "p", 1) || strings.Contains(cmd, "c") && utils.IsNumericIgnore(cmd, "c", 1) {
				cmds.strikPrice = cmd
			} else if strings.Contains(cmd, ".") && utils.IsNumericIgnore(cmd, "@", 1) {
				if cmds.price == 0 {
					cmds.price, err = utils.ToNumericIgnore(cmd, "@", 1)
					if err != nil {
						println("Error converting price '" + cmd + "' to float64: " + err.Error())
					}
				} else if cmds.stopLoss == 0 {
					cmds.stopLoss, err = utils.ToNumeric(cmd)
					if err != nil {
						println("Error converting stop loss '" + cmd + "' to float64: " + err.Error())
					}
				}
			}
		}
	}
	cmdsChannel <- cmds
}
