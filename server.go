package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const IPLocationAPIEndpoint = "http://ipwho.is/"

type GeoLocationAPIReponse struct {
	IP      string `json:"ip"`
	Country string `json:"country"`
	City    string `json:"city"`
	Region  string `json:"region"`
}

func index(res http.ResponseWriter, req *http.Request) {
	resp, err := http.Get(IPLocationAPIEndpoint)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Errorf("failed to request: %#v", err)
		http.Error(res, "Could not request geolocation endpoint", http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	var apiResponse GeoLocationAPIReponse

	if err := json.Unmarshal(body, &apiResponse); err != nil {
		fmt.Errorf("failed to unmarshall: %#v", err)
		http.Error(res, "Could not unmarshall geolocation data", http.StatusInternalServerError)
		return
	}

	responseString := fmt.Sprintf("Hello from %s, %s", apiResponse.City, apiResponse.Country)

	res.Write([]byte(responseString))
}

func main() {
	http.HandleFunc("/", index)
	fmt.Println("Serving on :7654")
	_ = http.ListenAndServe(":7654", nil)
}
