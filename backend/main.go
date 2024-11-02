package main

import (
    "log"
    "net/http"
    "./api"
)

func main() {
    http.HandleFunc("/api/devices", api.AddDeviceHandler)
    http.HandleFunc("/api/metrics", api.MetricsHandler)

    log.Println("Backend corriendo en el puerto 4000")
    log.Fatal(http.ListenAndServe(":4000", nil))
}
