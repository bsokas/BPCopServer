package webserver

import (
	"strings"
	"regexp"
	"net/http"
)

const LOCAL_APP_ORIGIN = "http://localhost:3000"

// Modifies the time string if it contains datetime delimiters
func ParseJSONRecordedAt(reading *BloodPressureEntryRequest) {
	regex := regexp.MustCompile(`[A-Za-z]`)
	reading.RecordedAt = strings.TrimSpace(regex.ReplaceAllString(reading.RecordedAt, " "))
}

func UseCORSHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Origin")
	w.Header().Set("Access-Control-Allow-Origin", LOCAL_APP_ORIGIN)
}
