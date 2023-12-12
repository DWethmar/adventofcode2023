package main

import (
	"strings"
	"testing"

	_ "embed"
)

//go:embed test_input_1.txt
var testInput []byte

//go:embed input.txt
var testInput2 []byte

func TestSumPossibleGameIds(t *testing.T) {
	input := strings.NewReader(string(testInput))

	expected := 8
	actual := SumPossibleGameIds(input, map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	})

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestSumPowerOfSets(t *testing.T) {
	input := strings.NewReader(string(testInput2))

	expected := 2286
	actual := SumPowerOfSets(input)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestPossibleGameIds(t *testing.T) {

}
