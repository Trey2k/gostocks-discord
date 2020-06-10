package main

import (
	"strings"
)

func chatParse(msg string, commands *[5]string) {
	if last := len(msg) - 1; last >= 0 && msg[last] == '.' {
		msg = msg[:last]
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
				} else if strings.Contains(cmd, ".") {
					commands[4] = strings.Replace(cmd, "@", "", 1)
					ci++
				}
			}
			if ci == 5 {
				return
			}
		}
	}
}
