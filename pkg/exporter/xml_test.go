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

	tests := []struct {
		name     string
		input    interface{}
		expected interface{}
	}{
		{"SolarFlux", data.SolarFlux, 145},
		{"AIndex", data.AIndex, 19},
		{"KIndex", data.KIndex, 3},
		{"KIndexNT", data.KIndexNT, "No Report"},
		{"XRay", data.XRay, XRay(720)},
		{"Sunspots", data.Sunspots, uint(105)},
		{"HeliumLine", data.HeliumLine, float32(142.3)},
		{"ProtonFlux", data.ProtonFlux, uint(80)},
		{"ElectronFlux", data.ElectronFlux, uint(2100)},
		{"Aurora", data.Aurora, 1},
		{"AuroraLatitude", data.AuroraLatitude, float32(67.5)},
		{"Normalization", data.Normalization, float32(1.99)},
		{"SolarWind", data.SolarWind, float32(604.7)},
		{"MagneticField", data.MagneticField, float32(1.9)},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.input != test.expected {
				t.Errorf("%s: Expected value %v, got %v", test.name, test.input, test.expected)
			}
		})
	}
}
