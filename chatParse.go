package main

import (
	"strings"
)

//ChatParse : Parse a chat message and build a array of commands
func ChatParse(msg string) [7]string {
	var commands [7]string
	msg = strings.ToLower(msg)
	msgs := strings.Split(msg, " ")

	commands[5] = "safe"
	if strings.Contains(msg, "risky") || strings.Contains(msg, "lotto") {
		commands[5] = "risky"
	}

	for i := 0; i < len(msgs); i++ {
		cmd := msgs[i]
		switch cmd {
		case "stc":
			commands[0] = "STC"
		case "bto":
			commands[0] = "BTO"
		default:
			if i == 2 { //TODO: instead of basing ticker off of count check if cmd is equal to a valid ticker
				commands[1] = cmd
			} else {
				if strings.Contains(cmd, "/") && isNumericIgnore(cmd, "/", 2) {
					commands[2] = cmd
				} else if strings.Contains(cmd, "p") && isNumericIgnore(cmd, "p", 1) || strings.Contains(cmd, "c") && isNumericIgnore(cmd, "c", 1) {
					commands[3] = cmd
				} else if strings.Contains(cmd, ".") && isNumericIgnore(cmd, "@", 1) {
					if commands[4] == "" {
						commands[4] = strings.Replace(cmd, "@", "", 1)
					} else if commands[6] == "" {
						commands[6] = cmd
					}
				}
			}
		}
	}
	if commands[6] == "" {
		commands[6] = "0.00"
	}
	return commands
}
