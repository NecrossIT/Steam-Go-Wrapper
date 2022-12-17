package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Set the Steam app ID and your API key
	appID := 123456
	apiKey := "your-api-key"

	// Build the URL for the API request
	url := fmt.Sprintf("https://store.steampowered.com/api/appdetails?appids=%d&cc=us&filters=minimum_requirement&key=%s", appID, apiKey)

	// Make the request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Parse the JSON response and extract the system requirements
	var data struct {
		AppID struct {
			Data struct {
				MinimumRequirements struct {
					Linux struct {
						Minimum string `json:"minimum"`
						Recommended string `json:"recommended"`
					} `json:"linux"`
					Mac struct {
						Minimum string `json:"minimum"`
						Recommended string `json:"recommended"`
					} `json:"mac"`
					Windows struct {
						Minimum string `json:"minimum"`
						Recommended string `json:"recommended"`
					} `json:"windows"`
				} `json:"minimum_requirement"`
			} `json:"data"`
		} `json:"appid"`
	}
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(err)
		return
	}

	// Print the system requirements as JSON
	reqs := data.AppID.Data.MinimumRequirements
	reqsJSON, err := json.Marshal(reqs)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(reqsJSON))
}
