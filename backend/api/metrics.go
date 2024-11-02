package api

import (
    "encoding/json"
    "net/http"
)

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
    metrics := map[string]string{"cpu": "20%", "memory": "512MB"}
    json.NewEncoder(w).Encode(metrics)
}
