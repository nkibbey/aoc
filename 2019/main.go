package main

import (
	"io/ioutil"
	"strconv"
	"strings"

	"aoc/2019/intcode"
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
