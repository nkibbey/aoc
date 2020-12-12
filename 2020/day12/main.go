package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
)

func move(x, y, x2, y2 int, dir string, dist int) (int, int, int, int) {
	if dir == "E" {
		return x, y, (x2 + dist), y2
	} else if dir == "N" {
		return x, y, x2, (y2 + dist)
	} else if dir == "W" {
		return x, y, (x2 - dist), y2
	} else if dir == "S" {
		return x, y, x2, (y2 - dist)
	} else {
		x += x2 * dist
		y += y2 * dist
		return x, y, x2, y2
	}
}

func parseAction(a string, x, y int) (dir string, dist int, nx int, ny int) {
	dir, dist, nx, ny = "N", 0, x, y
	if len(a) < 1 {
		fmt.Println("aaaaa")
		return
	}
	d := a[:1]
	amount, _ := strconv.Atoi(a[1:])

	if d == "E" || d == "N" || d == "W" || d == "S" || d == "F" {
		dist = amount
		dir = d
		return
	}
	amount = amount / 90 % 4
	dist = 0
	if d == "L" {
		if amount == 1 {
			amount = 3
		} else if amount == 3 {
			amount = 1
		}
	}

	for i := 0; i < amount; i++ {
		nx = y
		ny = x * -1
		x, y = nx, ny
	}
	return

}

func sail(actions []string) {
	dir, dist, x, y, x2, y2 := "F", 0, 0, 0, 10, 1
	for _, action := range actions {
		dir, dist, x2, y2 = parseAction(action, x2, y2)
		x, y, x2, y2 = move(x, y, x2, y2, dir, dist)
		// fmt.Println("wx: ", x2, ", wy: ", y2)
	}
	fmt.Println("x: ", x, ", y: ", y)
}

func main() {
	sail(utils.FileToStringsBy("samp.txt", "\n"))
	sail(utils.FileToStringsBy("day12.txt", "\n"))
}
