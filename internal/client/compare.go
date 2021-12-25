package client

import "gameservice/internal/model"

func (c *Client) GetAllPricesByName(name string) ([]model.GamePriceResponse, error) {
	SteamPrice, err := c.GetSteamPriceByName(name)
	if err != nil {
		return nil, err
	}
	SteamPayPrice, err := c.GetSteamPayPriceByName(name)
	if err != nil {
		return nil, err
	}
	GOGPrice, err := c.GetGOGPriceByName(name)
	if err != nil {
		return nil, err
	}
	PlatiruPrice, err := c.GetPlatiruPriceByName(name)
	if err != nil {
		return nil, err
	}
	return []model.GamePriceResponse{SteamPrice, SteamPayPrice, GOGPrice, PlatiruPrice}, nil
}
