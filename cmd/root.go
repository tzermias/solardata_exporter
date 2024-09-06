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
package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/tzermias/solardata_exporter/pkg/exporter"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
)

// Server port
var port int

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "solardata_exporter",
	Short: "Export HAMQSL.com Solar Data as Prometheus metrics",
	Long: `This exporter, scrapes N0NBH Solar data RSS feed and transforms
it into OpenMetrics format, ready to be scraped by Prometheus`,
	Run: func(cmd *cobra.Command, args []string) {

		// Create new exporter
		exporter := exporter.NewExporter()
		prometheus.MustRegister(exporter)

		http.Handle("/metrics", promhttp.Handler())
		log.Println("Starting exporter ...")
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 9101, "Port where the exporter should listen to")
}
