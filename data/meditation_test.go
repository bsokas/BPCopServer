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

// TODO
func TestValidateRatingInvalidInput(t *testing.T) {}
