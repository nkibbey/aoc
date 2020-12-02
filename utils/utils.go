package utils

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// GetIntSlice takes csvfile and splits values into an int slice
func GetIntSlice(csvFile string) []int {
	input, _ := ioutil.ReadFile(csvFile)
	is := make([]int, 0)
	for _, val := range strings.Split(string(input), ",") {
		currInt, _ := strconv.Atoi(val)
		is = append(is, currInt)
	}

	return is
}
