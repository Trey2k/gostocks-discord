package main

import (
	"fmt"
	"time"

	"github.com/Trey2k/gostocks-discord/mysql"
	"github.com/Trey2k/gostocks-discord/td"
	"github.com/Trey2k/gostocks-discord/utils"
	"github.com/pkg/errors"
)

func placeOrder(order utils.OrderStruct) {
	var accountInfo td.GetAccountResponse
	err := td.GetAccount(utils.Config.TD.AccountID, &accountInfo)
	if err != nil {
		utils.Log("Getting account info: "+errors.WithStack(err).Error(), utils.LogError)
	}

	tradeSettings := utils.Config.Settings.Trade

	var makeTrade bool = true
	var marketClosed bool = false
	var whitelistFail bool = false
	var expDateFail bool = false
	var completeFail bool = false
	var dataFound bool = false

	var optionData td.ExpDateOption

	tradeBalance := accountInfo.Account.CurrentBalances.CashAvailableForTrading - accountInfo.Account.CurrentBalances.PendingDeposits
	initalBallance := accountInfo.Account.InitialBalances.CashAvailableForTrading - accountInfo.Account.InitialBalances.PendingDeposits
	riskyInvestPercent := tradeSettings.RiskyInvestPercent
	safeInvestPercent := tradeSettings.SafeInvestPercent
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

	if marketHours.Option.EQO.IsOpen == false { //Checking if market is open
		if marketOpen.YearDay() != time.Now().YearDay() {
			updateMarketHours()
		}
		makeTrade = false
		marketClosed = true
	} else {
		if !utils.InTimeSpan(marketOpen, marketClose, time.Now()) {
			if marketOpen.YearDay() != time.Now().YearDay() {
				updateMarketHours()
			}
			makeTrade = false
			marketClosed = true
		}
	}

	if order.Price == 0 || order.StrikPrice == 0 || order.ContractType == "" || order.Ticker == "" || order.ExpDate.IsZero() {
		makeTrade = false
		completeFail = true
	} else {
		optionData, dataFound, err = td.GetOptionData(order)
		if err != nil {
			utils.Log("Getting option data: "+errors.WithStack(err).Error(), utils.LogError)
		}
	}

	if order.ExpDate.YearDay() < time.Now().YearDay() || order.ExpDate.Year() < time.Now().Year() {
		makeTrade = false
		expDateFail = true
	}

	if makeTrade && dataFound {
		if order.Buy {
			if order.Risky && tradeSettings.MakeRiskyTrades {
				buy(tradeBalance, initalBallance, riskyInvestPercent, order, optionData)
			} else if !order.Risky {
				buy(tradeBalance, initalBallance, safeInvestPercent, order, optionData)
			} else { //105
				failMessage := "Risky tradding is not enabled."
				failLog(order, 105, failMessage)
			}
		} else {
			sell(order, optionData)
		}
	} else {
		if marketClosed { //100
			failMessage := "The market is currently closed."
			failLog(order, 100, failMessage)
		} else if whitelistFail { //101
			failMessage := "Sender it not in whitelist."
			failLog(order, 101, failMessage)
		} else if completeFail { //103
			failMessage := "Could not find enough instructions."
			failLog(order, 102, failMessage)
		} else if expDateFail { //102
			failMessage := "Contract is expired."
			failLog(order, 103, failMessage)
		} else if !dataFound { //104
			failMessage := "Could not find this contract."
			failLog(order, 104, failMessage)
		}
	}
}

