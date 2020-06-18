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

//IsValidTicker test if string is a valid ticker
func IsValidTicker(ticker string) bool {
	if len(ticker) <= 5 && utils.NoNumbers(ticker) && ticker != "bto" && ticker != "stc" {
		var quoteResponse GetQuoteResponse

		err := GetQuote(ticker, &quoteResponse)
		utils.ErrCheck("Error testing is valid ticker for "+ticker, err)

		if quoteResponse.Symbol == strings.ToUpper(ticker) {
			return true
		}
	}
	return false
}
