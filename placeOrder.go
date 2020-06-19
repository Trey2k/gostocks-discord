package main

import (
	"fmt"
	"time"

	"github.com/Trey2k/gostocks-discord/mysql"
	"github.com/Trey2k/gostocks-discord/td"
	"github.com/Trey2k/gostocks-discord/utils"
)

func placeOrder(order utils.OrderStruct) {
	var accountInfo td.GetAccountResponse
	err := td.GetAccount(utils.Config.TD.AccountID, &accountInfo)
	utils.ErrCheck("Error getting Account Info", err)

	tradeSettings := utils.Config.Settings.Trade

	var makeTrade bool = true
	var whitelistFail bool = false
	var expDateFail bool = false
	var completeFail bool = false
	var dataFound bool = false

	var optionData td.ExpDateOption

	tradeBalance := accountInfo.SecuritiesAccount.CurrentBalances.CashAvailableForTrading
	initalBallance := accountInfo.SecuritiesAccount.InitialBalances.CashBalance
	riskyInvestPercent := tradeSettings.RiskyInvestPercentage
	safeInvestPercent := tradeSettings.SafeInvestPercentage
	useUserWhitelist := tradeSettings.UseUserWhitlist
	whitelistedUsers := utils.Config.Settings.Trade.WhitelistUserIDs

	if useUserWhitelist {
		makeTrade = false
		whitelistFail = true
		for i := 0; i < len(whitelistedUsers); i++ {
			if order.Sender.ID == whitelistedUsers[i] {
				makeTrade = true
				whitelistFail = false
			}
		}
	}

	if order.ExpDate.YearDay() < time.Now().YearDay() || order.ExpDate.Year() < time.Now().Year() {
		makeTrade = false
		expDateFail = true
	}

	if order.Price == 0 || order.StrikPrice == 0 || order.ContractType == "" || order.Ticker == "" {
		makeTrade = false
		completeFail = true
	} else {
		optionData, dataFound, err = td.GetOptionData(order)
		utils.ErrCheck("Error getting Option Data", err)
	}

	if makeTrade && order.Risky && tradeSettings.MakeRiskyTrades && dataFound {
		if tradeBalance >= initalBallance*riskyInvestPercent {
			mysql.StoreOrder(order, optionData)
			fmt.Println("I made a risky order")

		} else {
			fmt.Println("I'm too poor for this trade")
		}
	} else if makeTrade && order.Risky == false && dataFound {
		if tradeBalance >= initalBallance*safeInvestPercent {
			mysql.StoreOrder(order, optionData)
			fmt.Println("I made a safe order")

		} else {
			fmt.Println("I'm too poor for this trade")
		}
	} else {
		if whitelistFail {
			fmt.Println("I did not make a order, sender not in whitelist")
		} else if expDateFail {
			fmt.Println("I did not make a order, trade expired")
		} else if completeFail {
			fmt.Println("I did not make a order, not enough instructions")
		} else if !dataFound {
			fmt.Println("I did not make a order, could not find option data")
		} else {
			fmt.Println("I did not make a order")
		}
	}
}
