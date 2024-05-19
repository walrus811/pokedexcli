package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageUrl *string) (LocationArea, error) {
	url := baseURL + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}

	if val, ok := c.cache.Get(url); ok {
		result := LocationArea{}
		marshalErr := json.Unmarshal(val, &result)
		if marshalErr != nil {
			return LocationArea{}, marshalErr
		}
		return result, nil

	}

	req, reqErr := http.NewRequest("GET", url, nil)
	if reqErr != nil {
		return LocationArea{}, reqErr
	}

	res, resErr := c.httpClient.Do(req)
	if resErr != nil {
		return LocationArea{}, resErr
	}
	defer res.Body.Close()

	bodyData, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		return LocationArea{}, readErr
	}

	result := LocationArea{}
	marshalErr := json.Unmarshal(bodyData, &result)

	if marshalErr != nil {
		return LocationArea{}, marshalErr
	}

	c.cache.Add(url, bodyData)
	return result, nil
}

func (c *Client) GetLocationDetails(location string) (LocationAreaDetails, error) {
	url := baseURL + "/location-area/" + location

	if val, ok := c.cache.Get(url); ok {
		result := LocationAreaDetails{}
		marshalErr := json.Unmarshal(val, &result)
		if marshalErr != nil {
			return LocationAreaDetails{}, marshalErr
		}
		return result, nil

	}

	req, reqErr := http.NewRequest("GET", url, nil)
	if reqErr != nil {
		return LocationAreaDetails{}, reqErr
	}

	res, resErr := c.httpClient.Do(req)
	if resErr != nil {
		return LocationAreaDetails{}, resErr
	}
	defer res.Body.Close()

	bodyData, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		return LocationAreaDetails{}, readErr
	}

	result := LocationAreaDetails{}
	marshalErr := json.Unmarshal(bodyData, &result)

	if marshalErr != nil {
		return LocationAreaDetails{}, marshalErr
	}

	c.cache.Add(url, bodyData)
	return result, nil
}
