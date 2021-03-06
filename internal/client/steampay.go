package client

import (
	"encoding/json"
	"gameservice/internal/model"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func (c *Client) GetSteamPayPriceByName(name string) (model.GamePriceResponse, error) {
	StoreName := "steampay"
	if p, ok := c.repo.Load(name + StoreName); ok {
		log.Println("SteamPay price for", name, "found in cache")
		t := time.Now()
		diff := t.Sub(p.Timestamp)
		if diff < c.repo.GetTimeout() {
			return p, nil
		}
	}
	link := "https://steampay.com/api/search?"

	params := url.Values{}
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

	var SteamPayResponse model.SteamPayResponseModel

	err = json.Unmarshal(b, &SteamPayResponse)
	if err != nil {
		log.Println(err)
		return model.GamePriceResponse{}, nil
	}

	var PriceResponse model.GamePriceResponse
	PriceResponse.StoreName = "steampay"
	PriceResponse.StoreAppName = name

	found := false

	for _, i := range SteamPayResponse.Products {
		if i.Title == name {
			PriceResponse.StorePrice = strconv.Itoa(i.Prices.Rub) + " руб."
			PriceResponse.StoreImage = i.Image
			PriceResponse.StoreAppURL = i.URL
			found = true
		}
	}

	if !found {
		PriceResponse.Status = "game not found in store"
	}

	c.repo.Store(name+StoreName, PriceResponse)
	return PriceResponse, nil
}
