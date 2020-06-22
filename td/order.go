package td

import "time"

//GetOrdersResponse array of GetOrderResponse
type GetOrdersResponse []GetOrderResponse

//GetOrderResponse for GetOrder
type GetOrderResponse []struct {
	Session                  string                          `json:"session"`
	Duration                 string                          `json:"duration"`
	OrderType                string                          `json:"orderType"`
	ComplexOrderStrategyType string                          `json:"complexOrderStrategyType"`
	Quantity                 int                             `json:"quantity"`
	FilledQuantity           int                             `json:"filledQuantity"`
	RemainingQuantity        int                             `json:"remainingQuantity"`
	RequestedDestination     string                          `json:"requestedDestination"`
	DestinationLinkName      string                          `json:"destinationLinkName"`
	Price                    float64                         `json:"price"`
	OrderLegCollection       []OrderLegCollectionStruct      `json:"orderLegCollection"`
	OrderStrategyType        string                          `json:"orderStrategyType"`
	OrderID                  int64                           `json:"orderId"`
	Cancelable               bool                            `json:"cancelable"`
	Editable                 bool                            `json:"editable"`
	Status                   string                          `json:"status"`
	EnteredTime              string                          `json:"enteredTime"`
	CloseTime                string                          `json:"closeTime"`
	AccountID                int                             `json:"accountId"`
	OrderActivityCollection  []OrderActivityCollectionStruct `json:"orderActivityCollection"`
}

//OrderActivityCollectionStruct stuff
type OrderActivityCollectionStruct struct {
	ActivityType           string                `json:"activityType"`
	ExecutionType          string                `json:"executionType"`
	Quantity               int                   `json:"quantity"`
	OrderRemainingQuantity int                   `json:"orderRemainingQuantity"`
	ExecutionLegs          []ExecutionLegsStruct `json:"executionLegs"`
}

//OrderLegCollectionStruct stuff
type OrderLegCollectionStruct struct {
	OrderLegType   string           `json:"orderLegType"`
	LegID          int              `json:"legId"`
	Instrument     InstrumentStruct `json:"instrument"`
	Instruction    string           `json:"instruction"`
	PositionEffect string           `json:"positionEffect"`
	Quantity       int              `json:"quantity"`
}

//ExecutionLegsStruct Stuff
type ExecutionLegsStruct struct {
	LegID             int     `json:"legId"`
	Quantity          int     `json:"quantity"`
	MismarkedQuantity int     `json:"mismarkedQuantity"`
	Price             float64 `json:"price"`
	Time              string  `json:"time"`
}

//InstrumentStruct stuff
type InstrumentStruct struct {
	AssetType   string `json:"assetType"`
	Cusip       string `json:"cusip"`
	Symbol      string `json:"symbol"`
	Description string `json:"description"`
}

//PlaceOrderPayload Payload for PlaceOrder
type PlaceOrderPayload struct {
	ComplexOrderStrategyType string                      `json:"complexOrderStrategyType"`
	OrderType                string                      `json:"orderType"`
	Session                  string                      `json:"session"`
	Price                    string                      `json:"price"`
	Duration                 string                      `json:"duration"`
	OrderStrategyType        string                      `json:"orderStrategyType"`
	OrderLegCollection       []OrderLegCollectionPayload `json:"orderLegCollection"`
}

//OrderLegCollectionPayload Stuff
type OrderLegCollectionPayload struct {
	Instruction string            `json:"instruction"`
	Quantity    int               `json:"quantity"`
	Instrument  InstrumentPayload `json:"instrument"`
}

//InstrumentPayload Stuff
type InstrumentPayload struct {
	Symbol    string `json:"symbol"`
	AssetType string `json:"assetType"`
}

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
func GetOrders(accountID string, fromTime time.Time, toTime time.Time, response *GetOrdersResponse) error {
	var dateFormat string = "2006-01-02"
	err := getRequest("https://api.tdameritrade.com/v1/orders?accountId="+accountID+"&fromEnteredTime="+fromTime.Format(dateFormat)+"&toEnteredTime="+toTime.Format(dateFormat), accessToken, &response)
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
