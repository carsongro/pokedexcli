package pokeapi

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemonName
	fullURL := baseURL + endpoint

	var response Pokemon
	response, err := performRequest(fullURL, c, response)
	if err != nil {
		panic(err)
	}
	return response, nil
}
