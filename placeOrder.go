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
		fmt.Println("Error getting account info: " + errors.WithStack(err).Error())
	}

	tradeSettings := utils.Config.Settings.Trade

	var makeTrade bool = true
	var marketClosed bool = false
	var whitelistFail bool = false
	var expDateFail bool = false
	var completeFail bool = false
	var dataFound bool = false

	var optionData td.ExpDateOption

	tradeBalance := accountInfo.Account.CurrentBalances.CashAvailableForTrading
	initalBallance := accountInfo.Account.InitialBalances.CashBalance
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

	marketHours, err := td.GetMarketHours()
	if err != nil {
		fmt.Println("Error getting market hours: " + errors.WithStack(err).Error())
	}

	if marketHours.Option.EQO.IsOpen == false {
		makeTrade = false
		marketClosed = true
	} else {
		start, err := time.Parse("2006-01-02T15:04:05Z07:00", marketHours.Option.EQO.SessionHours.RegularMarket[0].Start)
		if err != nil {
			fmt.Println("Error parsing time: " + errors.WithStack(err).Error())
		}

		end, err := time.Parse("2006-01-02T15:04:05Z07:00", marketHours.Option.EQO.SessionHours.RegularMarket[0].End)
		if err != nil {
			fmt.Println("Error parsing time: " + errors.WithStack(err).Error())
		}

		if !utils.InTimeSpan(start, end, time.Now()) {
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
			fmt.Println("Error getting option data: " + errors.WithStack(err).Error())
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
			sell(order, optionData, tradeBalance)
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
		fmt.Println("Error querying db: " + errors.WithStack(err).Error())
	}
	if !aleadyOwn {
		tradeSettings := utils.Config.Settings.Trade
		contracts := int64((initalBallance * investPercent) / (optionData.Last * 100))
		if int(tradeBalance*100) >= int((initalBallance*investPercent)*100) && contracts != 0 {
			if int(optionData.Last*100) <= int(order.Price*100) || int((optionData.Last-order.Price)/optionData.Last*100) <= int(tradeSettings.AllowedPriceIncreasePercent*100) {

				mysql.NewOrder(order, optionData, contracts)

				totalPurchasePrice := float64(contracts) * (optionData.Last * 100)
				err = utils.SetTradeBal(tradeBalance - totalPurchasePrice)
				if err != nil {
					fmt.Println("Error Setting trade ball: " + errors.WithStack(err).Error())
				}

				payload := td.PlaceOrderBuyPayload{
					ComplexOrderStrategyType: "NONE",
					OrderType:                "LIMIT",
					Session:                  "NORMAL",
					Price:                    fmt.Sprint(float64(int((order.Price+(order.Price*tradeSettings.AllowedPriceIncreasePercent))*100)) / 100),
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
					fmt.Println(order.Message)
					utils.ErrCheck("BAD BAD BAD ERROR MAKEING ORDER: ", err)
				}

				fmt.Println("I made a order of " + fmt.Sprint(contracts) + " contracts at $" + fmt.Sprint(optionData.Last*100) + " or option price of $" + fmt.Sprint(optionData.Last) + " each for a total price of $" + fmt.Sprint(totalPurchasePrice))
				utils.PrintOrder(order)

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

func sell(order utils.OrderStruct, optionData td.ExpDateOption, tradeBalance float64) {
	owned, err := mysql.AlreadyOwn(optionData.Symbol)
	if err != nil {
		fmt.Println("Error querying db: " + errors.WithStack(err).Error())
	}
	if owned {
		resp, err := mysql.RetriveActiveOrder(optionData.Symbol)
		if err != nil {
			fmt.Println("Error querying db: " + errors.WithStack(err).Error())
		}

		if resp.Status != "PENDING" {

			err = mysql.SellContract(order)
			if err != nil {
				fmt.Println("Error querying db: " + errors.WithStack(err).Error())
			}

			sellPrice := float64(resp.Contracts) * optionData.Last
			purchasePrice := float64(resp.Contracts) * resp.PurchasePrice
			totalProfit := sellPrice - purchasePrice

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
				fmt.Println(order.Message)
				utils.ErrCheck("BAD BAD BAD ERROR SELLING ORDER: ", err)
			}

			fmt.Println("I sold " + fmt.Sprint(resp.Contracts) + " contracts at $" + fmt.Sprint(optionData.Last) + " each for a total of $" + fmt.Sprint(sellPrice*100))
			fmt.Println("The total purchase price was $" + fmt.Sprint(purchasePrice*100) + " that makes our total profit $" + fmt.Sprint(totalProfit*100))

			utils.PrintOrder(order)
		} else {
			fmt.Println("I do have some contracts for this option but they have not been filled so i cant sell them!")
			fmt.Println("Message: " + order.Message)
		}
	} else {
		fmt.Println("I do not own any contracts for this option")
		fmt.Println("Message: " + order.Message)
	}
}

func failLog(order utils.OrderStruct, failCode int, failMessage string) {
	fmt.Println("I did not make an order: " + failMessage)

	fmt.Println("Message: " + order.Message)

	err := mysql.FailedOrder(order, failCode, failMessage)
	if err != nil {
		fmt.Println("Error saving failedOrder in db: " + errors.WithStack(err).Error())
	}
}
