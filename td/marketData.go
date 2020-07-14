package td

import (
	"fmt"
	"strings"

	"github.com/Trey2k/gostocks-discord/utils"
)

//GetQuotesResponse Response struct for GetQuotes
type GetQuotesResponse struct {
	Quote map[string]GetQuoteResponse `json:"-"`
}

//GetQuoteResponse response struct for GetQuote
type GetQuoteResponse struct {
	AssetType                          string  `json:"assetType"`
	AssetMainType                      string  `json:"assetMainType"`
	Cusip                              string  `json:"cusip"`
	Symbol                             string  `json:"symbol"`
	Description                        string  `json:"description"`
	BidPrice                           float64 `json:"bidPrice"`
	BidSize                            int     `json:"bidSize"`
	BidID                              string  `json:"bidId"`
	AskPrice                           float64 `json:"askPrice"`
	AskSize                            int     `json:"askSize"`
	AskID                              string  `json:"askId"`
	LastPrice                          float64 `json:"lastPrice"`
	LastSize                           int     `json:"lastSize"`
	LastID                             string  `json:"lastId"`
	OpenPrice                          float64 `json:"openPrice"`
	HighPrice                          float64 `json:"highPrice"`
	LowPrice                           float64 `json:"lowPrice"`
	BidTick                            string  `json:"bidTick"`
	ClosePrice                         float64 `json:"closePrice"`
	NetChange                          float64 `json:"netChange"`
	TotalVolume                        float64 `json:"totalVolume"`
	QuoteTimeInLong                    int64   `json:"quoteTimeInLong"`
	TradeTimeInLong                    int64   `json:"tradeTimeInLong"`
	Mark                               float64 `json:"mark"`
	Exchange                           string  `json:"exchange"`
	ExchangeName                       string  `json:"exchangeName"`
	Marginable                         bool    `json:"marginable"`
	Shortable                          bool    `json:"shortable"`
	Volatility                         float64 `json:"volatility"`
	Digits                             int     `json:"digits"`
	Five2WkHigh                        float64 `json:"52WkHigh"`
	Five2WkLow                         float64 `json:"52WkLow"`
	NAV                                float64 `json:"nAV"`
	PeRatio                            float64 `json:"peRatio"`
	DivAmount                          float64 `json:"divAmount"`
	DivYield                           float64 `json:"divYield"`
	DivDate                            string  `json:"divDate"`
	SecurityStatus                     string  `json:"securityStatus"`
	RegularMarketLastPrice             float64 `json:"regularMarketLastPrice"`
	RegularMarketLastSize              int     `json:"regularMarketLastSize"`
	RegularMarketNetChange             float64 `json:"regularMarketNetChange"`
	RegularMarketTradeTimeInLong       int64   `json:"regularMarketTradeTimeInLong"`
	NetPercentChangeInDouble           float64 `json:"netPercentChangeInDouble"`
	MarkChangeInDouble                 float64 `json:"markChangeInDouble"`
	MarkPercentChangeInDouble          float64 `json:"markPercentChangeInDouble"`
	RegularMarketPercentChangeInDouble float64 `json:"regularMarketPercentChangeInDouble"`
	Delayed                            bool    `json:"delayed"`
}

//CHAINS

type Underlying struct {
	Symbol            string  `json:"symbol"`
	Description       string  `json:"description"`
	Change            float64 `json:"change"`
	PercentChange     float64 `json:"percentChange"`
	Close             float64 `json:"close"`
	QuoteTime         int     `json:"quoteTime"`
	TradeTime         int     `json:"tradeTime"`
	Bid               float64 `json:"bid"`
	Ask               float64 `json:"ask"`
	Last              float64 `json:"last"`
	Mark              float64 `json:"mark"`
	MarkChange        float64 `json:"markChange"`
	MarkPercentChange float64 `json:"markPercentChange"`
	BidSize           int     `json:"bidSize"`
	AskSize           int     `json:"askSize"`
	HighPrice         float64 `json:"highPrice"`
	LowPrice          float64 `json:"lowPrice"`
	OpenPrice         float64 `json:"openPrice"`
	TotalVolume       int     `json:"totalVolume"`
	ExchangeName      string  `json:"exchangeName"`
	FiftyTwoWeekHigh  float64 `json:"fiftyTwoWeekHigh"`
	FiftyTwoWeekLow   float64 `json:"fiftyTwoWeekLow"`
	Delayed           bool    `json:"delayed"`
}

