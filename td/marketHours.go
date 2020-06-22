package td

//MarketHourse Stuff
type MarketHourse struct {
	Option OptionHours `json:"option"`
}

//OptionHours stuff
type OptionHours struct {
	EQO Hours `json:"EQO"`
	IND Hours `json:"IND"`
}

//Hours stuff
type Hours struct {
	Date         string             `json:"date"`
	MarketType   string             `json:"marketType"`
	Exchange     string             `json:"exchange"`
	Category     string             `json:"category"`
	Product      string             `json:"product"`
	ProductName  string             `json:"productName"`
	IsOpen       bool               `json:"isOpen"`
	SessionHours SessionHoursStruct `json:"sessionHours"`
}

//SessionHoursStruct Stuff
type SessionHoursStruct struct {
	RegularMarket []MarketSession `json:"regularMarket"`
}

//MarketSession Stuff
type MarketSession struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

//GetMarketHours get the market hours
func GetMarketHours() (MarketHourse, error) {
	var response MarketHourse
	err := getRequest("https://api.tdameritrade.com/v1/marketdata/OPTION/hours", accessToken, &response)
	return response, err
}
