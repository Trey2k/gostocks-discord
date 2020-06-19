package td

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var tokenRefreshes int = 0

//proccess request
func atuhRequest(data url.Values, endpoint string, response interface{}) error {

	client := &http.Client{}
	request, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return err
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := client.Do(request)
	if err != nil {
		return err
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	return err
}

func getRequest(endpoint string, token string, response interface{}) error {
	request, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return err
	}

	request.Header.Add("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	switch resp.StatusCode {
	case 200:

		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			return err
		}
		return nil
	case 400:
		return errors.New("There was a validation problem with the request")

	case 401:
		tokenRefreshes++
		if tokenRefreshes == 1 {
			var response RequestTokensResponse
			err := refreshTokens(refreshToken, clientCode, &response)
			if err != nil {
				return err
			}
			accessToken = response.AccessToken
			err = getRequest(endpoint, token, &response)
			if err != nil {
				return err
			}
			tokenRefreshes = 0
			return nil
		}
		return errors.New("Error generating new access token. Please launch application again and reauthenticate")
	case 403:
		return errors.New("The caller is forbidden from accessing this page")
	case 404:
		return errors.New("The order was not found")
	case 500:
		return errors.New("There was an unexpected server error")
	default:
		return errors.New("Unkown status code")
	}

}

func postRequest(endpoint string, token string, payload interface{}) error {
	bodyBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	request, err := http.NewRequest("POST", endpoint, bytes.NewReader(bodyBytes))
	if err != nil {
		return err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	switch resp.StatusCode {
	case 200:
		return nil
	case 400:
		return errors.New("There was a validation problem with the request")

	case 401:
		tokenRefreshes++
		if tokenRefreshes == 1 {
			var response RequestTokensResponse
			err := refreshTokens(refreshToken, clientCode, &response)
			if err != nil {
				return err
			}
			accessToken = response.AccessToken
			err = getRequest(endpoint, token, &response)
			if err != nil {
				return err
			}
			tokenRefreshes = 0
			return nil
		}
		return errors.New("Error generating new access token. Please launch application again and reauthenticate")
	case 403:
		return errors.New("The caller is forbidden from accessing this page")
	case 404:
		return errors.New("The order was not found")
	case 500:
		return errors.New("There was an unexpected server error")
	default:
		return errors.New("Unkown status code")
	}

}

func putRequest(endpoint string, token string, payload interface{}) error {
	bodyBytes, err := json.Marshal(payload)

	request, err := http.NewRequest("PUT", endpoint, bytes.NewReader(bodyBytes))
	if err != nil {
		return err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	switch resp.StatusCode {
	case 200:
		return nil
	case 400:
		return errors.New("There was a validation problem with the request")

	case 401:
		tokenRefreshes++
		if tokenRefreshes == 1 {
			var response RequestTokensResponse
			err := refreshTokens(refreshToken, clientCode, &response)
			if err != nil {
				return err
			}
			accessToken = response.AccessToken
			err = getRequest(endpoint, token, &response)
			if err != nil {
				return err
			}
			tokenRefreshes = 0
			return nil
		}
		return errors.New("Error generating new access token. Please launch application again and reauthenticate")
	case 403:
		return errors.New("The caller is forbidden from accessing this page")
	case 404:
		return errors.New("The order was not found")
	case 500:
		return errors.New("There was an unexpected server error")
	default:
		return errors.New("Unkown status code")
	}

}

func deleteRequest(endpoint string, token string) error {
	request, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		return err
	}

	request.Header.Add("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return err
	}
	switch resp.StatusCode {
	case 200:
		return nil
	case 400:
		return errors.New("There was a validation problem with the request")

	case 401:
		tokenRefreshes++
		if tokenRefreshes == 1 {
			var response RequestTokensResponse
			err := refreshTokens(refreshToken, clientCode, &response)
			if err != nil {
				return err
			}
			accessToken = response.AccessToken
			err = getRequest(endpoint, token, &response)
			if err != nil {
				return err
			}
			tokenRefreshes = 0
			return nil
		}
		return errors.New("Error generating new access token. Please launch application again and reauthenticate")
	case 403:
		return errors.New("The caller is forbidden from accessing this page")
	case 404:
		return errors.New("The order was not found")
	case 500:
		return errors.New("There was an unexpected server error")
	default:
		return errors.New("Unkown status code")
	}

}
