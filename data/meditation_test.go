package data

import (
  "testing"
  "strings"
)

func TestValidateRating(t *testing.T) {
  ratingStr := "   3   "
  actual, err := ValidateRating(ratingStr)

  expected := 3
  if err != nil {
    t.Fatal(err.Error())
  } else if actual != expected {
    t.Fatalf("Expected rating value of %d, got %d\n", expected, actual)
  }
}

func TestValidateRatingInvalidRange(t *testing.T) {
  ratingStrHigh := "   6"
  _, err := ValidateRating(ratingStrHigh)
  if err == nil {
    t.Fatalf("Expected invalid range error but got none")
  } else if !strings.Contains(err.Error(), "is invalid") {
    t.Fatalf("Received improper error")
  }

  ratingStrLow := "-2   "
  _, err = ValidateRating(ratingStrLow)
  if err == nil {
    t.Fatalf("Expected invalid range error but got none")
  } else if !strings.Contains(err.Error(), "is invalid") {
    t.Fatalf("Received improper error")
  }
}

func TestValidateRatingInvalidInput(t *testing.T) {
  rating := "abc123"
  _, err := ValidateRating(rating)
  if err == nil {
    t.Fatal("Expected error when passing invalid value")
  }
}

func TestDurationMinToSeconds(t *testing.T) {
  durationStr := "   10   "
  result, err := DurationMinToSeconds(durationStr)
  if err != nil {
    t.Fatal(err)
  }

  expected := 600
  if result != expected {
    t.Fatalf("Expected value of %d, instead received %d\n", expected, result)
  }
}

func TestValidateDurationSecondsNegative(t *testing.T) {
  durationNeg := "-40"
  _, err := DurationMinToSeconds(durationNeg)
  if err == nil {
    t.Fatal("Expected error when passing in a negative input\n")
  } else if !strings.Contains(err.Error(), "Did you only meditate for a few seconds") {
    t.Fatal("Function returned error, but the wrong error")
  }
}
