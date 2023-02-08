package data

import (
	"fmt"
	"strconv"
	"strings"
)

func UpsertMeditationLog(meditatedAt string, rating int, durationSeconds int, comments string) (int64, error) {
	result, err := BPDatabase.Exec(`INSERT INTO meditation_log
    (meditated_at, rating, duration_seconds, comments)
    VALUES (?, ?, ?, ?)
    `, meditatedAt, rating, durationSeconds, comments)

	if err != nil {
		return -1, err
	}

	return result.LastInsertId()
}

func ValidateRating(ratingStr string) (int, error) {
	rating, err := strconv.Atoi(strings.TrimSpace(ratingStr))
	if err != nil {
		return -1, err
	} else if rating <= 0 || rating > 5 {
		return rating, fmt.Errorf("Rating of %d is invalid, requires a value from 1-5", rating)
	}

	return rating, nil
}

func DurationMinToSeconds(durationMinStr string) (int, error) {
	minInt, err := strconv.Atoi(strings.TrimSpace(durationMinStr))
	if err != nil {
		return -1, err
	} else if minInt < 1 {
		return minInt, fmt.Errorf("Entered duration value less than one. Did you only meditate for a few seconds?")
	}

	durationSeconds := minInt * 60
	return durationSeconds, nil
}

func ValidateComments(comments string) string {
	formatted := strings.TrimSpace(comments)
	if len(comments) > 255 {
		fmt.Println("Meditation comments exceed character limit and will be shortened.")
		formatted = formatted[:255]
	}

	return formatted
}
