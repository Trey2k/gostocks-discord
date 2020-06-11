package main

import (
	"fmt"

	ib "github.com/hadrianl/ibapi"
)

func connectIB(channel chan Commands) {
	var err error
	log := ib.GetLogger().Sugar()
	defer log.Sync()
	ibwrapper := &ib.Wrapper{}
	client := ib.NewIbClient(ibwrapper)
	err = client.Connect(config.IB.IP, config.IB.Port, config.IB.ClientID)
	if err != nil {
		log.Panic("Connect failed:", err)
		return
	}

	err = client.HandShake()
	if err != nil {
		log.Info("HandShake failed:", err)
		return
	}

	client.Run()
	for {
		requestIB(client, <-channel)
	}
}

func requestIB(client *ib.IbClient, commands Commands) {
	fmt.Println("------------------------------------------------------------------------------------------------------------")
	fmt.Println("Buy/Sell: " + commands.buysell + ", Ticker: " + commands.ticker + ", ExpDate: " + commands.expDate + ", StrikerPrice: " + commands.strikPrice + ", Buy Price: " + fmt.Sprint(commands.price) + ", Danger: " + commands.danger + ", Stop Loss: " + fmt.Sprint(commands.stopLoss))
	fmt.Println("------------------------------------------------------------------------------------------------------------")
}
