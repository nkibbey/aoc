package main

import (
	"aoc/utils"
	"fmt"
)

var (
	samp = []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}
)

func pairSumExists(vals []int, sum int) bool {
	// fmt.Println("looking for ", sum, ", amidst ", vals)
	for i := range vals {
		for j := i + 1; j < len(vals); j++ {
			if vals[i]+vals[j] == sum {
				return true
			}
		}
	}

	return false
}

func findPatternBreak(vals []int, psz int) {
	for i := psz; i < len(vals); i++ {
		if !pairSumExists(vals[i-psz:i], vals[i]) {
			fmt.Println(vals[i])
			break
		}
	}
}

func contigSetSums(vals []int, sum int) {
	for i := range vals {
		big := -1
		smol := 20000000000000000
		csum := 0
		for j := i; j < len(vals); j++ {
			curr := vals[j]
			csum += curr
			if curr > big {
				big = curr
			}
			if curr < smol {
				smol = curr
			}
			if csum == sum {
				fmt.Println("Answer is ", (big + smol))
				fmt.Println("this got hit at start ", i, " and end ", j)
				return
			} else if csum > sum {
				break
			}
		}
	}
}

func main() {
	findPatternBreak(samp, 5)
	day9 := utils.FileToIntsBy("day9.txt", "\n")
	findPatternBreak(day9, 25)

	contigSetSums(samp, 127)
	contigSetSums(day9, 27911108)
}
