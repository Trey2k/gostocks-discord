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
