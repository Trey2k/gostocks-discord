package td

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
/*func GetQuotes(symbol string) (Quote, error) {

	type resp struct {
		quote Quote `json:"AMZN"`
	}

	err := getRequest("https://api.tdameritrade.com/v1/marketdata/quotes?symbol="+symbol, accessToken, &resp)
	return resp.quote, err
}*/
