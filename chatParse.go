package main

import (
	"strings"
)

//Commands list of commands built from messages
type Commands struct {
	buysell    string
	ticker     string
	expDate    string
	strikPrice string
	price      float64
	danger     string
	stopLoss   float64
}

//ChatParse : Parse a chat message and build a array of commands
func ChatParse(msg string, cmds *Commands) {
	msg = strings.ToLower(msg)
	msgs := strings.Split(msg, " ")

	cmds.danger = "safe"
	if strings.Contains(msg, "risky") || strings.Contains(msg, "lotto") {
		cmds.danger = "risky"
	}

	for i := 0; i < len(msgs); i++ {
		cmd := msgs[i]
		switch cmd {
		case "stc":
			cmds.buysell = "STC"
		case "bto":
			cmds.buysell = "BTO"
		default:
			if i <= 4 && IsValidTicker(cmd) {
				cmds.ticker = cmd
			} else {
				if strings.Contains(cmd, "/") && isNumericIgnore(cmd, "/", 2) {
					cmds.expDate = cmd
				} else if strings.Contains(cmd, "p") && isNumericIgnore(cmd, "p", 1) || strings.Contains(cmd, "c") && isNumericIgnore(cmd, "c", 1) {
					cmds.strikPrice = cmd
				} else if strings.Contains(cmd, ".") && isNumericIgnore(cmd, "@", 1) {
					if cmds.price == 0 {
						var err error
						cmds.price, err = toNumericIgnore(cmd, "@", 1)
						errCheck("error converting price to float64", err)
					} else if cmds.stopLoss == 0 {
						var err error
						cmds.stopLoss, err = toNumeric(cmd)
						errCheck("error converting stop loss to float64", err)
					}
				}
			}
		}
	}
	if cmds.stopLoss == 0 {
		cmds.stopLoss = 0.00
	}
}
