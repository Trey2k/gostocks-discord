package main

import (
	"fmt"
	"time"

	"github.com/Trey2k/gostocks-discord/td"
	"github.com/Trey2k/gostocks-discord/utils"
)

func placeOrder(order Order) {
	var accountInfo td.GetAccountResponse
	err := td.GetAccount(utils.Config.TD.AccountID, &accountInfo)
	utils.ErrCheck("Error getting Account Info", err)

	tradeSettings := utils.Config.Settings.Trade

	var makeTrade bool = true

	tradeBalance := accountInfo.SecuritiesAccount.CurrentBalances.CashAvailableForTrading
	initalBallance := accountInfo.SecuritiesAccount.InitialBalances.CashBalance
	riskyInvestPercent := tradeSettings.RiskyInvestPercentage
	safeInvestPercent := tradeSettings.SafeInvestPercentage
	useUserWhitelist := tradeSettings.UseUserWhitlist
	whitelistedUsers := utils.Config.Settings.Trade.WhitelistUserIDs

	if useUserWhitelist {
		makeTrade = false
		for i := 0; i < len(whitelistedUsers); i++ {
			if order.sender.ID == whitelistedUsers[i] {
				makeTrade = true
			}
		}
	}

	if order.expDate.YearDay() < time.Now().YearDay() || order.expDate.Year() < time.Now().Year() {
		makeTrade = false
	}

	if order.price == 0 || order.strikPrice == "" || order.ticker == "" {
		makeTrade = false
	}

	if makeTrade && order.risky && tradeSettings.MakeRiskyTrades {
		if tradeBalance >= initalBallance*riskyInvestPercent {
			fmt.Println("I made a risky order")
		} else {
			fmt.Println("I'm too poor for this trade")
		}
	} else if makeTrade && order.risky == false {
		if tradeBalance >= initalBallance*safeInvestPercent {
			fmt.Println("I made a safe order")
		} else {
			fmt.Println("I'm too poor for this trade")
		}
	} else {
		fmt.Println("I did not make a order")
	}

}