func buy(tradeBalance float64, initalBallance float64, investPercent float64, order utils.OrderStruct, optionData td.ExpDateOption) {
	aleadyOwn, err := mysql.AlreadyOwn(optionData.Symbol)
	if err != nil {
		utils.Log("Querying db: "+errors.WithStack(err).Error(), utils.LogError)
	}
	if !aleadyOwn {
		tradeSettings := utils.Config.Settings.Trade
		multiplier := optionData.Multiplier

		contracts := int64((initalBallance * investPercent) / (optionData.Last * multiplier))
		if int(tradeBalance*100) >= int((initalBallance*investPercent)*100) && contracts != 0 {
			if int(optionData.Last*100) <= int(order.Price*100) || int((optionData.Last-order.Price)/optionData.Last*100) <= int(tradeSettings.AllowedPriceIncreasePercent*100) {

				mysql.NewOrder(order, optionData, contracts)

				totalPurchasePrice := float64((int(contracts) * int((optionData.Last*multiplier)*100)) / 100)

				if utils.Config.Settings.Trade.MakeOrders {
					payload := td.PlaceOrderBuyPayload{
						ComplexOrderStrategyType: "NONE",
						OrderType:                "LIMIT",
						Session:                  "NORMAL",
						Price:                    fmt.Sprint(float64(int((optionData.Last+(optionData.Last*tradeSettings.AllowedPriceIncreasePercent))*100)) / 100), //Normiliing for real price
						Duration:                 "DAY",
						OrderStrategyType:        "SINGLE",
						OrderLegCollection: []td.OrderLegCollectionPayload{
							{
								Instruction: "BUY_TO_OPEN",
								Quantity:    int(contracts),
								Instrument: td.InstrumentPayload{
									Symbol:    optionData.Symbol,
									AssetType: "OPTION",
								},
							},
						},
					}

					err := td.PlaceOrderBuy(utils.Config.TD.AccountID, payload)
					if err != nil {
						utils.Log("BAD BAD BAD I FAILED TO MAKE THIS ORDER THROUGH TD\n"+order.Message+"\n"+errors.WithStack(err).Error(), utils.LogError)
					}
				}

				utils.Log("I made a order of "+fmt.Sprint(contracts)+" contracts at $"+fmt.Sprint(float64(int((optionData.Last*multiplier)*100)/100))+" or option price of $"+fmt.Sprint(optionData.Last)+" each for a total price of $"+fmt.Sprint(totalPurchasePrice)+"\n"+utils.PrintOrder(order), utils.LogOrder)

			} else { //106
				failMessage := "The price increase is greater than " + fmt.Sprint(int(tradeSettings.AllowedPriceIncreasePercent*100)) + "% at " + fmt.Sprint(int((optionData.Last-order.Price)/optionData.Last*100)) + "%"
				failLog(order, 106, failMessage)
			}
		} else { //107
			failMessage := "I do not have enough trading funds to make this trade."
			failLog(order, 107, failMessage)
		}
	} else { //108
		failMessage := "I already own contracts for this option."
		failLog(order, 108, failMessage)
	}
}

func sell(order utils.OrderStruct, optionData td.ExpDateOption) {
	owned, err := mysql.AlreadyOwn(optionData.Symbol)
	if err != nil {
		utils.Log("Querying db: "+errors.WithStack(err).Error(), utils.LogError)
	}
	if owned {
		multiplier := optionData.Multiplier
		resp, err := mysql.RetriveActiveOrder(optionData.Symbol)
		if err != nil {
			utils.Log("Querying db: "+errors.WithStack(err).Error(), utils.LogError)
		}

		if resp.Status != "PENDING" {

			err = mysql.SellContract(order)
			if err != nil {
				utils.Log("Querying db: "+errors.WithStack(err).Error(), utils.LogError)
			}

			sellPrice := float64(resp.Contracts) * optionData.Last
			purchasePrice := float64(resp.Contracts) * resp.PurchasePrice
			totalProfit := sellPrice - purchasePrice

			if utils.Config.Settings.Trade.MakeOrders {
				payload := td.PlaceOrderSellPayload{
					ComplexOrderStrategyType: "NONE",
					OrderType:                "MARKET",
					Session:                  "NORMAL",
					Duration:                 "DAY",
					OrderStrategyType:        "SINGLE",
					OrderLegCollection: []td.OrderLegCollectionPayload{
						{
							Instruction: "SELL_TO_CLOSE",
							Quantity:    resp.Contracts,
							Instrument: td.InstrumentPayload{
								Symbol:    optionData.Symbol,
								AssetType: "OPTION",
							},
						},
					},
				}

				err = td.PlaceOrderSell(utils.Config.TD.AccountID, payload)
				if err != nil {
					utils.Log("BAD BAD BAD I FAILED TO MAKE THIS ORDER THROUGH TD\n"+order.Message+"\n"+errors.WithStack(err).Error(), utils.LogError)
				}
			}

			utils.Log("I sold "+fmt.Sprint(resp.Contracts)+" contracts at $"+fmt.Sprint(optionData.Last)+" each for a total of $"+fmt.Sprint(float64(int(sellPrice*multiplier)*100)/100)+
				"\nThe total purchase price was $"+fmt.Sprint(float64(int(purchasePrice*multiplier)*100)/100)+" that makes our total profit $"+fmt.Sprint(float64(int(totalProfit*multiplier)*100)/100), utils.LogOrder)

			utils.PrintOrder(order)
		} else {
			utils.Log("I do have some contracts for this option but they have not been filled so i cant sell them!\nMessage: "+order.Message, utils.LogError)
		}
	} else {
		utils.Log("I do not own any contracts for this option\nMessage: "+order.Message, utils.LogError)
	}
}

func failLog(order utils.OrderStruct, failCode int, failMessage string) {
	utils.Log("Order could not be made. "+failMessage+"\nMessage: "+order.Message, utils.LogError)

	err := mysql.FailedOrder(order, failCode, failMessage)
	if err != nil {
		utils.Log("Saving failedOrder in db: "+errors.WithStack(err).Error(), utils.LogError)
	}
}
