package main

import (
	"strings"
	"aoc/utils"
	"fmt"
)

// func march(i, j int, seats[]string, dir string) int {
// 	if dir == "L" {
// 		for j > 0; j-- {
// 			c := string(seats[i][j])
// 			if c:
// 		}
// 	}
// 	return 0
// }

func applyRules(seats []string) []string {
	ns := []string{}
	for i, line := range seats {
		nl := ""
		for j, c := range line {
			c := string(c)
			if c == "." {
				nl += "."
				continue
			}
			target := "#"
			fellas := 0
			if c != "#" && c != "L" {
				fmt.Println("ahhhh")
				return []string{}
			}
			// up down left right
			if j > 0 && string(line[j-1]) == target {
				fellas++
			}
			if j < len(line)-1 && string(line[j+1]) == target {
				fellas++
			}
			if i > 0 && string(seats[i-1][j]) == target {
				fellas++
			}
			if i < len(seats)-1 && string(seats[i+1][j]) == target {
				fellas++
			}

			//upleft
			if j > 0 && i > 0 && string(seats[i-1][j-1]) == target {
				fellas++
			}
			//upright
			if j < len(line)-1 && i > 0 && string(seats[i-1][j+1]) == target {
				fellas++
			}
			//dleft
			if j > 0 && i < len(seats)-1 && string(seats[i+1][j-1]) == target {
				fellas++
			}
			//dright
			if j < len(line)-1 && i < len(seats)-1 && string(seats[i+1][j+1]) == target {
				fellas++
			}
			if c == "#" {
				if fellas >= 4 {
					nl += "L"
				} else {
					nl += "#"
				}
			} else {
				if fellas == 0 {
					nl += "#"
				} else {
					nl += "L"
				}
			}
		}
		ns = append(ns, nl)
	}
	return ns
}

func eq(a, b []string) bool {
    if (a == nil) != (b == nil) { 
        return false; 
    }

    if len(a) != len(b) {
        return false
    }

    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }

    return true
}


func occupied(seats []string) int {
	sum := 0
	for _, l := range seats {
		sum+= len(l)-len(strings.ReplaceAll(l, "#", ""))
	}
	return sum
}

func main() {
	samp := utils.FileToStringsBy("day11.txt", "\n")
	// fmt.Println(samp)
	// prev := []string{}
	y := samp
	for i := 0; i < 1000; i++ {
		// fmt.Println(y)
		tmp := applyRules(y)
		if eq(tmp, y) {
			fmt.Println("yuh")
			break
		}
		y = tmp
	}
	fmt.Println(occupied(y))
}
