package webserver

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

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
	_, err = w.Write(body)

	return err
}

func EnterNewBPReading(req *http.Request, w http.ResponseWriter) {
	if req.Body == nil {
		http.Error(w, "Can not pass empty request body", http.StatusBadRequest)
		return
	}

	defer req.Body.Close()

	decoder := json.NewDecoder(req.Body)
	var reading data.BloodPressureReading
	if err := decoder.Decode(&reading); err != nil && err != io.EOF {
		// TODO worth migrating error handling to a helper
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newId, createErr := data.CreateBPReading(reading)
	if createErr != nil {
		http.Error(w, createErr.Error(), http.StatusInternalServerError)
		return
	}

	// fmt.Printf("Pressure: %d/%d, Heart Rate: %d\n", reading.SystolicMMHg, reading.DiastolicMMHg, reading.HeartRateBpm)
	body, writeErr := json.Marshal(data.CreateSuccessResponse{newId})
	if writeErr != nil {
		http.Error(w, writeErr.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(body)
	if err != nil {
		log.Fatal(err.Error())
	}
}
