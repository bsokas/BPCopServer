package data

const DBName = "blood_pressure"

// TODO table name list?

type BloodPressureReading struct {
	ID            int `json:"id"`
	SystolicMMHg  int	`json:"systolicMMHg"`
	DiastolicMMHg int `json:"diastolicMMHg"`
	HeartRateBpm  int `json:"heartRateBpm"`
	CreatedAt     string `json:"createdAt"`
	RecordedAt    string `json:"recordedAt"`
	TripleReading bool `json:"tripleReading"`
	Notes         string `json:"notes"`
}

type MeditationLog struct {
	ID              int
	CreatedAt       string
	MeditatedAt     string
	DurationSeconds int
	Rating          uint8
	Comments        string
}
