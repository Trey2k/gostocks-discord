package td

//GetAccountResponses array of GetAccountResponse
type GetAccountResponses []GetAccountResponse

//GetAccountResponse struct
type GetAccountResponse struct {
	Account AccountInfo `json:"securitiesAccount"`
}

//AccountInfo stuff
type AccountInfo struct {
	Type                    string                  `json:"type"`
	AccountID               string                  `json:"accountId"`
	RoundTrips              int                     `json:"roundTrips"`
	IsDayTrader             bool                    `json:"isDayTrader"`
	IsClosingOnlyRestricted bool                    `json:"isClosingOnlyRestricted"`
	InitialBalances         Ballance                `json:"initialBalances"`
	CurrentBalances         Ballance                `json:"currentBalances"`
	ProjectedBalances       ProjectedBalancesStruct `json:"projectedBalances"`
}

//Ballance stuff
type Ballance struct {
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
}

//ProjectedBalancesStruct Stuff
type ProjectedBalancesStruct struct {
	CashAvailableForTrading    float64 `json:"cashAvailableForTrading"`
	CashAvailableForWithdrawal float64 `json:"cashAvailableForWithdrawal"`
}

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
