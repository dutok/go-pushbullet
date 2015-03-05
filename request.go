package main

import (
	"io/ioutil"
	"net/http"
	"bytes"
	"encoding/json"
	"errors"
)

var (
    key = ""
    api = "https://api.pushbullet.com/v2/"
)

type ApiError struct {
	Error struct {
		Cat     string `json:"cat"`
		Message string `json:"message"`
		Type    string `json:"type"`
	} `json:"error"`
}


func request(action, object, value string) ([]byte, error) {
    url := api + object
	req, err := http.NewRequest(action, url, bytes.NewBuffer([]byte(value)))
	req.Header.Set("Authorization", "Bearer " + key)
	req.Header.Set("Content-Type", "application/json")

	// Handle the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	
	if resp.StatusCode >= 500 {
	    err = errors.New("Server Error - Something went wrong on Pushbullet's side.")
	    return nil, err
	}
	
	respbytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	
	apiError := ApiError{}
	if resp.StatusCode >= 400 {
        err = json.Unmarshal(respbytes, &apiError)
        if err != nil {
            return nil, err
        }
        err = errors.New(apiError.Error.Message)
        return nil,  err
	}

	return respbytes, nil
}