package config

import (
    "github.com/influxdata/influxdb1-client/v2"
    "log"
)

func ConnectToInflux() client.Client {
    influx, err := client.NewHTTPClient(client.HTTPConfig{
        Addr: "http://influxdb:8086",
    })
    if err != nil {
        log.Fatal(err)
    }
    return influx
}
