package client

import (
	"encoding/json"
	"gameservice/internal/model"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

func (c *Client) GetGOGPriceByName(name string) (model.GamePriceResponse, error) {
	StoreName := "gog"
	if p, ok := c.repo.Load(name + StoreName); ok {
		log.Println("GOG price for", name, "found in cache")
		t := time.Now()
		diff := t.Sub(p.Timestamp)
		if diff < c.repo.GetTimeout() {
			return p, nil
		}
	}
	link := "https://embed.gog.com/games/ajax/filtered?mediaType=game&"

	params := url.Values{}
	params.Add("search", name)
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

	var GOGResponse model.GOGResponseModel

	err = json.Unmarshal(b, &GOGResponse)
	if err != nil {
		log.Println("Unmarshal error", err)
		return model.GamePriceResponse{}, err
	}

	var PriceResponse model.GamePriceResponse
	PriceResponse.StoreName = "gog"
	PriceResponse.StoreAppName = name

	found := false

	for _, i := range GOGResponse.Products {
		if i.Title == name {
			PriceResponse.StoreAppID = i.ID
			PriceResponse.StoreImage = "https:" + i.Image
			PriceResponse.StorePrice = i.Price.FinalAmount + " руб."
			PriceResponse.StoreAppURL = "https://gog.com" + i.URL
			found = true
		}
	}

	if !found {
		PriceResponse.Status = "game not found in store"
	}
	c.repo.Store(name+StoreName, PriceResponse)
	return PriceResponse, nil
}
