package main

func getAccounts(token string, response *GetAccountResponses) error {

	err := getRequest("https://api.tdameritrade.com/v1/accounts", token, &response)
	return err
}
