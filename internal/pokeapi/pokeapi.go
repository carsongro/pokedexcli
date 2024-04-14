package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/carsongro/pokedexcli/internal/pokecache"
)

const baseURL = "https://pokeapi.co/api/v2"

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}

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

	// execute request
	resp, err := client.httpClient.Do(req)
	if err != nil {
		return responseType, err
	}

	if resp.StatusCode > 399 {
		return responseType, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	// get data from request
	data, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return responseType, err
	}

	client.cache.Add(url, data)

	// decode type from data
	var respObj T
	err = json.Unmarshal(data, &respObj)
	if err != nil {
		return responseType, err
	}

	return respObj, nil
}
