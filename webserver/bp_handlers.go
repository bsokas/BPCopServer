package webserver

import (
  "net/http"
  "encoding/json"
  "fmt"
  // "io"
  "github.com/bsokas/BPCopServer/data"
)

func SendAllBPReadings(w http.ResponseWriter) error {
  readings, readErr := data.GetBPReadings(true)
  if readErr != nil {
    return readErr
  }

  // TODO might be worth typing this
  respBody := map[string][]data.BloodPressureReading{
    "readings": readings,
  }
  body, err := json.Marshal(respBody)
  if err != nil {
    return err
  }

  // TODO not sure if there's a more effective way to handle CORS
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Header().Set("Content-Type", "application/json")
  written, err := w.Write(body)

  fmt.Printf("Bytes written: %d\n", written)
  return err
}
