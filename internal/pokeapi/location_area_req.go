package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func performRequest[T any](url string, client *Client, responseType T) (T, error) {
	// Create request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return responseType, err
	}

	// check the cache
	dat, ok := client.cache.Get(url)
	if ok {
		var respObj T
		err := json.Unmarshal(dat, &respObj)
		if err != nil {
			return responseType, err
		}
		return respObj, nil
	}

	// Execute request
	resp, err := client.httpClient.Do(req)
	if err != nil {
		return responseType, err
	}

	if resp.StatusCode > 399 {
		return responseType, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	// Get data from request
	data, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return responseType, err
	}

	client.cache.Add(url, data)

	// Decode type from data
	var respObj T
	err = json.Unmarshal(data, &respObj)
	if err != nil {
		return responseType, err
	}

	return respObj, nil
}

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint

	if pageURL != nil {
		fullURL = *pageURL
	}

	var response LocationAreasResp
	response, err := performRequest(fullURL, c, response)
	if err != nil {
		panic(err)
	}
	return response, nil
}

func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + locationAreaName
	fullURL := baseURL + endpoint

	var response LocationArea
	response, err := performRequest(fullURL, c, response)
	if err != nil {
		panic(err)
	}
	return response, nil
}
