package scraping

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/AppleBoiy/pydumper/go/crawler/app/lib/config"
)

func getURLFromWord(w string) ([]string, error) {
	apiURL := fmt.Sprintf(
		"https://www.googleapis.com/customsearch/v1?key=%s&cx=%s&q=%s",
		config.Config.APIKey, config.Config.SearchEngineID, url.QueryEscape(w),
	)

	resp, err := http.Get(apiURL)
	if err != nil {
		return []string{}, fmt.Errorf("failed to call Google API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []string{}, fmt.Errorf("HTTP Status Error: %d", resp.StatusCode)
	}

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return []string{}, fmt.Errorf("failed to parse JSON: %w", err)
	}

	items, have := result["items"].([]interface{})
	if !have {
		return []string{}, nil
	}

	urls := []string{}

	for index, it := range items {
		if config.Config.WebMax != 0 && index == config.Config.WebMax {
			break
		}
		url := it.(map[string]interface{})["link"].(string)
		urls = append(urls, url)
	}

	return urls, nil
}
