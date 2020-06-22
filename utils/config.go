package utils

import (
	"encoding/json"
	"io/ioutil"
)

//Config object
var Config *Configuration

//TradeSettings trade settings
type TradeSettings struct {
	MakeRiskyTrades             bool
	RiskyInvestPercent          float64
	SafeInvestPercent           float64
	SafeStopLossPercent         float64
	RiskyStopLossPercent        float64
	AllowedPriceIncreasePercent float64
	UpdateInterval              int
	AutoSellProfitPercent       float64
	UseUserWhitlist             bool
	WhitelistUserIDs            []string
}

//SettingsStruct this is exported
type SettingsStruct struct {
	Trade TradeSettings
}

//DiscordInfo : this is exported
type DiscordInfo struct {
	Token      string
	GuildIDs   []string
	ChannelIDs []string
}

//TDA : this is exported
type TDA struct {
	ClientKey       string
	CallbackAddress string
	AccountID       string
}

//MySQLStruct : this is exported
type MySQLStruct struct {
	Username string
	Password string
	IP       string
	Port     string
	Database string
}

//Configuration : this is exported
type Configuration struct {
	Discord  DiscordInfo
	TD       TDA
	MySQL    MySQLStruct
	Settings SettingsStruct
}

//GetConfig object
func GetConfig() (configuration *Configuration, err error) {
	if FileExists("config.json") { //Get user set configuration from config.ini if file exists
		b, err := ioutil.ReadFile("config.json")
		if err != nil {
			ErrCheck("Failed to unmarshal configuration file", err)
			return nil, err
		}

		config := &Configuration{}

		err = json.Unmarshal(b, config)
		if err != nil {
			ErrCheck("Failed to unmarshal configuration file", err)
			return nil, err
		}

		return config, nil

	} //Otherwise, create config.ini, init it with the default config, then return the default config

	config := &Configuration{ //Default configuration
		DiscordInfo{
			Token:      "",
			GuildIDs:   []string{"Guild ID1", "Guils ID2"},
			ChannelIDs: []string{"Channel ID1", "Channel ID2"},
		},
		TDA{
			ClientKey:       "Client Key",
			CallbackAddress: "127.0.0.1:8080",
			AccountID:       "Account ID",
		},
		MySQLStruct{
			Username: "root",
			Password: "",
			IP:       "127.0.0.1",
			Port:     "3306",
			Database: "GoStocks",
		},
		SettingsStruct{
			Trade: TradeSettings{
				MakeRiskyTrades:             true,
				RiskyInvestPercent:          0.05,
				SafeInvestPercent:           0.10,
				SafeStopLossPercent:         0.60,
				RiskyStopLossPercent:        0.80,
				AllowedPriceIncreasePercent: 0.11,
				UpdateInterval:              10,
				AutoSellProfitPercent:       0.4,
				UseUserWhitlist:             true,
				WhitelistUserIDs:            []string{"116377104035086339", "105036460108865536"},
			},
		},
	}

	b, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		ErrCheck("Failed to marshal configuration file", err)
		return nil, err
	}

	err = ioutil.WriteFile("config.json", b, 0644)
	if err != nil {
		ErrCheck("Failed to write config file. Is it locked by another process?", err)
		return nil, err
	}

	return config, nil //Return default configuration

}
