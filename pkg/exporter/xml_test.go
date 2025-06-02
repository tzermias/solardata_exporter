package exporter

import (
	_ "embed"
	"encoding/xml"
	"testing"

	"github.com/mmcdole/gofeed"
)

//go:embed testdata/success.xml
var success_data string

func TestXMLParsing(t *testing.T) {
	var data SolarData

	fp := gofeed.NewParser()
	feed, err := fp.ParseString(success_data)
	if err != nil {
		t.Fatalf("Couldn't parse RSS feed. Error: %v", err)
	}

	err = xml.Unmarshal([]byte(feed.Items[0].Custom["solar"]), &data)
	if err != nil {
		t.Fatalf("Couldn't parse XML data. Error: %v", err)
	}

	//Solarflux
	if data.SolarFlux != 145 {
		t.Errorf("Expected SolarFlux value 145, got %d", data.SolarFlux)
	}
}
