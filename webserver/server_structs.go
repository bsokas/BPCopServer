package webserver

type BloodPressureEntryRequest struct {
	SystolicMMHg  int    `json:"systolicMMHg,string"`
	DiastolicMMHg int    `json:"diastolicMMHg,string"`
	HeartRateBpm  int    `json:"heartRateBpm,string"`
	RecordedAt    string `json:"recordedAt"`
	TripleReading bool   `json:"tripleReading"`
	Notes         string `json:"notes"`
}