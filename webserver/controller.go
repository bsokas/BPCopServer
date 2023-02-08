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
  http.HandleFunc("/", defaultHandler)
  http.ListenAndServe(USE_PORT, nil)
}

func defaultHandler(w http.ResponseWriter, req *http.Request){
  io.WriteString(w, "Endpoint for blood pressure API service!")
}
