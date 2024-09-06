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
)

// SolarData XML representation.
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

	CalculatedConditions []struct {
		Value string `xml:",cdata"`
		Band  string `xml:"name,attr"`
		Time  string `xml:"time,attr"`
	} `xml:"calculatedconditions>band"`

	CalculatedVHFConditions []struct {
		Value      string `xml:",cdata"`
		Phenomenon string `xml:"name,attr"`
		Location   string `xml:"location,attr"`
	} `xml:"calculatedvhfconditions>phenomenon"`
}
