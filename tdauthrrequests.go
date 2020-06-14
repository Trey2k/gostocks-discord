package main

import "net/url"

func requestTokens(oauth string, clientCode string, response *RequestTokensResponse) error {

	endpoint := "https://api.tdameritrade.com/v1/oauth2/token"

	data := url.Values{}
	data.Set("grant_type", "authorization_code")
	data.Set("access_type", "offline")
	data.Set("code", oauth)
	data.Set("client_id", clientCode)
	data.Set("redirect_uri", "https://127.0.0.1:8080")

	err := procRequest(data, endpoint, &response)
	return err
}

func refreshTokens(refreshToken string, clientCode string, response *RequestTokensResponse) error {

	endpoint := "https://api.tdameritrade.com/v1/oauth2/token"

	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", refreshToken)
	data.Set("access_type", "request")
	data.Set("client_id", clientCode)
	data.Set("redirect_uri", "https://127.0.0.1:8080")

	err := procRequest(data, endpoint, &response)
	return err
}
