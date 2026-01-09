package exporter

import (
	_ "embed"
	"encoding/xml"
	"testing"
	"time"
)

//go:embed testdata/success.xml
var success_data string

func TestXMLParsing(t *testing.T) {
	var solar Solar
	var data SolarData

	err := xml.Unmarshal([]byte(success_data), &solar)
	if err != nil {
		t.Fatalf("Couldn't parse XML data. Error: %v", err)
	}

	data = solar.Data
	//Solarflux
	if data.SolarFlux != 145 {
		t.Errorf("Expected SolarFlux value 145, got %d", data.SolarFlux)
	}

	tests := []struct {
		name     string
		input    interface{}
		expected interface{}
	}{
		{"Updated", data.Updated.UTC(), time.Date(2025, 03, 01, 23, 56, 00, 0, time.UTC)},
		{"SolarFlux", data.SolarFlux, 145},
		{"AIndex", data.AIndex, 19},
		{"KIndex", data.KIndex, 3},
		{"KIndexNT", data.KIndexNT, "No Report"},
		{"XRay", data.XRay, XRay(720)},
		{"Sunspots", data.Sunspots, uint(105)},
		{"HeliumLine", data.HeliumLine, float32(142.3)},
		{"ProtonFlux", data.ProtonFlux, ProtonFlux(80)},
		{"ElectronFlux", data.ElectronFlux, ElectronFlux(2100)},
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

func TestParseNoRpt(t *testing.T) {
	// A rudimentary test to check "NoRpt" parsing. Should always return 0
	test_data := `
	<?xml version="1.0"?>
	<solar>
			<solardata>
				<source url="http://www.hamqsl.com/solar.html">N0NBH</source>
				<updated> 09 Jan 2026 0851 GMT</updated>
				<solarflux>140</solarflux>
				<sunspots>84</sunspots>
				<heliumline>128.2</heliumline>
				<protonflux>NoRpt</protonflux>
				<electonflux>NoRpt</electonflux>
			</solardata>
		</solar>
`
	var solar Solar
	var data SolarData

	err := xml.Unmarshal([]byte(test_data), &solar)
	if err != nil {
		t.Fatalf("Couldn't parse XML data. Error: %v", err)
	}

	data = solar.Data

	//Solarflux
	if data.ProtonFlux != 0 {
		t.Errorf("ProtonFlux: Expected value: 0, got %v", data.ProtonFlux)
	}
	tests := []struct {
		name     string
		input    interface{}
		expected interface{}
	}{
		{"ProtonFlux", data.ProtonFlux, ProtonFlux(0)},
		{"ElectronFlux", data.ElectronFlux, ElectronFlux(0)},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.input != test.expected {
				t.Errorf("NoRpt in %s field: Expected value %v, got %v", test.name, test.input, test.expected)
			}
		})
	}
}
