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
	"log"

	"github.com/mmcdole/gofeed"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	// RSS Feed URL
	feed_url = "https://www.hamqsl.com/solarrss.php"

	// Prometheus namespace
	namespace = "solar"
)

var (
	up = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "up"),
		"Was the last scrape of the feed successful.",
		nil, nil,
	)
	solarflux = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "solarflux"),
		"Solar Flux Index.",
		nil, nil,
	)
	sunspots = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "sunspots"),
		"Sunspot Number.",
		nil, nil,
	)
	aindex = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "aindex"),
		"Planetary A Index.",
		nil, nil,
	)
	kindex = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "kindex"),
		"Planetary K Index.",
		nil, nil,
	)
	protonflux = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "protonflux"),
		"Proton Flux.",
		nil, nil,
	)
	electronflux = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "electronflux"),
		"Electron Flux.",
		nil, nil,
	)
	aurora = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "aurora"),
		"Aurora.",
		nil, nil,
	)
	aurora_lat = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "aurora_latitude"),
		"Aurora Latitude.",
		nil, nil,
	)
	solarwind = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "solarwind"),
		"Solar Wind.",
		nil, nil,
	)

	hf_condition = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "hf_condition"),
		"Calculated HF conditions.",
		[]string{"band_name", "time", "value"}, nil,
	)
	vhf_condition = prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "vhf_condition"),
		"Calculated VHF conditions.",
		[]string{"phenomenon", "location", "value"}, nil,
	)
)

type Exporter struct {
}

func NewExporter() *Exporter {
	return &Exporter{}
}
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- up
	ch <- solarflux
	ch <- sunspots
	ch <- aindex
	ch <- kindex
	ch <- protonflux
	ch <- electronflux
	ch <- aurora
	ch <- aurora_lat
	ch <- solarwind
	ch <- hf_condition
	ch <- vhf_condition
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	var data SolarData

	// Fetch data from the URL and parse them
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(feed_url)
	if err != nil {
		log.Println("Could not fetch RSS data")
		log.Println(err)
		ch <- prometheus.MustNewConstMetric(
			up, prometheus.GaugeValue, 0,
		)
		return
	}
	// Parse XML
	err = xml.Unmarshal([]byte(feed.Items[0].Custom["solar"]), &data)
	if err != nil {
		log.Println("Could not parse XML data from RSS feed")
		log.Println(err)
		ch <- prometheus.MustNewConstMetric(
			up, prometheus.GaugeValue, 0,
		)
		return
	}

	ch <- prometheus.MustNewConstMetric(
		up, prometheus.GaugeValue, 1,
	)

	ch <- prometheus.MustNewConstMetric(
		solarflux, prometheus.GaugeValue, float64(data.SolarFlux),
	)
	ch <- prometheus.MustNewConstMetric(
		sunspots, prometheus.GaugeValue, float64(data.Sunspots),
	)
	ch <- prometheus.MustNewConstMetric(
		aindex, prometheus.GaugeValue, float64(data.AIndex),
	)
	ch <- prometheus.MustNewConstMetric(
		kindex, prometheus.GaugeValue, float64(data.KIndex),
	)
	ch <- prometheus.MustNewConstMetric(
		protonflux, prometheus.GaugeValue, float64(data.ProtonFlux),
	)
	ch <- prometheus.MustNewConstMetric(
		electronflux, prometheus.GaugeValue, float64(data.ElectronFlux),
	)
	ch <- prometheus.MustNewConstMetric(
		aurora, prometheus.GaugeValue, float64(data.Aurora),
	)
	ch <- prometheus.MustNewConstMetric(
		aurora_lat, prometheus.GaugeValue, float64(data.AuroraLatitude),
	)
	ch <- prometheus.MustNewConstMetric(
		solarwind, prometheus.GaugeValue, float64(data.SolarWind),
	)

	// Calculated HF conditions
	for _, condition := range data.CalculatedConditions {
		// Create metric with values
		ch <- prometheus.MustNewConstMetric(
			hf_condition, prometheus.GaugeValue, 1, condition.Band, condition.Time, condition.Value,
		)
	}

	// Calculated VHF conditions
	for _, condition := range data.CalculatedVHFConditions {
		ch <- prometheus.MustNewConstMetric(
			vhf_condition, prometheus.GaugeValue, 1, condition.Phenomenon, condition.Location, condition.Value,
		)
	}
}