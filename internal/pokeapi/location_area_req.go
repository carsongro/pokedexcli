package pokeapi

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
