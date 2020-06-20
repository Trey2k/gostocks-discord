package td

//RequestTokensResponse response struct for token request
type RequestTokensResponse struct {
	AccessToken           string `json:"access_token"`
	RefreshToken          string `json:"refresh_token"`
	Scope                 string `json:"scope"`
	ExpiresIn             int    `json:"expires_in"`
	RefreshTokenExpiresIn int    `json:"refresh_token_expires_in"`
	TokenType             string `json:"token_type"`
}

//GetAccountResponses array of GetAccountResponse
type GetAccountResponses []GetAccountResponse

//GetAccountResponse struct
type GetAccountResponse struct {
	SecuritiesAccount struct {
		Type                    string `json:"type"`
		AccountID               string `json:"accountId"`
		RoundTrips              int    `json:"roundTrips"`
		IsDayTrader             bool   `json:"isDayTrader"`
		IsClosingOnlyRestricted bool   `json:"isClosingOnlyRestricted"`
		InitialBalances         struct {
			AccruedInterest            float64 `json:"accruedInterest"`
			CashAvailableForTrading    float64 `json:"cashAvailableForTrading"`
			CashAvailableForWithdrawal float64 `json:"cashAvailableForWithdrawal"`
			CashBalance                float64 `json:"cashBalance"`
			BondValue                  float64 `json:"bondValue"`
			CashReceipts               float64 `json:"cashReceipts"`
			LiquidationValue           float64 `json:"liquidationValue"`
			LongOptionMarketValue      float64 `json:"longOptionMarketValue"`
			LongStockValue             float64 `json:"longStockValue"`
			MoneyMarketFund            float64 `json:"moneyMarketFund"`
			MutualFundValue            float64 `json:"mutualFundValue"`
			ShortOptionMarketValue     float64 `json:"shortOptionMarketValue"`
			ShortStockValue            float64 `json:"shortStockValue"`
			IsInCall                   bool    `json:"isInCall"`
			UnsettledCash              float64 `json:"unsettledCash"`
			CashDebitCallValue         float64 `json:"cashDebitCallValue"`
			PendingDeposits            float64 `json:"pendingDeposits"`
			AccountValue               float64 `json:"accountValue"`
		} `json:"initialBalances"`
		CurrentBalances struct {
			AccruedInterest              float64 `json:"accruedInterest"`
			CashBalance                  float64 `json:"cashBalance"`
			CashReceipts                 float64 `json:"cashReceipts"`
			LongOptionMarketValue        float64 `json:"longOptionMarketValue"`
			LiquidationValue             float64 `json:"liquidationValue"`
			LongMarketValue              float64 `json:"longMarketValue"`
			MoneyMarketFund              float64 `json:"moneyMarketFund"`
			Savings                      float64 `json:"savings"`
			ShortMarketValue             float64 `json:"shortMarketValue"`
			PendingDeposits              float64 `json:"pendingDeposits"`
			CashAvailableForTrading      float64 `json:"cashAvailableForTrading"`
			CashAvailableForWithdrawal   float64 `json:"cashAvailableForWithdrawal"`
			CashCall                     float64 `json:"cashCall"`
			LongNonMarginableMarketValue float64 `json:"longNonMarginableMarketValue"`
			TotalCash                    float64 `json:"totalCash"`
			ShortOptionMarketValue       float64 `json:"shortOptionMarketValue"`
			MutualFundValue              float64 `json:"mutualFundValue"`
			BondValue                    float64 `json:"bondValue"`
			CashDebitCallValue           float64 `json:"cashDebitCallValue"`
			UnsettledCash                float64 `json:"unsettledCash"`
		} `json:"currentBalances"`
		ProjectedBalances struct {
			CashAvailableForTrading    float64 `json:"cashAvailableForTrading"`
			CashAvailableForWithdrawal float64 `json:"cashAvailableForWithdrawal"`
		} `json:"projectedBalances"`
	} `json:"securitiesAccount"`
}

//GetOrderResponses array of GetOrderResponse
type GetOrderResponses []GetOrderResponse

//GetOrderResponse for GetOrder
type GetOrderResponse struct {
	Session    string `json:"session"`
	Duration   string `json:"duration"`
	OrderType  string `json:"orderType"`
	CancelTime struct {
		Date        string `json:"date"`
		ShortFormat bool   `json:"shortFormat"`
	} `json:"cancelTime"`
	ComplexOrderStrategyType string  `json:"complexOrderStrategyType"`
	Quantity                 int     `json:"quantity"`
	FilledQuantity           int     `json:"filledQuantity"`
	RemainingQuantity        int     `json:"remainingQuantity"`
	RequestedDestination     string  `json:"requestedDestination"`
	DestinationLinkName      string  `json:"destinationLinkName"`
	ReleaseTime              string  `json:"releaseTime"`
	StopPrice                float64 `json:"stopPrice"`
	StopPriceLinkBasis       string  `json:"stopPriceLinkBasis"`
	StopPriceLinkType        string  `json:"stopPriceLinkType"`
	StopPriceOffset          float64 `json:"stopPriceOffset"`
	StopType                 string  `json:"stopType"`
	PriceLinkBasis           string  `json:"priceLinkBasis"`
	PriceLinkType            string  `json:"priceLinkType"`
	Price                    float64 `json:"price"`
	TaxLotMethod             string  `json:"taxLotMethod"`
	OrderLegCollection       []struct {
		OrderLegType   string `json:"orderLegType"`
		LegID          int    `json:"legId"`
		Instrument     string `json:"instrument"`
		Instruction    string `json:"instruction"`
		PositionEffect string `json:"positionEffect"`
		Quantity       int    `json:"quantity"`
		QuantityType   string `json:"quantityType"`
	} `json:"orderLegCollection"`
	ActivationPrice          float64  `json:"activationPrice"`
	SpecialInstruction       string   `json:"specialInstruction"`
	OrderStrategyType        string   `json:"orderStrategyType"`
	OrderID                  int      `json:"orderId"`
	Cancelable               bool     `json:"cancelable"`
	Editable                 bool     `json:"editable"`
	Status                   string   `json:"status"`
	EnteredTime              string   `json:"enteredTime"`
	CloseTime                string   `json:"closeTime"`
	Tag                      string   `json:"tag"`
	AccountID                int      `json:"accountId"`
	OrderActivityCollection  []string `json:"orderActivityCollection"`
	ReplacingOrderCollection []struct {
	} `json:"replacingOrderCollection"`
	ChildOrderStrategies []struct {
	} `json:"childOrderStrategies"`
	StatusDescription string `json:"statusDescription"`
}

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
	Theta                  float64 `json:"theta"`
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

//MarketHoursResponse stuff
type MarketHoursResponse struct {
	Option struct {
		Option OptionHours `json:"option"`
	} `json:"option"`
}

//OptionHours stuff
type OptionHours struct {
	Date       string `json:"date"`
	MarketType string `json:"marketType"`
	Product    string `json:"product"`
	IsOpen     bool   `json:"isOpen"`
}
