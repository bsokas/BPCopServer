package webserver

import (
  "net/http"
  "fmt"
  "io"
)

const USE_PORT = ":8013" // port for test use

func Start() {
  // TODO should replace with ListenAndServeTLS
  fmt.Printf("Starting API on port %s\n.........", USE_PORT)

  // http.HandleFunc("/", defaultHandler)
  http.HandleFunc("/bp", bpHandler)

  http.ListenAndServe(USE_PORT, nil)
}

func defaultHandler(w http.ResponseWriter, req *http.Request){
  io.WriteString(w, "Endpoint for blood pressure API service!")
}

func bpHandler(w http.ResponseWriter, req *http.Request) {
  switch req.Method {
  case http.MethodGet:
    // fetch one or many
    if err := SendAllBPReadings(w); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
    }
  case http.MethodPost:
    // add a brand new entry
  case http.MethodDelete:
    // delete a reading
  case http.MethodPatch:
    // edit a reading
  }
}
