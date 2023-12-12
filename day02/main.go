package main

import "os"

func main() {
	sumOfIds := SumPossibleGameIds(os.Stdin, map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	})

	println(sumOfIds)
}
