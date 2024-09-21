# solardata_exporter

Scrape solar and terrestrial data feed from [N0NBH](https://hamqsl.com/) and expose them as Prometheus metrics.

## Usage

TBD

## Configuration

TBD

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

Metrics exported can be found on the following table. For additiona data please refer to 
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
| solar_hf_condition    | Calculated HF conditions per band and time of day                                                            |
| solar_vhf_condition   | Calculated VHF conditions per phenomenon and location                                                        | 



