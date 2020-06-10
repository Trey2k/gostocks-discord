package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//TickerInfo is a json struct from yahoo finance
type TickerInfo struct {
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

//GetTickerInfo returns a json struct from yahoo finance
func GetTickerInfo(ticker string) TickerInfo {
	resp, err := http.Get("http://d.yimg.com/autoc.finance.yahoo.com/autoc?query=" + ticker + "&region=1&lang=en")
	errCheck("Error connecting to yahoo REST api", err)

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to tickerInfor struct
	var response TickerInfo
	json.Unmarshal(bodyBytes, &response)
	return response
}

//IsValidTicker test if string is a valid ticker
func IsValidTicker(s string) bool {
	if len(s) <= 5 {
		if isLetter(s) {
			tikInf := GetTickerInfo(s)
			if len(tikInf.ResultSet.Result) >= 1 {
				return true
			}
		}
	}
	return false
}
