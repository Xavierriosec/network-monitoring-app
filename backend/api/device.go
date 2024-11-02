package api

import (
    "encoding/json"
    "net/http"
)

type Device struct {
    ID   string `json:"id"`
    Name string `json:"name"`
    IP   string `json:"ip"`
}

var devices []Device

func AddDeviceHandler(w http.ResponseWriter, r *http.Request) {
    var device Device
    _ = json.NewDecoder(r.Body).Decode(&device)
    devices = append(devices, device)
    json.NewEncoder(w).Encode(devices)
}
