package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

//proccess request
func procRequest(data url.Values, endpoint string, response interface{}) error {

	client := &http.Client{}
	r, err := http.NewRequest("POST", endpoint, strings.NewReader(data.Encode()))
	errCheck("Error creating post request", err)

	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, err := client.Do(r)
	if err != nil {
		return err
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	return err
}
