package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
)

var (
	// 939
	samp      = []string{"7", "13", "x", "x", "59", "x", "31", "19"}
	smartsamp = map[int]int{7: -1, 13: 0, 59: 3, 31: 5, 19: 6}
)

func part1(t int, ids []string) {
	lwait := 5000000
	res := 0
	for _, id := range ids {
		if id == "x" {
			continue
		}
		i, _ := strconv.Atoi(id)
		twait := i - (t % i)
		if twait < lwait {
			lwait = twait
			res = twait * i
		}
	}
	fmt.Println(res)
}

func part2(first, offset int, ids map[int]int) {
	// res := 0
	for t := 0; t < 200000000000000; t += first {
		yuh := true
		for key, val := range ids {
			if ((t + val) % key) != 0 {
				yuh = false
				break
			}
		}
		if yuh {
			fmt.Println(t - offset)
			return
		}
	}

	// for i, id := range ids {
	// 	if id == "x" {
	// 		continue
	// 	}
	// 	i, _ := strconv.Atoi(id)
	// 	twait := i-(t%i)
	// 	if twait < lwait {
	// 		lwait = twait
	// 		res = twait * i
	// 	}
	// }
	// fmt.Println(res)
}

func hardCodeBoys() {
	// smart13 := map[int]int{41: -41, 37: -6, 911: 0, 13: 13, 17: 14, 23: 23, 29: 29, 827: 31, 19: 50}
	a := 0
	for i := 2000000000; i < 2000000000000000; i++ {
		a = 911 * i
		if i%100000000000 == 0 {
			fmt.Println("not yet ", i)
		}
		if (a+31)%827 != 0 {
			continue
		}

		if (a-41)%41 != 0 {
			continue
		}

		if (a-6)%37 != 0 {
			continue
		}

		if (a+29)%29 != 0 {
			continue
		}

		if (a+23)%23 != 0 {
			continue
		}

		if (a+50)%19 != 0 {
			continue
		}

		if (a+14)%17 != 0 {
			continue
		}

		if (a+13)%13 != 0 {
			continue
		}

		fmt.Println(a)
		return
	}

}

func printy(s []string) {
	fmt.Printf("map[int]int{")
	for i, elt := range s {
		if elt == "x" {
			continue
		}
		fmt.Printf("%s: %d, ", elt, i)
	}
	fmt.Printf("\n}\n")
}

func main() {
	part1(939, samp)
	day13 := utils.FileToStringsBy("day13.txt", ",")
	// smart13 := map[int]int{41: -41, 37: -6, 911: 0, 13: 13, 17: 14, 23: 23, 29: 29, 827: 31, 19: 50}
	part1(1011416, day13)

	part2(13, 1, smartsamp)
	hardCodeBoys()
	// printy(day13)
	// part2(911, 41, smart13)
	// part2(7, samp)
}
