package main

import "fmt"

var (
	samp = []int{0, 3, 6}
	day15 = []int{18,11,9,0,5,1}
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

func main() {
	s := keepGoing(2222, samp)
	fmt.Println(s[2019])
	d := keepGoing(2222, day15)
	fmt.Println(d[2019])
}
