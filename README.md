# solardata_exporter

[![Go Report Card](https://goreportcard.com/badge/github.com/tzermias/solardata_exporter)](https://goreportcard.com/report/github.com/tzermias/solardata_exporter)

Scrape solar and terrestrial data feed from [N0NBH](https://hamqsl.com/) and expose them as Prometheus metrics.

## Installation and Usage

### Docker

The easiest way to spin up the exporter is by running it as a Docker container

```bash
docker run -d \
  -p 9101:9101 \
  --name solardata_exporter \
  tzermias/solardata_exporter:latest
```

If you use Docker compose, put the following in your `docker-compose.yaml` file
```yaml
---
version: 3

services:
	solardata_exporter:
		image: tzermias/solardata_exporter:latest
		container_name: solardata_exporter
		restart: unless-stopped
		ports:
		- "9101:9101"
```

## Configuration

The exporter listens to port `9101` by default. To override the port where the exporter listens to, run the exporter with `-p <PORT>` flag, where `<PORT>` indicates your new port.

## Prometheus Configuration

Assuming that `solardata_exporter` runs on the same instance as Prometheus listening to default port (`9101`), use the following example configuration:
```
scrape_configs:
 - job_name: 'solardata'
   static_configs:
     - targets: ['127.0.0.1:9101']
   scrape_interval: 120s
```

## Metrics

Metrics exported can be found on the following table. For additional data please refer to 
[this page](https://www.hamqsl.com/solar2.html#usingdata) on N0NBH site.

| Name                  | Description                                                                                                  |
| ------                | -----------                                                                                                  |
| solar_up              | Last scrape from N0NBH feed was successful/unsuccessful                                                      |
| solar_solarflux       | Solar Flux Index (SFI)                                                                                       |
| solar_sunspots        | Sunspot Number                                                                                               |
| solar_aindex          | Planetary A Index                                                                                            |
| solar_kindex          | Planetary K Index                                                                                            |
| solar_xrays		| Solar X-Rays												       |
| solar_protonflux      | Proton Flux                                                                                                  |
| solar_electronflux    | Electron Flux                                                                                                |
| solar_aurora          | Aurora                                                                                                       |
| solar_aurora_latitude | Aurora Latitude                                                                                              |
| solar_solarwind       | Solar Wind                                                                                                   |
| solar_magneticfield   | Magnetic Field                                                                                               |
| solar_hf_condition    | Calculated HF conditions per band and time of day. Check below for value mappings.                                                             |
| solar_vhf_condition   | Calculated VHF conditions per phenomenon and location. Check below for value mappings.                                                        | 

### solar_hf_condition
Mapping for this metric is the following:
| Value | Mapped Condition |
| ----- | ---------------- |
| 0 | *Poor* |
| 1 | *Fair* |
| 2 | *Good* |

More information can be found at [this page](https://www.hamqsl.com/solar2.html#usingdata).

### solar_hf_condition
Mapping for this metric is the following:
| Value | Mapped Condition |
| ----- | ---------------- |
| 0 | *Band Closed* (applicable to all locations) |
| 1 | *High MUF* (applicable only when `location` is `europe` or `north_america`) |
| 2 | *50MHz ES* E-sporadic on 6-meters band (only when `location` is `europe_6m`) |
| 3 | *70MHz ES* E-sporadic on 4-meters band (only when `location` is `europe_4m`) |
| 4 | *144MHz ES* E-sporadic on 2-meters band (only when `location` is `europe`)|
| 5 | *MID LAT AUR* Auroral activity between 60° and 30°N. (applicable when `phenomenon` is `vhf-aurora` only) |
| 6 | *High LAT AUR* Auroral activity greater than 60°N.  (applicable when `phenomenon` is `vhf-aurora` only) |

More information can be found at [this page](https://www.hamqsl.com/solar2.html#usingdata).
