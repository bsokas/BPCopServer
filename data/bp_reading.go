package data

import (
	"fmt"
	"strconv"
	"strings"
)

const MaxBPMMHg = 370

func GetBPReadings(selectAll bool, ids ...[]string) ([]BloodPressureReading , error) {
	// TODO add reading specific values
	rows, readErr := BPDatabase.Query("SELECT * FROM blood_pressure_reading")
	if readErr != nil {
		return nil, readErr
	}

	defer rows.Close()

	readings := make([]BloodPressureReading, 0)
	for rows.Next() {
		var reading BloodPressureReading
		if readErr := rows.Scan(&reading.ID, &reading.SystolicMMHg, &reading.DiastolicMMHg, &reading.HeartRateBpm, &reading.CreatedAt, &reading.RecordedAt, &reading.TripleReading, &reading.Notes); readErr != nil {
			return nil, readErr
		}

		readings = append(readings, reading)
	}

	return readings, nil
}

func ValidatePressure(pressure string) (int, error) {
	mmHg, err := strconv.Atoi(strings.TrimSpace(pressure))
	if err != nil {
		return -1, err
	} else if mmHg >= MaxBPMMHg {
		return mmHg, fmt.Errorf("Input blood pressure value of %d exceeds highest known recorded value", mmHg)
	} else if mmHg < 0 {
		return mmHg, fmt.Errorf("Input blood pressure value of %d is negative and therefore impossible", mmHg)
	}

	return mmHg, nil
}

// TODO: Validate in future for heart rate in danger levels
func ValidateHeartRate(heartRate string) (int, error) {
	bpm, err := strconv.Atoi(strings.TrimSpace(heartRate))
	if err != nil {
		return -1, err
	} else if bpm <= 0 {
		return bpm, fmt.Errorf("You entered a heart rate of %d bpm. Are you dead?", bpm)
	}

	return bpm, nil
}

// TODO probably could add error handling
func ValidateTripleReading(input string) bool {
	formatted := strings.ToLower(strings.TrimSpace(input))

	switch formatted {
	case "y":
		return true
	case "n":
		return false
	}

	fmt.Errorf("input value %s for triple reading is invalid, defaulting to false", formatted)
	return false
}

func ValidateNotes(notes string) string {
	formatted := strings.TrimSpace(notes)
	if len(formatted) > 255 {
		fmt.Println("Provided notes input exceeds character limit. Notes will be truncated to fall inside the limit.")
		formatted = formatted[:255]
	}

	return formatted
}
