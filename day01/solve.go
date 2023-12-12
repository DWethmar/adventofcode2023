package main

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
)

var (
	regex  = regexp.MustCompile(`^.*?(\d|one|two|three|four|five|six|seven|eight|nine)(?:.*(\d|one|two|three|four|five|six|seven|eight|nine))?.*$`)
	regex2 = regexp.MustCompile(`\d`)
)

func CalibrateDocument(reader io.Reader) int {
	scanner := bufio.NewScanner(reader)

	var sum int

	for scan := scanner.Scan(); scan; {
		sum += CalibrationValue(scanner.Text())
		scan = scanner.Scan()
	}

	return sum
}

// CalibrationValue creates a digit from the first and last digits of a string.
func CalibrationValue(i string) int {
	r := regex.FindStringSubmatch(i)

	var first, last string

	if len(r) == 3 {
		first = ValueToDigit(r[1])
		if r[2] == "" {
			last = first
		} else {
			last = ValueToDigit(r[2])
		}
	}

	s, err := strconv.Atoi(first + last)
	if err != nil {
		fmt.Printf("Error converting %s to int: %s", first+last, err)
		return 0
	}

	return s
}

func ValueToDigit(i string) string {
	switch i {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	}

	return i
}
