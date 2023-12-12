package main

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

var (
	gameIDRegex = regexp.MustCompile(`^Game (\d+): `)
)

func SumPossibleGameIds(reader io.Reader, configuration map[string]int) int {
	scanner := bufio.NewScanner(reader)

	var sum int

	for scan := scanner.Scan(); scan; {
		id := PossibleGameIds(scanner.Text(), configuration)

		fmt.Println(id)
		sum += id

		scan = scanner.Scan()
	}

	return sum
}

func SumPowerOfSets(reader io.Reader) int {
	scanner := bufio.NewScanner(reader)

	var sum int

	for scan := scanner.Scan(); scan; {
		s := PowerOfLowestPossibleSet(scanner.Text())

		fmt.Println(s)
		sum += s

		scan = scanner.Scan()
	}

	return sum
}

func PossibleGameIds(i string, configuration map[string]int) (ID int) {
	r := gameIDRegex.FindStringSubmatch(i)
	var err error

	if len(r) == 2 {
		if ID, err = strconv.Atoi(r[1]); err != nil {
			fmt.Printf("Error converting %s to int: %s", r[1], err)
			return 0
		}
	}

	// go trough each set and check if it's possible
	for _, v := range strings.Split(gameIDRegex.ReplaceAllString(i, ""), ";") {
		fmt.Printf("handling: %q\n", v)

		set := CreateSet(strings.Trim(v, " "))

		for color, amount := range configuration {
			if set[color] > amount {
				fmt.Printf("Not possible: %q\n", v)
				return 0
			}
		}
	}

	return ID
}

func PowerOfLowestPossibleSet(i string) (sum int) {
	fewest := map[string]int{}

	// go trough each set and check if it's possible
	for _, v := range strings.Split(gameIDRegex.ReplaceAllString(i, ""), ";") {
		fmt.Printf("handling: %q\n", v)

		for color, amount := range CreateSet(strings.Trim(v, " ")) {
			if _, ok := fewest[color]; !ok {
				fewest[color] = amount
				continue
			}

			if amount > fewest[color] {
				fewest[color] = amount
			}
		}
	}

	sum = 1

	for _, v := range fewest {
		sum *= v
	}

	return sum
}

func CreateSet(i string) map[string]int {
	set := make(map[string]int)

	for _, v := range strings.Split(i, ",") {
		v = strings.Trim(v, " ")
		p := strings.Split(v, " ")

		amount, err := strconv.Atoi(p[0])
		if err != nil {
			fmt.Printf("Error converting %s to int: %s\n", p[0], err)
			return nil
		}

		color := p[1]
		set[color] = amount
	}

	return set
}
