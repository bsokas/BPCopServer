package webserver

import (
	"fmt"
	"net/http"
)

const USE_PORT = ":8013" // port for test use

func Start() {
	fmt.Printf("Starting API on port %s\n.........", USE_PORT)

	http.HandleFunc("/bp", bpHandler)

	// TODO should replace with ListenAndServeTLS
	http.ListenAndServe(USE_PORT, nil)
}

func bpHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		// fetch one or many
		if err := SendAllBPReadings(w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case http.MethodPost:
		// add a brand new entry, errors handled within function
		UseCORSHeaders(w) // TODO why does this not work within the handler
		EnterNewBPReading(req, w)
	case http.MethodDelete:
		// delete a reading
	case http.MethodPatch:
		// edit a reading
	case http.MethodOptions:
		UseCORSHeaders(w)
	}
}
