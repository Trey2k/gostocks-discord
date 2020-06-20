package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/Trey2k/gostocks-discord/td"
	"github.com/Trey2k/gostocks-discord/utils"
	"github.com/bwmarrin/discordgo"
)

func init() {
	var err error
	utils.Config, err = utils.GetConfig()
	utils.ErrCheck("Error getting config", err)
	td.Init()

	td.Auth() //Holding call untill authed
}

func TestChatParse(t *testing.T) {

	var tests = []struct {
		input    string
		expected utils.OrderStruct
	}{
		{"@everyone BTO COST 310c 4/9 @ 0.39 very risky", utils.OrderStruct{true, true, "COST", time.Time{}, 310, "CALL", 0.39, 0.195, discordgo.User{}, "messageID", "message"}},
		{"@everyone BTO SPY 4/9 266p @ .97", utils.OrderStruct{true, true, "SPY", time.Time{}, 266, "PUT", .97, 0.485, discordgo.User{}, "messageID", "message"}},
		{"@everyone STC DRI 5/15/2021 60c @ 9.70\n\nALL/REMAINING (2) for 500% gain from 1.52 on Friday. This is a great example of a hedge to my SPY puts I got destroyed on this week.", utils.OrderStruct{false, false, "DRI", time.Time{}, 60, "CALL", 9.70, 4.85, discordgo.User{}, "messageID", "message"}},
		{"@everyone BTO GO 10/16/2021 40C @0.9 (Starter Positions)", utils.OrderStruct{true, false, "GO", time.Time{}, 40, "CALL", 0.9, 0.45, discordgo.User{}, "messageID", "message"}},
		{"@everyone BTO BYND 06/12/2021 170C @3.35 SL set at 3.00", utils.OrderStruct{true, false, "BYND", time.Time{}, 170, "CALL", 3.35, 3.00, discordgo.User{}, "messageID", "message"}},
		{"@everyone STC ROKU 120c 6/19/2021 @ 3.73 partial on target for 91%.", utils.OrderStruct{false, false, "ROKU", time.Time{}, 120, "CALL", 3.73, 1.865, discordgo.User{}, "messageID", "message"}},
		{"@everyone BTO AMZN 1/15/2021 3000c @ 67.95", utils.OrderStruct{true, false, "AMZN", time.Time{}, 3000, "CALL", 67.95, 33.975, discordgo.User{}, "messageID", "message"}},
		{"@everyone BTO AAPL 340c 6/19/2021 @ 2.11", utils.OrderStruct{true, false, "AAPL", time.Time{}, 340, "CALL", 2.11, 1.055, discordgo.User{}, "messageID", "message"}},
		{"@everyone STC SPY 06/05/2021 312C @ 6.3", utils.OrderStruct{false, false, "SPY", time.Time{}, 312, "CALL", 6.3, 3.15, discordgo.User{}, "messageID", "message"}},
		{"@everyone BTO AMZN 2550C 6/5 @ 1.09 lotto", utils.OrderStruct{true, true, "AMZN", time.Time{}, 2550, "CALL", 1.09, 0.545, discordgo.User{}, "messageID", "message"}},
		{"@everyone BTO SPY 06/03/2021 312P @0.42 LOTTO", utils.OrderStruct{true, true, "SPY", time.Time{}, 312, "PUT", 0.42, 0.21, discordgo.User{}, "messageID", "message"}},
	}

	for _, test := range tests {
		output := ChatParse(test.input, discordgo.User{}, "messageID")
		output.ExpDate = time.Time{}
		output.Message = "message"
		if output != test.expected {
			t.Error("Test failed, inputted: '" + test.input + "', expected: '" + fmt.Sprint(test.expected) + "', received: '" + fmt.Sprint(output) + "'")
		}
	}
}
