package utils

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// GetIntSlice takes csvfile and splits values into an int slice
func GetIntSlice(csvFile string) []int {
	return FileToIntsBy(csvFile, ",")
}

func FileToIntsBy(file, separator string) []int {
	input, _ := ioutil.ReadFile(file)
	var is []int
	for _, val := range strings.Split(string(input), separator) {
		currInt, _ := strconv.Atoi(val)
		is = append(is, currInt)
	}
	return is[:len(is)-1]
}

// GetStringSlice takes file where slice is made from splitting on newlines
func GetStringSlice(fn string) []string {
	input, err := ioutil.ReadFile(fn)
	if err != nil {
		fmt.Println("asdffasd")
		return []string{}
	}
	return strings.Split(string(input), "\n")
}
