package main

import (
	"encoding/json"
	"fmt"
)

type Me struct {
	Created         float64 `json:"created"`
	Email           string  `json:"email"`
	EmailNormalized string  `json:"email_normalized"`
	Iden            string  `json:"iden"`
	ImageURL        string  `json:"image_url"`
	Modified        float64 `json:"modified"`
	Name            string  `json:"name"`
	Preferences     struct {
		Onboarding struct {
			App       bool `json:"app"`
			Extension bool `json:"extension"`
			Friends   bool `json:"friends"`
		} `json:"onboarding"`
		Social bool `json:"social"`
	} `json:"preferences"`
}

func getMe() (Me, error) {
	me := Me{}
	respbytes, err := request("GET", "users/me", "")
	if err != nil {
		return me, err
	}
	err = json.Unmarshal(respbytes, &me)
	if err != nil {
		fmt.Println(err)
	}
	return me, nil
}
