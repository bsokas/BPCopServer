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

  body, err := json.Marshal(readings)
  if err != nil {
    return err
  }

  w.Header().Set("Content-Type", "application/json")
  written, err := w.Write(body)

  fmt.Printf("Bytes written: %d\n", written)
  return err
}
