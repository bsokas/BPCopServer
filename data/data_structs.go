package data

const DBName = "blood_pressure"

// TODO table name list?

type BloodPressureReading struct {
  ID int
  SystolicMMHg int
  DiastolicMMHg int
  HeartRateBpm int
  CreatedAt string
  RecordedAt string
  TripleReading bool
  Notes string
}

type MeditationLog struct {
  ID int
  CreatedAt string
  MeditatedAt string
  DurationSeconds int
  Rating uint8
  Comments string
}
