package main

import (
	"fmt"
	"time"

	"github.com/Trey2k/gostocks-discord/mysql"
	"github.com/Trey2k/gostocks-discord/td"
	"github.com/Trey2k/gostocks-discord/utils"
	"github.com/pkg/errors"
)

var marketHours td.MarketHours
var marketOpen time.Time
var marketClose time.Time

func updateMarketHours() {

	var err error

	marketHours, err = td.GetMarketHours()
	if err != nil {
		utils.Log("Getting market hours:"+errors.WithStack(err).Error(), utils.LogError)
	}
	if marketHours.Option.EQO.IsOpen == true {
		marketOpen, err = time.Parse("2006-01-02T15:04:05Z07:00", marketHours.Option.EQO.SessionHours.RegularMarket[0].Start)
		if err != nil {
			utils.Log("Parsing time:"+errors.WithStack(err).Error(), utils.LogError)
		}

		marketClose, err = time.Parse("2006-01-02T15:04:05Z07:00", marketHours.Option.EQO.SessionHours.RegularMarket[0].End)
		if err != nil {
			utils.Log("Parsing time:"+errors.WithStack(err).Error(), utils.LogError)
		}
	}
}

func update(period int) {
	for {
		var accountInfo td.GetAccountResponse
		err := td.GetAccount(utils.Config.TD.AccountID, &accountInfo)
		if err != nil {
			utils.Log("Getting account info: "+errors.WithStack(err).Error()+"\ntrying again next update", utils.LogError)
		} else {

			tradeBalance := accountInfo.Account.CurrentBalances.CashAvailableForTrading

			var update bool = true

			if marketHours.Option.EQO.IsOpen == false { //Checking if market is open
				if marketOpen.YearDay() != time.Now().YearDay() {
					updateMarketHours()
				}
				update = false
			} else {
				if !utils.InTimeSpan(marketOpen, marketClose, time.Now()) {
					if marketOpen.YearDay() != time.Now().YearDay() {
						updateMarketHours()
					}
					update = false
				}
			}

			if update {

				resp, err := mysql.GetOrders()
				if err != nil {
					utils.Log("Querying db: "+errors.WithStack(err).Error(), utils.LogError)
				}

				for i := 0; i < len(resp); i++ {
					optionData, dataFound, err := td.GetOptionData(resp[i].Order)
					if err != nil {
						utils.Log("Getting option data: "+errors.WithStack(err).Error(), utils.LogError)
					}

					if dataFound {
						purchasePrice := resp[i].PurchasePrice
						currentPrice := optionData.Last

						if resp[i].Status == "PENDING" {
							response, err := td.GetOrders(utils.Config.TD.AccountID, time.Now(), time.Now())
							if err != nil {
								utils.Log("Error getting list of orders from TD: "+errors.WithStack(err).Error(), utils.LogError)
							}

							for j := 0; j < len(response); j++ {
								if response[j].OrderLegCollection[0].Instrument.Symbol == optionData.Symbol && response[j].OrderLegCollection[0].Instruction == "BUY_TO_OPEN" {
									if response[j].Status == "FILLED" {
										err := mysql.OrderFilled(resp[i].Order, int(response[j].FilledQuantity))
										if err != nil {
											utils.Log("Updateing db:"+errors.WithStack(err).Error(), utils.LogError)
										}
										utils.Log("The order for "+resp[i].Symbol+" has been filled. The filled quantity is "+fmt.Sprint(response[j].FilledQuantity), utils.LogInfo)
										break
									}
								}
							}

						} else {
							if currentPrice > purchasePrice && int((currentPrice-purchasePrice)/purchasePrice*100) >= int(utils.Config.Settings.Trade.AutoSellProfitPercent*100) {
								utils.Log("Auto selling for profit baby!", utils.LogInfo)
								sell(resp[i].Order, optionData, tradeBalance)
							} else if currentPrice < purchasePrice && int(currentPrice*100) <= int(resp[i].Order.StopLoss*100) {
								utils.Log("Auto selling to save our ass!", utils.LogInfo)
								sell(resp[i].Order, optionData, tradeBalance)
							}
						}

					} else {
						utils.Log("Data could not be found for: "+resp[i].Symbol, utils.LogError)
					}
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
