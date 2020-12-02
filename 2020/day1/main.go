package main

import (
	"aoc/utils"
	"fmt"
)

func Find2EltSum(elts []int, sum int) (int, int) {
	return 0, 0
}

func main() {
	a, b := Find2EltSum(utils.GetIntSlice("resources/day1.csv"), 2020)
	fmt.Println(a * b)
}
