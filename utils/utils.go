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
		val = strings.TrimSuffix(val, "\n")
		currInt, _ := strconv.Atoi(val)
		is = append(is, currInt)
	}
	if separator == "\n" {
		return is[:len(is)-1]
	}
	return is
}

// GetStringSlice takes file where slice is made from splitting on newlines
func GetStringSlice(fn string) []string {
	return FileToStringsBy(fn, "\n")
}

func FileToStringsBy(file, sep string) []string {
	input, _ := ioutil.ReadFile(file)
	s := strings.Split(string(input), sep)
	if sep == "\n" {
		return s[:len(s)-1]
	}
	s[len(s)-1] = strings.TrimSuffix(s[len(s)-1], "\n")
	return s
}
