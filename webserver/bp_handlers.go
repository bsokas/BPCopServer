package webserver

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"fmt"

	"github.com/bsokas/BPCopServer/data"
)

type BloodPressurePostBody struct {
	SystolicMMHg int
	DiastolicMMHg int
	HeartRateBpm int 
	TripleReading bool
	Notes string
	RecordedAt string
}

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
	var newReading BloodPressureEntryRequest
	if err := decoder.Decode(&newReading); err != nil && err != io.EOF {
		// TODO worth migrating error handling to a helper
		fmt.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ParseJSONRecordedAt(&newReading)

	newId, createErr := data.CreateBPReading(data.BloodPressureReading{
		SystolicMMHg: newReading.SystolicMMHg,
		DiastolicMMHg: newReading.DiastolicMMHg,
		HeartRateBpm: newReading.HeartRateBpm,
		RecordedAt: newReading.RecordedAt,
		TripleReading: newReading.TripleReading,
		Notes: newReading.Notes,
	})

	if createErr != nil {
		http.Error(w, createErr.Error(), http.StatusInternalServerError)
		return
	}

	body, writeErr := json.Marshal(data.CreateSuccessResponse{newId})
	if writeErr != nil {
		http.Error(w, writeErr.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err := w.Write(body)
	if err != nil {
		log.Fatal(err.Error())
	}
}
