package utils

import (
	"encoding/json"
	"io/ioutil"
)

//Config object
var Config *Configuration

//TradeSettings trade settings
type TradeSettings struct {
	MakeRiskyTrades       bool
	RiskyInvestPercentage float64
	SafeInvestPercentage  float64
	StopLossPercentage    float64
	UseUserWhitlist       bool
	WhitelistUserIDs      []string
}

//SettingsStruct this is exported
type SettingsStruct struct {
	Trade TradeSettings
}

//DiscordInfo : this is exported
type DiscordInfo struct {
	Username   string
	Password   string
	GuildID    string
	ChannelID  string
	GameStatus string
}

//TDA : this is exported
type TDA struct {
	ClientKey       string
	CallbackAddress string
	AccountID       string
}

//Configuration : this is exported
type Configuration struct {
	Discord  DiscordInfo
	TD       TDA
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
			Username:   "Email",
			Password:   "",
			GuildID:    "",
			ChannelID:  "",
			GameStatus: "The stock markets",
		},
		TDA{
			ClientKey:       "Client Key",
			CallbackAddress: "127.0.0.1:8080",
			AccountID:       "Account ID",
		},
		SettingsStruct{
			Trade: TradeSettings{
				MakeRiskyTrades:       true,
				RiskyInvestPercentage: 0.05,
				SafeInvestPercentage:  0.10,
				StopLossPercentage:    0.50,
				UseUserWhitlist:       true,
				WhitelistUserIDs:      []string{"116377104035086339", "105036460108865536"},
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
