package main

import (
	"fmt"
	"time"

	"github.com/Trey2k/gostocks-discord/mysql"
	"github.com/Trey2k/gostocks-discord/td"
	"github.com/Trey2k/gostocks-discord/utils"
	"github.com/pkg/errors"
)

func update(period int) {
	for {
		var accountInfo td.GetAccountResponse
		err := td.GetAccount(utils.Config.TD.AccountID, &accountInfo)
		if err != nil {
			fmt.Println("Error getting account info: " + errors.WithStack(err).Error())
		}

		tradeBalance, err := utils.GetTradeBal(accountInfo.Account.CurrentBalances.CashAvailableForTrading)
		if err != nil {
			fmt.Println("Error getting trade ballance: " + errors.WithStack(err).Error())
		}

		var update bool = true

		marketHours, err := td.GetMarketHours()
		if err != nil {
			fmt.Println("Error getting market hours: " + errors.WithStack(err).Error())
		}

		if marketHours.Option.EQO.IsOpen == false {
			update = false
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
				update = false
			}
		}

		if update {

			resp, err := mysql.GetOrders()
			if err != nil {
				fmt.Println("Error querying db: " + errors.WithStack(err).Error())
			}

			for i := 0; i < len(resp); i++ {
				optionData, dataFound, err := td.GetOptionData(resp[i].Order)
				if err != nil {
					fmt.Println("Error getting option data: " + errors.WithStack(err).Error())
				}

				if dataFound {
					purchasePrice := resp[i].PurchasePrice
					currentPrice := optionData.Last

					if currentPrice > purchasePrice && int((currentPrice-purchasePrice)/purchasePrice*100) >= int(utils.Config.Settings.Trade.AutoSellProfitPercent*100) {
						sell(resp[i].Order, optionData, tradeBalance)
					} else if currentPrice < purchasePrice && int(currentPrice*100) <= int(resp[i].Order.StopLoss*100) {
						sell(resp[i].Order, optionData, tradeBalance)
					}

				} else {
					fmt.Println("Error data could not be found for: " + resp[i].Symbol)
				}
			}
		}
		time.Sleep(time.Second * time.Duration(period))
	}
}

func procOrder(ordersChan chan utils.OrderStruct) {
	for {
		placeOrder(<-ordersChan)
	}
}
