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

func makeUnanimous(group, indiv string, first bool) string{
	res := ""
	if first {
		return indiv
	}
	for i := 0; i<len(indiv); i++ {
		if strings.Contains(group, string(indiv[i])) {
			res += string(indiv[i])
		}
	}

	return res
}

func uniqs(ans string) int {
	u := ""
	for i := 0; i < len(ans); i++ {
		if !strings.Contains(u, string(ans[i])) {
			u+=string(ans[i])
		}
	}
	return len(u)
}

// func answerSums(answers []string) int {
// 	sum := 0
// 	curr := ""
// 	for _, answer := range answers {
// 		if answer == "" {
// 			sum+=uniqs(curr)
// 			curr = ""
// 		} else {
// 			curr += answer
// 		}
// 	}
// 	return sum
// }

func answerSums(answers []string) int {
	sum := 0
	curr := ""
	first := true
	for _, answer := range answers {
		if answer == "" {
			fmt.Println(curr)
			sum+=len(curr)
			curr = ""
			first = true
		} else {
			curr = makeUnanimous(curr, answer, first)
			first = false
		}
	}
	return sum
}

func main() {
	// fmt.Println(getSlice("samp.txt"))
	fmt.Println(answerSums(getSlice("day6.txt")))
}
