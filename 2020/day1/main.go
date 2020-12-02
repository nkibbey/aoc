package main

import (
	"aoc/utils"
	"fmt"
)

func Find2EltSum(elts []int, sum int) (int, int) {
	for i := range elts {
		for j := i + 1; j < len(elts); j++ {
			if elts[i] + elts[j] == sum {
				return elts[i], elts[j]
			}
		}
	}
	return 0, 0
}

func main() {
	a, b := Find2EltSum(utils.GetIntSlice("resources/day1.csv"), 2020)
	fmt.Println(a * b)
}
