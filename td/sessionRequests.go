package td

import (
	"fmt"
	"net/url"
	"time"

	"github.com/Trey2k/gostocks-discord/utils"
)

//RequestTokensResponse response struct for token request
type RequestTokensResponse struct {
	AccessToken           string `json:"access_token"`
	RefreshToken          string `json:"refresh_token"`
	Scope                 string `json:"scope"`
	ExpiresIn             int    `json:"expires_in"`
	RefreshTokenExpiresIn int    `json:"refresh_token_expires_in"`
	TokenType             string `json:"token_type"`
}

func requestTokens(oauth string, clientCode string) (string, string, error) {
	var response RequestTokensResponse
	endpoint := "https://api.tdameritrade.com/v1/oauth2/token"

	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("access_type", "offline")
	data.Set("code", oauth)
	data.Set("client_id", clientCode)
	data.Set("redirect_uri", "https://127.0.0.1:8080")
	var attempts int
	for {
		attempts++
		code, err := atuhRequest(data, endpoint, &response)
		if err == nil && code == 200 {
			return response.AccessToken, response.RefreshToken, err
		}
		utils.Log("Failed to generate new tokens "+fmt.Sprint(attempts)+" time(s) with status code "+fmt.Sprint(code)+" retrying in 20 secounds", utils.LogError)
		time.Sleep(20 * time.Second)
	}
}

func refreshTokens(refreshToken string, clientCode string) (string, error) {
	var response RequestTokensResponse
	endpoint := "https://api.tdameritrade.com/v1/oauth2/token"

	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", refreshToken)
	data.Set("access_type", "request")
	data.Set("client_id", clientCode)
	data.Set("redirect_uri", "https://127.0.0.1:8080")
	var attempts int
	for {
		attempts++
		code, err := atuhRequest(data, endpoint, &response)
		if err == nil && code == 200 {
			return response.AccessToken, err
		}
		utils.Log("Failed to generate new access token "+fmt.Sprint(attempts)+" time(s) with status code "+fmt.Sprint(code)+" retrying in 20 secounds", utils.LogError)
		time.Sleep(20 * time.Second)
	}
}
