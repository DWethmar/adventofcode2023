package main

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
)

var regex = regexp.MustCompile(`^.*?(\d)(?:.*(\d))?.*$`)

func CalibrateDocument(reader io.Reader) int {
	scanner := bufio.NewScanner(reader)

	var sum int

	for scan := scanner.Scan(); scan; {
		sum += CalibrationValue(scanner.Text())
		scan = scanner.Scan()
	}

	return sum
}

// CalibrationValue sums the first and last digit of the string
func CalibrationValue(i string) int {
	r := regex.FindStringSubmatch(i)

	var first, last string

	if len(r) == 3 {
		first = r[1]
		if r[2] == "" {
			last = r[1]
		} else {
			last = r[2]
		}
	}

	s, err := strconv.Atoi(first + last)
	if err != nil {
		fmt.Printf("Error converting %s to int: %s", first+last, err)
		return 0
	}

	return s
}
