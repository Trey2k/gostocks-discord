package main

import (
	"fmt"
	"testing"
)

func TestChatParse(t *testing.T) {
	var tests = []struct {
		input    string
		expected Commands
	}{
		{"@everyone BTO COST 310c 4/9 @ 0.39 very risky", Commands{true, "cost", "4/9", "310c", 0.39, true, 0.00}},
		{"@everyone BTO SPY 4/9 266p @ .97", Commands{true, "spy", "4/9", "266p", .97, false, 0.00}},
		{"@everyone STC DRI 5/15 60c @ 9.70\n\nALL/REMAINING (2) for 500% gain from 1.52 on Friday. This is a great example of a hedge to my SPY puts I got destroyed on this week.", Commands{false, "dri", "5/15", "60c", 9.70, false, 0.00}},
		{"@everyone BTO GO 10/16 40C @0.9 (Starter Positions)", Commands{true, "go", "10/16", "40c", 0.9, false, 0.00}},
		{"@everyone BTO BYND 06/12 170C @3.35 SL set at 3.00", Commands{true, "bynd", "06/12", "170c", 3.35, false, 3.00}},
		{"@everyone STC ROKU 120c 6/19 @ 3.73 partial on target for 91%.", Commands{false, "roku", "6/19", "120c", 3.73, false, 0.00}},
		{"@everyone BTO AMZN 1/15/2021 3000c @ 67.95", Commands{true, "amzn", "1/15/2021", "3000c", 67.95, false, 0.00}},
		{"@everyone BTO AAPL 340c 6/19 @ 2.11", Commands{true, "aapl", "6/19", "340c", 2.11, false, 0.00}},
		{"@everyone STC SPY 06/05 312C @ 6.3", Commands{false, "spy", "06/05", "312c", 6.3, false, 0.00}},
		{"@everyone BTO AMZN 2550C 6/5 @ 1.09 lotto", Commands{true, "amzn", "6/5", "2550c", 1.09, true, 0.00}},
		{"@everyone BTO SPY 06/03 312P @0.42 LOTTO", Commands{true, "spy", "06/03", "312p", 0.42, true, 0.00}},
	}

	for _, test := range tests {
		go ChatParse(test.input)
		output := <-channel
		if output != test.expected {
			t.Error("Test failed, inputted: '" + test.input + "', expected: '" + fmt.Sprint(test.expected) + "', received: '" + fmt.Sprint(output) + "'")
		}
	}
}
