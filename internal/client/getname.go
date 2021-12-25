package client

import (
	"encoding/json"
	"io"
	"net/http"
	"sort"
	"strings"

	"github.com/adrg/strutil"
	"github.com/adrg/strutil/metrics"
)

func (c *Client) GetNameFromInput(name string) ([]string, error) {
	URL := "https://api.steampowered.com/ISteamApps/GetAppList/v2/"
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	//dec := json.NewDecoder(resp.Body)
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var AppList AppListResponse

	//err = dec.Decode(&AppList)
	err = json.Unmarshal(b, &AppList)
	if err != nil {
		return nil, err
	}

	j := metrics.NewSorensenDice()

	similarityMap := make(map[string]float64)
	splitName := strings.Split(name, " ")
	for _, item := range AppList.Applist.Apps {
		similarity := strutil.Similarity(name, item.Name, j)
		similarityMap[item.Name] = similarity
		splitItemName := strings.Split(item.Name, " ")
		for _, word := range splitName {
			for _, word2 := range splitItemName {
				if strings.EqualFold(word, word2) {
					similarityMap[item.Name] += float64(len(word)) * 0.05
				}
			}
		}
		if strings.Contains(item.Name, name) {
			similarityMap[item.Name] += 0.3
		}
		if strings.EqualFold(item.Name, name) {
			return []string{name}, nil
		}
	}
	return rankMapStringfloat(similarityMap), nil
}

func rankMapStringfloat(values map[string]float64) []string {
	type kv struct {
		Key   string
		Value float64
	}
	var ss []kv
	for k, v := range values {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	ranked := make([]string, len(values))
	for i, kv := range ss {
		ranked[i] = kv.Key
	}
	return ranked[0:5]
}
