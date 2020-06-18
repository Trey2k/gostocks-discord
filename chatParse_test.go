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
		expected Order
	}{
		{"@everyone BTO COST 310c 4/9 @ 0.39 very risky", Order{true, "cost", time.Time{}, "310c", 0.39, true, 0.195, discordgo.User{}, "messageID"}},
		{"@everyone BTO SPY 4/9 266p @ .97", Order{true, "spy", time.Time{}, "266p", .97, true, 0.485, discordgo.User{}, "messageID"}},
		{"@everyone STC DRI 5/15/2021 60c @ 9.70\n\nALL/REMAINING (2) for 500% gain from 1.52 on Friday. This is a great example of a hedge to my SPY puts I got destroyed on this week.", Order{false, "dri", time.Time{}, "60c", 9.70, false, 4.85, discordgo.User{}, "messageID"}},
		{"@everyone BTO GO 10/16/2021 40C @0.9 (Starter Positions)", Order{true, "go", time.Time{}, "40c", 0.9, false, 0.45, discordgo.User{}, "messageID"}},
		{"@everyone BTO BYND 06/12/2021 170C @3.35 SL set at 3.00", Order{true, "bynd", time.Time{}, "170c", 3.35, false, 3.00, discordgo.User{}, "messageID"}},
		{"@everyone STC ROKU 120c 6/19/2021 @ 3.73 partial on target for 91%.", Order{false, "roku", time.Time{}, "120c", 3.73, false, 1.865, discordgo.User{}, "messageID"}},
		{"@everyone BTO AMZN 1/15/2021 3000c @ 67.95", Order{true, "amzn", time.Time{}, "3000c", 67.95, false, 33.975, discordgo.User{}, "messageID"}},
		{"@everyone BTO AAPL 340c 6/19/2021 @ 2.11", Order{true, "aapl", time.Time{}, "340c", 2.11, false, 1.055, discordgo.User{}, "messageID"}},
		{"@everyone STC SPY 06/05/2021 312C @ 6.3", Order{false, "spy", time.Time{}, "312c", 6.3, false, 3.15, discordgo.User{}, "messageID"}},
		{"@everyone BTO AMZN 2550C 6/5 @ 1.09 lotto", Order{true, "amzn", time.Time{}, "2550c", 1.09, true, 0.545, discordgo.User{}, "messageID"}},
		{"@everyone BTO SPY 06/03/2021 312P @0.42 LOTTO", Order{true, "spy", time.Time{}, "312p", 0.42, true, 0.21, discordgo.User{}, "messageID"}},
	}

	for _, test := range tests {
		go ChatParse(test.input, discordgo.User{}, "messageID")
		output := <-ordersChannel
		output.expDate = time.Time{}
		if output != test.expected {
			t.Error("Test failed, inputted: '" + test.input + "', expected: '" + fmt.Sprint(test.expected) + "', received: '" + fmt.Sprint(output) + "'")
		}
	}
}
