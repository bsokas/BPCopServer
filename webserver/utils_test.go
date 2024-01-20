package webserver

import (
	"testing"
)

func TestParseJSONRecordedAtValid(t *testing.T) {
	testVal := "2024-01-14T16:45:00.000Z" // 1/14 @ 11:45am EST
	expected := "2024-01-14 16:45:00.000"

	TestEntryRequest := BloodPressureEntryRequest{RecordedAt: testVal}
	ParseJSONRecordedAt(&TestEntryRequest)

	if TestEntryRequest.RecordedAt != expected {
		t.Fatalf("RecordedAt value is %s, expected %s\n", TestEntryRequest.RecordedAt, expected)
	}
}