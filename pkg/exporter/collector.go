package exporter

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

const (
	// User-Agent to use when performing requests
	user_agent = "solardata_exporter/%s (https://github.com/tzermias/solardata_exporter)"
)

// fetch data from the feed and parse them.
func fetchData(feed_url string) (SolarData, error) {
	var data SolarData
	var solar Solar

	// Request
	req, err := http.NewRequest(http.MethodGet, feed_url, nil)
	if err != nil {
		return data, err
	}
	req.Header.Set("User-Agent", fmt.Sprintf(user_agent, Version))

	// HTTP client
	client := http.Client{}

	// Make request
	res, err := client.Do(req)
	if err != nil {
		return data, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return data, err
	}

	// Check response statusCodes
	if res.StatusCode != http.StatusOK {
		return data, fmt.Errorf("Server responded with status %d: %s", res.StatusCode, body)
	}
	// Parse XML
	err = xml.Unmarshal([]byte(body), &solar)
	if err != nil {
		return data, err
	}

	return solar.Data, nil
}
