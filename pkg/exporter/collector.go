package exporter

import (
	"encoding/xml"
	"fmt"

	"github.com/mmcdole/gofeed"
)

const (
	// User-Agent to use when performing requests
	user_agent = "solardata_exporter/%s (https://github.com/tzermias/solardata_exporter)"
)

// fetch data from the feed and parse them.
func fetchData(feed_url string) (SolarData, error) {
	var data SolarData

	fp := gofeed.NewParser()
	fp.UserAgent = fmt.Sprintf(user_agent, Version)
	feed, err := fp.ParseURL(feed_url)
	if err != nil {
		return data, err
	}
	// Parse XML
	err = xml.Unmarshal([]byte(feed.Items[0].Custom["solar"]), &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
