package main

import (
	"strings"

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
	msgs := strings.Split(msg, " ")

	cmds.risky = false
	if strings.Contains(msg, "risky") || strings.Contains(msg, "lotto") {
		cmds.risky = true
	}

	cmds.buy = false
	if strings.Contains(msg, "bto") {
		cmds.buy = true
	}

	for i := 0; i < len(msgs); i++ {
		cmd := msgs[i]
		if i <= 4 && IsValidTicker(cmd) {
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
	channel <- cmds
}
