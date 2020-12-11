package main

import (
	"aoc/utils"
	"fmt"
)

var (
	samp = []int{1, 4, 5, 6, 7, 10, 11, 12, 15, 16, 19, 22}
)

// func combSet(top int64) {
// 	sum := big.NewInt(0)
// 	n := big.NewInt(0)

// 	for i := int64(0); i <= top; i++ {
// 		n.Binomial(top, i)
// 		sum.Add(sum, n)
// 	}
// 	fmt.Println(sum)
// }

// func replaceables(set []int) int {
// 	x := 0

// 	for i := 1; i < len(set)-1; i++ {
// 		if set[i+1]-set[i-1] <= 3 {
// 			x++
// 		}
// 	}
// 	return x
// }
func trib() []int {
	trib := []int{0, 1, 2, 4, 7, 13, 24}
	for i := 7; i < 80; i++ {
		trib = append(trib, trib[i-1]+trib[i-2]+trib[i-3])

	}
	return trib
}

func part2(in []int) {
	tr := trib()
	oneBoys := 1
	s := 1
	for i := 1; i < len(in); i++ {
		diff := in[i] - in[i-1]
		if diff == 1 {
			oneBoys++
		} else if oneBoys > 0 {
			s *= tr[oneBoys]
			oneBoys = 0
		}
	}
	fmt.Println(s)
}
func main() {

	// replaceables(samp)
	// replaceables(utils.FileToIntsBy("sortedd10.txt", "\n"))
	part2(utils.FileToIntsBy("ss2.txt", "\n"))
	part2(utils.FileToIntsBy("sortedd10.txt", "\n"))

}
