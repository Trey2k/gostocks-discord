package td

import (
	"fmt"
	"strings"

	"github.com/Trey2k/gostocks-discord/utils"
)

//GetQuote a quote on a option or stock
func GetQuote(symbol string, response *GetQuoteResponse) error {
	var resp GetQuotesResponse
	symbol = strings.ToUpper(symbol)
	err := getRequest("https://api.tdameritrade.com/v1/marketdata/"+symbol+"/quotes", accessToken, &resp.Quote)
	*response = resp.Quote[symbol]
	return err
}

//GetQuotes a quote on a option or stock Symbols seperate via commas ie IBM,AMZN
func GetQuotes(symbols string, response *GetQuotesResponse) error {
	symbols = strings.ToUpper(symbols)
	err := getRequest("https://api.tdameritrade.com/v1/marketdata/quotes?symbol="+symbols, accessToken, &response.Quote)
	return err
}

func findData(order utils.OrderStruct, expMap ExpDateMap) (ExpDateOption, bool) {
	for date, temp := range expMap {
		date = strings.Split(date, ":")[0]
		if date == order.ExpDate.Format("2006-01-02") {
			if strings.Contains(fmt.Sprint(order.StrikPrice), ".") {
				return temp[fmt.Sprint(order.StrikPrice)][0], true
			} else {
				return temp[fmt.Sprint(order.StrikPrice)+".0"][0], true
			}
		}
	}
	return ExpDateOption{}, false
}

//GetOptionData stuff
func GetOptionData(order utils.OrderStruct) (ExpDateOption, bool, error) {
	var response Chains
	err := getRequest("https://api.tdameritrade.com/v1/marketdata/chains?symbol="+order.Ticker+"&contractType="+order.ContractType+"&strike="+fmt.Sprint(order.StrikPrice)+"&expMonth="+strings.ToUpper(order.ExpDate.Format("Jan")), accessToken, &response)
	if err != nil {
		return ExpDateOption{}, false, err
	}

	if order.ContractType == "CALL" {
		resp, found := findData(order, response.CallExpDateMap)
		return resp, found, nil
	}
	resp, found := findData(order, response.PutExpDateMap)
	return resp, found, nil
}

//GetMarketHours get the market hours
func GetMarketHours() (OptionHours, error) {
	var response MarketHoursResponse
	err := getRequest("https://api.tdameritrade.com/v1/marketdata/OPTION/hours", accessToken, &response)
	return response.Option.Option, err
}
