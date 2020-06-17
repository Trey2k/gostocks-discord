package utils

import (
	"encoding/json"
	"io/ioutil"
)

//Config object
var Config *Configuration

//Settings this is exported
type Settings struct {
}

//DiscordInfo : this is exported
type DiscordInfo struct {
	Username   string
	Password   string
	GuildID    string
	ChannelID  string
	GameStatus string
}

//TDAPI : this is exported
type TDAPI struct {
	ClientKey       string
	CallbackAddress string
}

//Configuration : this is exported
type Configuration struct {
	Discord DiscordInfo
	TD      TDAPI
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
		TDAPI{
			ClientKey:       "Client Key",
			CallbackAddress: "127.0.0.1:8080",
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
