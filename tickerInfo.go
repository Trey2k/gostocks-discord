package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type apiResponse struct {
	ResultSet struct {
		Query  string `json:"Query"`
		Result []struct {
			Symbol   string `json:"symbol"`
			Name     string `json:"name"`
			Exch     string `json:"exch"`
			Type     string `json:"type"`
			ExchDisp string `json:"exchDisp"`
			TypeDisp string `json:"typeDisp"`
		} `json:"Result"`
	} `json:"ResultSet"`
}

//TickerInfo is a json struct from yahoo finance
type TickerInfo struct {
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	Exch     string `json:"exch"`
	Type     string `json:"type"`
	ExchDisp string `json:"exchDisp"`
	TypeDisp string `json:"typeDisp"`
}

//GetTickerInfo returns a json struct from yahoo finance
func GetTickerInfo(ticker string) (TickerInfo, error) {
	var ticketInfo TickerInfo
	resp, err := http.Get("http://d.yimg.com/autoc.finance.yahoo.com/autoc?query=" + ticker + "&region=1&lang=en")
	if err != nil {
		return ticketInfo, err
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ticketInfo, err
	}

	// Convert response body to tickerInfor struct
	var response apiResponse
	json.Unmarshal(bodyBytes, &response)

	for i := 0; i < len(response.ResultSet.Result); i++ {
		if strings.ToLower(response.ResultSet.Result[i].Symbol) == strings.ToLower(ticker) {
			ticketInfo = response.ResultSet.Result[i]
		}
	}
	return ticketInfo, nil
}

//IsValidTicker test if string is a valid ticker
func IsValidTicker(s string) bool {
	if len(s) <= 5 {
		if noNumbers(s) {
			tikInf, err := GetTickerInfo(s)
			if err != nil {
				println("Error getting ticker info: " + err.Error())
				return false
			}
			if tikInf.Symbol != "" {
				return true
			}
		}
	}
	return false
}
