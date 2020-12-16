package main

import "fmt"

var (
	samp  = []int{0, 3, 6}
	day15 = []int{18, 11, 9, 0, 5, 1}
)

func numBoy(nums []int) int {
	topI := len(nums) - 1
	target := nums[topI]
	for i := topI - 1; i >= 0; i-- {
		if nums[i] == target {
			return topI - i
		}
	}
	return 0
}

func keepGoing(cap int, start []int) []int {
	nums := start
	for i := 0; i < cap; i++ {
		nums = append(nums, numBoy(nums))
	}

	return nums
}

func hardcodeboys() {
	// day15 = []int{18,11,9,0,5,1}
	var tmp int
	x := map[int]int{
		18: 1,
		11: 2,
		9:  3,
		0:  4,
		5:  5,
		1:  6,
	}
	// x := map[int]int{
	// 	0: 1,
	// 	3: 2,
	// 	6: 3,
	// }

	curr := 0
	for i := 7; i <= 30000000; i++ {
		if i == 30000000 {
			fmt.Println(curr)
		}
		if _, ok := x[curr]; !ok {
			x[curr] = i
			curr = 0
			continue
		}
		tmp = i - x[curr]
		x[curr] = i
		curr = tmp
	}

}

func main() {
	s := keepGoing(2222, samp)
	fmt.Println(s[2019])
	// d := keepGoing(30000006, day15)
	// fmt.Println(d[(30000000 - 1)])
	hardcodeboys()
}
