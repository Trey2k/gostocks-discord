package td

//PlaceOrderPayload Payload for PlaceOrder
type PlaceOrderPayload struct {
	ComplexOrderStrategyType string `json:"complexOrderStrategyType"`
	OrderType                string `json:"orderType"`
	Session                  string `json:"session"`
	Price                    string `json:"price"`
	Duration                 string `json:"duration"`
	OrderStrategyType        string `json:"orderStrategyType"`
	OrderLegCollection       []struct {
		Instruction string `json:"instruction"`
		Quantity    int    `json:"quantity"`
		Instrument  struct {
			Symbol    string `json:"symbol"`
			AssetType string `json:"assetType"`
		} `json:"instrument"`
	} `json:"orderLegCollection"`
}
