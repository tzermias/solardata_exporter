/*
Copyright © 2024 Aris Tzermias

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package exporter

import (
	"encoding/xml"
	"errors"
	"fmt"
)

// SolarData represents the XML data retrieved from the RSS feed
type SolarData struct {
	XMLName        xml.Name `xml:"solardata"`
	Updated        string   `xml:"updated"`
	SolarFlux      int      `xml:"solarflux"`
	AIndex         int      `xml:"aindex"`
	KIndex         int      `xml:"kindex"`
	KIndexNT       string   `xml:"kindexnt"`
	XRay           string   `xml:"xray"`
	Sunspots       uint     `xml:"sunspots"`
	HeliumLine     float32  `xml:"heliumline"`
	ProtonFlux     uint     `xml:"protonflux"`
	ElectronFlux   uint     `xml:"electonflux"`
	Aurora         int      `xml:"aurora"`
	AuroraLatitude float32  `xml:"latdegree"`
	Normalization  float32  `xml:"normalization"`
	SolarWind      float32  `xml:"solarwind"`
	MagneticField  float32  `xml:"magneticfield"`

	CalculatedConditions    []HFCondition  `xml:"calculatedconditions>band"`
	CalculatedVHFConditions []VHFCondition `xml:"calculatedvhfconditions>phenomenon"`
}

// HFCondition represents an element of `<calculatedhfconditions>` list in the XML data
type HFCondition struct {
	Value HFStatus `xml:",cdata"`
	Band  string   `xml:"name,attr"`
	Time  string   `xml:"time,attr"`
}

// HFStatus
type HFStatus int

const (
	//HF status
	Poor HFStatus = iota
	Fair
	Good
)

// VHFCondition represents an element of `<calculatedvhfconditions>` list in the XML data
type VHFCondition struct {
	Value      VHFStatus `xml:",cdata"`
	Phenomenon string    `xml:"name,attr"`
	Location   string    `xml:"location,attr"`
}

// VHFStatus enum
type VHFStatus int

const (
	//VHF status
	BandClosed VHFStatus = iota
	HighMUF
	Es50MHz
	Es70MHz
	Es144MHz
	MidLatAur
	HighLatAur
)

var (
	// Distinct VHF Conditions reported by the feed.
	// The full list of values is collected from https://www.hamqsl.com/shortcut.html
	vhf_status = map[string]VHFStatus{
		"Band Closed":  BandClosed,
		"High MUF":     HighMUF,
		"50MHz ES":     Es50MHz,
		"70MHz ES":     Es70MHz,
		"144MHz ES":    Es144MHz,
		"MID LAT AUR":  MidLatAur,
		"High LAT AUR": HighLatAur,
	}

	// Distinct HF Conditions reported by the feed.
	// Reported values are the following (using trial and error.)
	hf_status = map[string]HFStatus{
		"Poor": Poor,
		"Fair": Fair,
		"Good": Good,
	}
)

// UnmarshalText function to map XML <phenomenon> values to enumerated VHF conditions using vhf_status map.
func (s *VHFStatus) UnmarshalText(text []byte) error {
	str := string(text)

	status, ok := vhf_status[str]
	if !ok {
		return errors.New(fmt.Sprintf("Unknown VHF Status \"%s\"", str))
	}

	*s = status
	return nil
}

// UnmarshalText function to map XML <band> values to enumerated HF conditions using hf_status map.
func (s *HFStatus) UnmarshalText(text []byte) error {
	str := string(text)

	status, ok := hf_status[str]
	if !ok {
		return errors.New(fmt.Sprintf("Unknown HF Status \"%s\"", str))
	}

	*s = status
	return nil
}
