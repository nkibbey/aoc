package main

import (
	"bufio"
	"fmt"
	"os"
)

func inputToSlice(fn string) []string {
	var m []string
	file, err := os.Open(fn)
	if err != nil {
		fmt.Println("no bueno open")
		return m
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	for i := 0; s.Scan(); i++ {
		// t := s.Text()
		// fmt.Printf("%v: %s\tlen: %v\n", i, t, len(t))
		m = append(m, s.Text())
	}

	if err := s.Err(); err != nil {
		fmt.Println(err)
		return m
	}

	return m
}

// skiCollisions calculates number of collisions given strail and the slope down it
func skiCollisions(strail []string, slope int) int {
	hits := 0
	for i, curr := range strail {
		pos := (i * slope) % len(curr)
		a := string(curr[pos])
		if a == "#" {
			hits++
		}
	}
	return hits
}

// skiCollisions calculates number of collisions given strail, the halved horizontal slope while only counting even lines
func skiCollisionsJ2(strail []string, slope int) int {
	hits := 0
	for i, curr := range strail {
		if i%2 == 0 {
			pos := (i * slope / 2) % len(curr)
			a := string(curr[pos])
			if a == "#" {
				hits++
			}
		}
	}
	return hits
}

func main() {
	strail := inputToSlice("day3.txt")
	n1, n3, n5, n7 := skiCollisions(strail, 1), skiCollisions(strail, 3), skiCollisions(strail, 5), skiCollisions(strail, 7)
	w1 := skiCollisionsJ2(strail, 1)
	fmt.Println(n1 * n3 * n5 * n7 * w1)
}
