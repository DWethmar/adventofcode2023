package main

import (
	"bytes"
	"io/ioutil"
	"os"
)

func main() {
	rawBody, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	c := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	println("---")
	println("day 02 part1", SumPossibleGameIds(bytes.NewBuffer(rawBody), c))
	os.Stdin.Seek(0, 0)
	println("day 02 part2", SumPowerOfSets(bytes.NewBuffer(rawBody)))
}
