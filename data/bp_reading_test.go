package data

import (
  "testing"
  "strings"
)

func TestValidatePressure(t *testing.T) {
  validSystolic := "   120   "
  val, err := ValidatePressure(validSystolic)
  if err != nil {
    t.Fatalf(err.Error())
  }

  expected := 120
  if val != expected {
    t.Fatalf("Expected %d, got %d", expected, val)
  }
}

func TestValidatePressureNegative(t *testing.T) {
  negSystolic := "-120   "
  _, err := ValidatePressure(negSystolic)
  if err == nil {
    t.Fatal("Expected error for inputting negative value")
  } else if !strings.Contains(err.Error(), "Input blood pressure value of -120 is negative") {
    t.Fatalf("Expected error to flag negative value, instead errored on: %s\n", err.Error())
  }
}

func TestValidatePressureNonNumber(t *testing.T) {
  invalidBp := "This is just a string"
  _, err := ValidatePressure(invalidBp)

  if err == nil {
    t.Fatal("Expected invliad value to trigger errors but it did not")
  }
}

func TestValidateHeartRate(t *testing.T) {
  bpmStr := "  61  "
  val, err := ValidateHeartRate(bpmStr)
  if err != nil {
    t.Fatal(err.Error())
  }

  expected := 61
  if val != expected {
    t.Fatalf("Expected heart rate %d, but got %d", expected, val)
  }
}

func TestValidateHeartRateZero(t *testing.T) {
  hr := "0\n"
  _, err := ValidateHeartRate(hr)
  if err == nil {
    t.Fatal("Expected error when parsing HR of 0")
  } else if !strings.Contains(err.Error(), "Are you dead?") {
    t.Fatalf("Function returned unexpected error: %s\n", err.Error())
  }
}

func TestValidateHeartRateInvalid(t *testing.T) {
  hr := "just some words\n\n"
  _, err := ValidateHeartRate(hr)
  if err == nil {
    t.Fatal("Expected error when parsing invalid input")
  }
}
