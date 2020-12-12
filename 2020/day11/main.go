package main

import (
	"aoc/utils"
	"fmt"
	"strings"
)

func march(i, j int, seats []string, dir string) int {
	if dir == "L" {
		for j > 0 {
			j--
			c := string(seats[i][j])
			if c == "#" {
				return 1
			} else if c == "L" {
				return 0
			}
		}
	} else if dir == "UL" {
		for j > 0 && i > 0 {
			i--
			j--
			c := string(seats[i][j])
			if c == "#" {
				return 1
			} else if c == "L" {
				return 0
			}
		}
	} else if dir == "U" {
		for i > 0 {
			i--
			c := string(seats[i][j])
			if c == "#" {
				return 1
			} else if c == "L" {
				return 0
			}
		}
	} else if dir == "UR" {
		for j < len(seats[0])-1 && i > 0 {
			i--
			j++
			c := string(seats[i][j])
			if c == "#" {
				return 1
			} else if c == "L" {
				return 0
			}
		}
	} else if dir == "R" {
		for j < len(seats[0])-1 {
			j++
			c := string(seats[i][j])
			if c == "#" {
				return 1
			} else if c == "L" {
				return 0
			}
		}
	} else if dir == "DR" {
		for j < len(seats[0])-1 && i < len(seats)-1 {
			i++
			j++
			c := string(seats[i][j])
			if c == "#" {
				return 1
			} else if c == "L" {
				return 0
			}
		}
	} else if dir == "D" {
		for i < len(seats)-1 {
			i++
			c := string(seats[i][j])
			if c == "#" {
				return 1
			} else if c == "L" {
				return 0
			}
		}
	} else if dir == "DL" {
		for j > 0 && i < len(seats) -1 {
			i++
			j--
			c := string(seats[i][j])
			if c == "#" {
				return 1
			} else if c == "L" {
				return 0
			}
		}
	} 

	return 0
}

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
			fellas := 0
			if c != "#" && c != "L" {
				fmt.Println("ahhhh")
				return []string{}
			}
			fellas+= march(i,j,seats, "L")
			fellas+= march(i,j,seats, "UL")
			fellas+= march(i,j,seats, "U")
			fellas+= march(i,j,seats, "UR")
			fellas+= march(i,j,seats, "R")
			fellas+= march(i,j,seats, "DR")
			fellas+= march(i,j,seats, "D")
			fellas+= march(i,j,seats, "DL")
			
			if c == "#" {
				if fellas >= 5 {
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
		return false
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
		sum += len(l) - len(strings.ReplaceAll(l, "#", ""))
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
