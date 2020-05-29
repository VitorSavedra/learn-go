package math

import "testing"

const defaultError = "Expected %v but returns %v."

func TestAverage(t *testing.T) {
	expectedValue := 7.28
	value := Average(7.2, 9.9, 6.1, 5.9)

	if value != expectedValue {
		t.Errorf(defaultError, expectedValue, value)
	}
}
