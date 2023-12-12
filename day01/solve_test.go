package main

import (
	_ "embed"
	"strings"
	"testing"
)

//go:embed test_input_1.txt
var testInput []byte

func TestCalibrateDocument(t *testing.T) {
	input := strings.NewReader(string(testInput))

	expected := 142
	actual := CalibrateDocument(input)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestCalibrationValue(t *testing.T) {

}