type ExpDateOption struct {
	PutCall                string  `json:"putCall"`
	Symbol                 string  `json:"symbol"`
	Description            string  `json:"description"`
	ExchangeName           string  `json:"exchangeName"`
	Bid                    float64 `json:"bid"`
	Ask                    float64 `json:"ask"`
	Last                   float64 `json:"last"`
	Mark                   float64 `json:"mark"`
	BidSize                int     `json:"bidSize"`
	AskSize                int     `json:"askSize"`
	BidAskSize             string  `json:"bidAskSize"`
	LastSize               float64 `json:"lastSize"`
	HighPrice              float64 `json:"highPrice"`
	LowPrice               float64 `json:"lowPrice"`
	OpenPrice              float64 `json:"openPrice"`
	ClosePrice             float64 `json:"closePrice"`
	TotalVolume            int     `json:"totalVolume"`
	TradeDate              string  `json:"tradeDate"`
	TradeTimeInLong        int     `json:"tradeTimeInLong"`
	QuoteTimeInLong        int     `json:"quoteTimeInLong"`
	NetChange              float64 `json:"netChange"`
	Volatility             float64 `json:"volatility"`
	Delta                  float64 `json:"delta"`
	Gamma                  float64 `json:"gamma"`
	Vega                   float64 `json:"vega"`
	Rho                    float64 `json:"rho"`
	OpenInterest           int     `json:"openInterest"`
	TimeValue              float64 `json:"timeValue"`
	TheoreticalOptionValue float64 `json:"theoreticalOptionValue"`
	TheoreticalVolatility  float64 `json:"theoreticalVolatility"`
	OptionDeliverablesList string  `json:"optionDeliverablesList"`
	StrikePrice            float64 `json:"strikePrice"`
	ExpirationDate         int     `json:"expirationDate"`
	DaysToExpiration       int     `json:"daysToExpiration"`
	ExpirationType         string  `json:"expirationType"`
	LastTradingDate        int     `json:"lastTradingDay"`
	Multiplier             float64 `json:"multiplier"`
	SettlementType         string  `json:"settlementType"`
	DeliverableNote        string  `json:"deliverableNote"`
	IsIndexOption          bool    `json:"isIndexOption"`
	PercentChange          float64 `json:"percentChange"`
	MarkChange             float64 `json:"markChange"`
	MarkPercentChange      float64 `json:"markPercentChange"`
	InTheMoney             bool    `json:"inTheMoney"`
	Mini                   bool    `json:"mini"`
	NonStandard            bool    `json:"nonStandard"`
}

//ExpDateMap the first string is the exp date.  the second string is the strike price.
type ExpDateMap map[string]map[string][]ExpDateOption

//Chains stuff
type Chains struct {
	Symbol            string     `json:"symbol"`
	Status            string     `json:"status"`
	Underlying        Underlying `json:"underlying"`
	Strategy          string     `json:"strategy"`
	Interval          float64    `json:"interval"`
	IsDelayed         bool       `json:"isDelayed"`
	IsIndex           bool       `json:"isIndex"`
	InterestRate      float64    `json:"interestRate"`
	UnderlyingPrice   float64    `json:"underlyingPrice"`
	Volatility        float64    `json:"volatility"`
	DaysToExpiration  float64    `json:"daysToExpiration"`
	NumberOfContracts int        `json:"numberOfContracts"`
	CallExpDateMap    ExpDateMap `json:"callExpDateMap"`
	PutExpDateMap     ExpDateMap `json:"putExpDateMap"`
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
