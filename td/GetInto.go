package td

import (
	"strings"

	"github.com/Trey2k/gostocks-discord/utils"
)

//GetAccounts returns all linked accounts
func GetAccounts(response *GetAccountResponses) error {

	err := getRequest("https://api.tdameritrade.com/v1/accounts", accessToken, &response)
	return err
}

//GetAccount returns specific account
func GetAccount(accountID string, response *GetAccountResponse) error {

	err := getRequest("https://api.tdameritrade.com/v1/accounts/"+accountID, accessToken, &response)
	return err
}

//GetQuotes a quote on a option or stock
func GetQuotes(symbol string, response *GetQuoteResponse) error {
	var resp GetQuoteTemp
	symbol = strings.ToUpper(symbol)
	err := getRequest("https://api.tdameritrade.com/v1/marketdata/quotes?symbol="+symbol, accessToken, &resp.Quote)
	*response = resp.Quote[symbol]
	return err
}

//IsValidTicker test if string is a valid ticker
func IsValidTicker(ticker string) bool {
	var quoteResponse GetQuoteResponse

	err := GetQuotes(ticker, &quoteResponse)
	utils.ErrCheck("Error testing is valid ticker for "+ticker, err)

	if quoteResponse.Symbol == strings.ToUpper(ticker) {
		return true
	}
	return false
}
