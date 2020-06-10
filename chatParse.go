package main

import (
	"strings"
)

//ChatParse : Parse a chat message and build a array of commands
func ChatParse(msg string) [7]string {
	var commands [7]string
	if last := len(msg) - 1; last >= 0 && msg[last] == '.' {
		msg = msg[:last]
	}
	if strings.Contains(strings.ToLower(msg), "risky") || strings.Contains(strings.ToLower(msg), "lotto") {
		commands[5] = "risky"
	} else {
		commands[5] = "safe"
	}
	msgs := strings.Split(msg, " ")
	ci := 0
	for i := 0; i < len(msgs); i++ {
		cmd := strings.ToLower(msgs[i])
		switch cmd {
		case "stc":
			commands[0] = "STC"
			ci++
		case "bto":
			commands[0] = "BTO"
			ci++
		default:
			if i == 2 {
				commands[1] = cmd
				ci++
			} else {
				if strings.Contains(cmd, "/") {
					if i <= 4 {
						commands[2] = cmd
						ci++
					}
				} else if strings.Contains(cmd, "p") || strings.Contains(cmd, "c") {
					if i <= 4 {
						commands[3] = cmd
						ci++
					}
				} else if strings.Contains(cmd, ".") && isNumeric(cmd) {
					if commands[4] == "" {
						commands[4] = strings.Replace(cmd, "@", "", 1)
						ci++
					} else if commands[6] == "" {
						commands[6] = cmd
						ci++
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
