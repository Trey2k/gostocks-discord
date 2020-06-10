package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var config *Configuration

//Configuration : this is exported
type Configuration struct {
	Token      string
	GuildID    string
	ChannelID  string
	GameStatus string
}

func getConfig() (configuration *Configuration, err error) {
	if fileExists("config.json") { //Get user set configuration from config.ini if file exists
		b, err := ioutil.ReadFile("config.json")
		if err != nil {
			errCheck("Failed to unmarshal configuration file", err)
			return nil, err
		}

		config := &Configuration{}

		err = json.Unmarshal(b, config)
		if err != nil {
			errCheck("Failed to unmarshal configuration file", err)
			return nil, err
		}

		return config, nil

	} //Otherwise, create config.ini, init it with the default config, then return the default config

	config := &Configuration{ //Default configuration
		Token:      "",
		GuildID:    "",
		ChannelID:  "",
		GameStatus: "The stock markets",
	}

	b, err := json.MarshalIndent(config, "", " ")
	if err != nil {
		errCheck("Failed to marshal configuration file", err)
		return nil, err
	}

	err = ioutil.WriteFile("config.json", b, 0644)
	if err != nil {
		errCheck("Failed to write config file. Is it locked by another process?", err)
		return nil, err
	}

	return config, nil //Return default configuration

}

//Check to see if a file exists by name. Return bool
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
