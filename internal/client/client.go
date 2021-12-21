package client

import (
	"encoding/json"
	"fmt"
	"gameservice/internal/model"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type Client struct {
	client http.Client
}

type AppListResponse struct {
	Applist struct {
		Apps []struct {
			AppID int
			Name  string
		}
	}
}

type AppInfoResponse struct {
	Name          string `json:"name"`
	PriceOverview struct {
		Currency string `json:"currency"`
		Final    string `json:"final"`
	} `json:"price_overview"`
}

func New() *Client {
	return &Client{
		client: http.Client{
			Timeout: time.Duration(time.Minute),
		},
	}
}

func (c *Client) GetAppIDByName(name string) (int, error) {
	log.Println(name)
	URL := "https://api.steampowered.com/ISteamApps/GetAppList/v2/"
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return -1, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return -1, err
	}

	defer resp.Body.Close()

	//dec := json.NewDecoder(resp.Body)
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return -1, err
	}

	var AppList AppListResponse

	//err = dec.Decode(&AppList)
	err = json.Unmarshal(b, &AppList)
	if err != nil {
		return -1, err
	}

	counter := 0
	for _, item := range AppList.Applist.Apps {
		counter++
		if strings.EqualFold(item.Name, name) {
			return item.AppID, nil
		}
	}
	log.Println(counter, " apps found.")
	return -1, nil
}

func (c *Client) GetSteamPriceByName(name string) (model.GamePriceResponse, error) {
	appID, err := c.GetAppIDByName(name)
	if err != nil {
		return model.GamePriceResponse{}, err
	}

	if appID == -1 {
		return model.GamePriceResponse{StoreName: "steam", StoreAppName: name, Status: "game not found in store"}, nil
	}

	Price, err := c.GetAppPriceByID(appID)
	if err != nil {
		return model.GamePriceResponse{}, err
	}

	return model.GamePriceResponse{StoreName: "steam", StoreAppID: appID, StoreAppName: name, StorePrice: Price}, nil
}

func (c *Client) GetAppPriceByID(ID int) (string, error) {
	fmt.Println("Looking for Steam app with ID - ", ID)
	baseURL := "https://store.steampowered.com/api/appdetails?"
	params := url.Values{}
	params.Add("appids", strconv.Itoa(ID))
	params.Add("cc", "ru")
	params.Add("l", "ru")

	link := baseURL + params.Encode()

	req, err := http.NewRequest(http.MethodGet, link, nil)
	if err != nil {
		return "", err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var objmap map[string]interface{}

	err = json.Unmarshal(b, &objmap)
	if err != nil {
		return "", err
	}

	var Final interface{}

	StringId := strconv.Itoa(ID)
	Info := objmap[StringId].(map[string]interface{})
	Data := Info["data"].(map[string]interface{})
	if Data["price_overview"] == nil {
		Final = 0.0
	} else {
		PriceOverview := Data["price_overview"].(map[string]interface{})
		Final = PriceOverview["final"]
	}
	return fmt.Sprintf("%d руб.", int(Final.(float64))/100), nil
}
