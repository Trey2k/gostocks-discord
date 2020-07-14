package td

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

var tokenRefreshes int = 0

//proccess request
func atuhRequest(data url.Values, endpoint string, response interface{}) (int, error) {

	client := &http.Client{}
	request, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode()))
	if err != nil {
		return 0, err
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := client.Do(request)
	if err != nil {
		return 0, err
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	return resp.StatusCode, err
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
		access, err := refreshTokens(refreshToken, clientCode)
		if err != nil {
			return err
		}
		accessToken = access
		err = getRequest(endpoint, accessToken, &response)
		if err != nil {
			return err
		}
		return nil
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

	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return err
	}
	request.Header.Add("Content-Length", fmt.Sprint(len(bodyBytes)))
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
	case 201:
		return nil
	case 400:
		return errors.New("There was a validation problem with the request")

	case 401:
		access, err := refreshTokens(refreshToken, clientCode)
		if err != nil {
			return err
		}
		accessToken = access
		err = postRequest(endpoint, accessToken, payload)
		if err != nil {
			return err
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
		access, err := refreshTokens(refreshToken, clientCode)
		if err != nil {
			return err
		}
		accessToken = access
		err = putRequest(endpoint, accessToken, payload)
		if err != nil {
			return err
		}
		tokenRefreshes = 0
		return nil

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
		access, err := refreshTokens(refreshToken, clientCode)
		if err != nil {
			return err
		}
		accessToken = access
		err = deleteRequest(endpoint, accessToken)
		if err != nil {
			return err
		}
		tokenRefreshes = 0
		return nil
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
