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
	whitelistedUsers := tradeSettings.WhitelistUserIDs

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

	if order.Price == 0 || order.StrikPrice == 0 || order.ContractType == "" || order.Ticker == "" || order.ExpDate.IsZero() {
		makeTrade = false
		completeFail = true
	} else {
		optionData, dataFound, err = td.GetOptionData(order)
		utils.ErrCheck("Error getting Option Data", err)
	}

	if makeTrade && dataFound {
		if order.Buy {
			if order.Risky && tradeSettings.MakeRiskyTrades {
				riskyBuy(tradeBalance, initalBallance, riskyInvestPercent, order, optionData)
			} else if !order.Risky {
				safeBuy(tradeBalance, initalBallance, safeInvestPercent, order, optionData)
			} else {
				fmt.Println("I did not make a order, risky tradding is off")
			}
		} else {
			sell(order, optionData)
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

func riskyBuy(tradeBalance float64, initalBallance float64, riskyInvestPercent float64, order utils.OrderStruct, optionData td.ExpDateOption) {
	tradeSettings := utils.Config.Settings.Trade
	if tradeBalance >= initalBallance*riskyInvestPercent {
		if optionData.Last >= order.Price || int((order.Price-optionData.Last)/order.Price*100) <= int(tradeSettings.AllowedPriceIncreasePercent*100) {
			amount := int64((initalBallance * riskyInvestPercent) / optionData.Last)
			mysql.StoreOrder(order, optionData)
			fmt.Println("I made a risky order of " + fmt.Sprint(amount) + " contracts at the price of " + fmt.Sprint(optionData.Last) + " each")
		} else {
			fmt.Println("I did not make a order the price increase is greater than " + fmt.Sprint(int(tradeSettings.AllowedPriceIncreasePercent*100)) + "% at " + fmt.Sprint(int((order.Price-optionData.Last)/order.Price*100)) + "%")
		}
	} else {
		fmt.Println("I'm too poor for this trade")
	}
}

func safeBuy(tradeBalance float64, initalBallance float64, safeInvestPercent float64, order utils.OrderStruct, optionData td.ExpDateOption) {
	tradeSettings := utils.Config.Settings.Trade
	if tradeBalance >= initalBallance*safeInvestPercent {
		if optionData.Last <= order.Price || int((order.Price-optionData.Last)/order.Price*100) <= int(tradeSettings.AllowedPriceIncreasePercent*100) {
			amount := int64((initalBallance * safeInvestPercent) / optionData.Last)
			mysql.StoreOrder(order, optionData)
			fmt.Println("I made a risky order of " + fmt.Sprint(amount) + " contracts at the price of " + fmt.Sprint(optionData.Last) + " each")
		} else {
			fmt.Println("I did not make a order the price increase is greater than " + fmt.Sprint(int(tradeSettings.AllowedPriceIncreasePercent*100)) + "% at " + fmt.Sprint(int((order.Price-optionData.Last)/order.Price*100)) + "%")
		}
	} else {
		fmt.Println("I'm too poor for this trade")
	}
}

func sell(order utils.OrderStruct, optionData td.ExpDateOption) {
	//TODO: Code to find if order is in the DB
	fmt.Println("I made a sell")
}
