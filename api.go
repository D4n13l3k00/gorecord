package main

import (
	"encoding/json"
	"net/http"
)

type API struct {
	Endpoint     string
	CachedResult StationsAPIResult
}

func NewAPI() API {
	return API{
		Endpoint: "https://www.radiorecord.ru/api/stations/",
	}
}

func (api *API) Get() error {
	url := api.Endpoint
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var result StationsAPIResult
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return err
	}

	api.CachedResult = result
	return nil
}
