package main

import (
	"time"

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
		select {
		case c := <-channel:
			requestIB(client, c)
		default:
			time.Sleep(time.Second * 1)
		}
	}
}

func requestIB(client *ib.IbClient, commands Commands) {

}
