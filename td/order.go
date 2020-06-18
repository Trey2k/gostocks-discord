package td

//CancleOrder Cancle an order
func CancleOrder(accountID string, orderID string) error {

	err := deleteRequest("https://api.tdameritrade.com/v1/accounts/"+accountID+"/orders/"+orderID, accessToken)
	return err
}

//GetOrder Returns an order
func GetOrder(accountID string, orderID string, response *GetOrderResponse) error {

	err := getRequest("https://api.tdameritrade.com/v1/accounts/"+accountID+"/orders/"+orderID, accessToken, &response)
	return err
}

//GetOrders Returns all orders
func GetOrders(accountID string, response *GetOrderResponses) error {

	err := getRequest("https://api.tdameritrade.com/v1/accounts/"+accountID+"/orders/", accessToken, &response)
	return err
}

//PlaceOrder place an order
func PlaceOrder(accountID string, payload PlaceOrderPayload) error {

	err := postRequest("https://api.tdameritrade.com/v1/accounts/"+accountID+"/orders/", accessToken, payload)
	return err
}

//ReplaceOrder repalce an order
func ReplaceOrder(accountID string, orderID string, payload PlaceOrderPayload) error {

	err := putRequest("https://api.tdameritrade.com/v1/accounts/"+accountID+"/orders/"+orderID, accessToken, payload)
	return err
}
