package main

import (
	"github.com/nkibbey/aoc/intcode"
	"io/ioutil"
	"strconv"
	"strings"
)

func getIntArr(filepath string) []int {
	input, _ := ioutil.ReadFile(filepath)
	intArr := make([]int, 0)
	for _, val := range strings.Split(string(input), ",") {
		currInt, _ := strconv.Atoi(val)
		intArr = append(intArr, currInt)
	}

	return intArr
}

func main() {
	intcode.ParseOpcode(getIntArr("resources/day2.txt"))
}
