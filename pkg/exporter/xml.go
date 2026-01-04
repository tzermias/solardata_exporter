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
	"strconv"
	"time"
)

// Solar and SolarData represents the XML data retrieved from the RSS feed
type Solar struct {
	XMLName xml.Name  `xml:"solar"`
	Data    SolarData `xml:"solardata"`
}
type SolarData struct {
	Updated        customTimestamp `xml:"updated"`
	SolarFlux      int             `xml:"solarflux"`
	AIndex         int             `xml:"aindex"`
	KIndex         int             `xml:"kindex"`
	KIndexNT       string          `xml:"kindexnt"`
	XRay           XRay            `xml:"xray"`
	Sunspots       Sunspots        `xml:"sunspots"`
	HeliumLine     float32         `xml:"heliumline"`
	ProtonFlux     uint            `xml:"protonflux"`
	ElectronFlux   uint            `xml:"electonflux"`
	Aurora         int             `xml:"aurora"`
	AuroraLatitude float32         `xml:"latdegree"`
	Normalization  float32         `xml:"normalization"`
	SolarWind      float32         `xml:"solarwind"`
	MagneticField  float32         `xml:"magneticfield"`

	CalculatedConditions    []HFCondition  `xml:"calculatedconditions>band"`
	CalculatedVHFConditions []VHFCondition `xml:"calculatedvhfconditions>phenomenon"`
}

// customTimestamp type to parse the timestamp from the XML file.
type customTimestamp struct {
	time.Time
}

type Sunspots uint

// XRayClass represent the different X-Ray classses (A, B, C, M or X)
type XRayClass int

// XRay
type XRay float64

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
	// X-Ray Classes
	A XRayClass = 1
	B XRayClass = 10
	C XRayClass = 100
	M XRayClass = 1000
	X XRayClass = 10000
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

	// X-Ray Classes
	xray_class = map[string]XRayClass{
		"A": A,
		"B": B,
		"C": C,
		"M": M,
		"X": X,
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

// UnmarshalText function to convert X-Ray classes (e.g C1.1, M6.3, X9.9) to the actual scale
func (x *XRay) UnmarshalText(text []byte) error {
	// Ensure that first character is one of A, B, C, M or X
	m := string(text[0])
	n, err := strconv.ParseFloat(string(text[1:]), 64)
	if !(m == "A" || m == "B" || m == "C" || m == "M" || m == "X") || err != nil {
		return errors.New(fmt.Sprintf("Unknown Xray value \"%s\"", text))
	}

	*x = XRay(n * float64(xray_class[m]))
	return nil
}

// UnmarshalText function to convert the custom time format.
func (t *customTimestamp) UnmarshalText(text []byte) error {
	const customFormat = " 02 Jan 2006 1504 GMT"
	str := string(text)
	parsed, err := time.Parse(customFormat, str)
	if err != nil {
		return err
	}
	*t = customTimestamp{parsed}
	return nil
}

// UnmarshalText for Sunspots. In case of a parse error (usually a "NoRpt" string
// in the field), return 0.
func (s *Sunspots) UnmarshalText(text []byte) error {
	sunspots, err := strconv.ParseUint(string(text), 10, 32)
	if err != nil {
		sunspots = 0
	}
	*s = Sunspots(sunspots)

	return nil
}
