package main

import (
	"aoc/utils"
	"fmt"
)

// Find3EltSum finds 3 elts in list that sum to given sum number
func Find3EltSum(elts []int, sum int) (int, int, int) {
	for i := range elts {
		for j := i + 1; j < len(elts); j++ {
			for k := j + 1; k < len(elts); k++ {
				if elts[i]+elts[j]+elts[k] == sum {
					return elts[i], elts[j], elts[k]
				}
			}

		}
	}
	return 0, 0, 0
}

// Find2EltSum finds 2 elts in list that sum to given sum number
func Find2EltSum(elts []int, sum int) (int, int) {
	for i := range elts {
		for j := i + 1; j < len(elts); j++ {
			if elts[i]+elts[j] == sum {
				return elts[i], elts[j]
			}
		}
	}
	return 0, 0
}

func main() {
	a, b := Find2EltSum(utils.GetIntSlice("resources/day1.csv"), 2020)
	fmt.Println(a * b)

	a, b, c := Find3EltSum(utils.GetIntSlice("resources/day1.csv"), 2020)
	fmt.Println(a * b * c)
}
