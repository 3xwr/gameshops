package client

import (
	"gameservice/internal/model"
	"log"
	"sync"
)

func (c *Client) GetAllPricesByName(name string) ([]model.GamePriceResponse, error) {

	var SteamPrice model.GamePriceResponse
	var SteamPayPrice model.GamePriceResponse
	var GOGPrice model.GamePriceResponse
	var PlatiruPrice model.GamePriceResponse

	var err error

	var wg sync.WaitGroup
	wg.Add(4)

	go func() {
		defer wg.Done()
		SteamPrice, err = c.GetSteamPriceByName(name)
		if err != nil {
			log.Println(err)
		}
	}()

	go func() {
		defer wg.Done()
		SteamPayPrice, err = c.GetSteamPayPriceByName(name)
		if err != nil {
			log.Println(err)
		}
	}()

	go func() {
		defer wg.Done()
		GOGPrice, err = c.GetGOGPriceByName(name)
		if err != nil {
			log.Println(err)
		}
	}()

	go func() {
		defer wg.Done()
		PlatiruPrice, err = c.GetPlatiruPriceByName(name)
		if err != nil {
			log.Println(err)
		}
	}()

	wg.Wait()

	return []model.GamePriceResponse{SteamPrice, SteamPayPrice, GOGPrice, PlatiruPrice}, nil
}
