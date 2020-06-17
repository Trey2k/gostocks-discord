package td

//GetAccounts returns all linked accounts
func GetAccounts(response *GetAccountResponses) error {

	err := getRequest("https://api.tdameritrade.com/v1/accounts", accessToken, &response)
	return err
}
