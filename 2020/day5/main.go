package main

import (
	"aoc/utils"
	"fmt"
	"math"
	"sort"
)

func verifyInput(min, max int, dirs string) {
	if min > max || max < min {
		fmt.Printf("Ahh min: %d, max: %d, dirs: %s", min, max, dirs)
	}
}

func seatFind(min, max int, dirs string) int {
	// fmt.Println(dirs)
	// fmt.Printf("min: %d\tmax: %d\n", min, max)

	if len(dirs) == 0 {
		return min
	}
	verifyInput(min, max, dirs)
	dir := string(dirs[0])
	if dir == "F" || dir == "L" {
		max = min + ((max - min) / 2)
		return seatFind(min, max, dirs[1:])
	} else if dir == "B" || dir == "R" {
		// fmt.Println(float64(float64(max-min) / float64(2)))
		min = min + int(math.Ceil(float64(float64(max-min)/float64(2))))
		return seatFind(min, max, dirs[1:])
	}

	return -1
}

func PrintSeatingsAndIds(dirs []string) {
	for _, dir := range dirs {
		if len(dir) == 10 {
			row := seatFind(0, 127, dir[:7])
			col := seatFind(0, 7, dir[7:])
			fmt.Printf(" - %s: row %d, column %d, seat ID %d\n", dir, row, col, (row*8 + col))
		}
	}
}

func highestId(dirs []string) int {
	h := -1

	for _, dir := range dirs {
		if len(dir) == 10 {
			row := seatFind(0, 127, dir[:7])
			col := seatFind(0, 7, dir[7:])
			id := row*8 + col
			if id > h {
				h = id
			}
		}
	}

	return h
}

func getIds(dirs []string) []int {
	ids := []int{}

	for _, dir := range dirs {
		if len(dir) == 10 {
			row := seatFind(0, 127, dir[:7])
			col := seatFind(0, 7, dir[7:])
			id := row*8 + col
			ids = append(ids, id)
		}
	}

	return ids
}

func main() {
	// fmt.Println(seatFind(0, 7, "RLR"))
	// fmt.Println(seatFind(0, 127, "FBFBBFFRLR"))
	// fmt.Println("should be 5 44")

	// fmt.Println(seatFind(0, 127, "BFFFBBF"))
	// fmt.Println(seatFind(0, 7, "RRR"))
	// fmt.Println("should be 70 7")

	// fmt.Println(seatFind(0, 127, "FFFBBBF"))
	// fmt.Println(seatFind(0, 7, "RRR"))
	// fmt.Println("should be 14 7")

	// fmt.Println(seatFind(0, 127, "BBFFBBF"))
	// fmt.Println(seatFind(0, 7, "RLL"))
	// fmt.Println("should be 102 4")
	dirs := utils.GetStringSlice("day5.txt")
	// PrintSeatingsAndIds(dirs)
	// fmt.Println(highestId(dirs))
	ids := getIds(dirs)
	// fmt.Println(ids)
	sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })
	fmt.Println(ids)
	fmt.Println(len(ids))
}
