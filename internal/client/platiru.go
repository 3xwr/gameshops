package client

import (
	"encoding/json"
	"gameservice/internal/model"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func (c *Client) GetPlatiruPriceByName(name string) (model.GamePriceResponse, error) {
	link := "https://plati.io/api/search.ashx?"
	params := url.Values{}
	params.Add("response", "json")
	params.Add("query", name)

	link = link + params.Encode()

	req, err := http.NewRequest(http.MethodGet, link, nil)
	if err != nil {
		log.Println(err)
		return model.GamePriceResponse{}, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		log.Println(err)
		return model.GamePriceResponse{}, err
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return model.GamePriceResponse{}, err
	}

	var PlatiruResponse model.PlatiruResponseModel

	err = json.Unmarshal(b, &PlatiruResponse)
	if err != nil {
		log.Println(err)
		return model.GamePriceResponse{}, nil
	}

	var PriceResponse model.GamePriceResponse
	PriceResponse.StoreName = "platiru"
	PriceResponse.StoreAppName = name

	minPrice := 9999999.9

	found := false

	for _, item := range PlatiruResponse.Items {
		if strings.Contains(item.NameEng, name) && strings.Contains(strings.ToLower(item.NameEng), "key") {
			if item.PriceRur < minPrice {
				minPrice = item.PriceRur
				PriceResponse.StoreAppID = item.ID
				found = true
			}
		}
	}

	if found {
		PriceResponse.StorePrice = strconv.Itoa(int(minPrice)) + " руб."
	}

	if !found {
		PriceResponse.Status = "game not found in store"
	}
	return PriceResponse, nil
}
