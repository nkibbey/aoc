package utils

import (
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
	return FileToStringsBy(fn, "\n")
}

func FileToStringsBy(file, sep string) []string {
	input, _ := ioutil.ReadFile(file)
	s := strings.Split(string(input), sep)
	return s[:len(s)-1]
}
