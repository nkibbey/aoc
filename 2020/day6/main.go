package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func getSlice(fn string) []string {
	input, err := ioutil.ReadFile(fn)
	if err != nil {
		fmt.Println("asdffasd")
		return []string{}
	}
	return strings.Split(string(input), "\n")
}

// func addUniq(group, indiv string) string{
// 	res 
// 	for i := 0; i < len(group); i++ {
// 		uniq := true
// 		for j := 0; j<len(indiv); j++ {
// 			if indiv[j] == indiv[i] {
// 				uniq = false
// 				continue
// 			}
// 		}
// 		if uniq {

// 		}
// 	}

// 	return ""
// }

func uniqs(ans string) int {
	u := ""
	for i := 0; i < len(ans); i++ {
		if !strings.Contains(u, string(ans[i])) {
			u+=string(ans[i])
		}
	}
	return len(u)
}

func answerSums(answers []string) int {
	sum := 0
	curr := ""
	for _, answer := range answers {
		if answer == "" {
			sum+=uniqs(curr)
			curr = ""
		} else {
			curr += answer
		}
	}
	return sum
}

func main() {
	// fmt.Println(getSlice("samp.txt"))
	fmt.Println(answerSums(getSlice("day6.txt")))
}
