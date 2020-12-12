package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
)

func move(x, y int, dir string, dist int) (int, int) {
	if dir == "E" { 
		return x + dist, y
	} else if dir == "N" {
		return x, y + dist
	} else if dir == "W" { 
		return x - dist, y
	} else { // south
		return x, y - dist
	}
}

func parseAction(a, pdir string) (dir string, dist int, change bool) {
	dir, dist, change = pdir, 0, false
	if len(a) < 1 {
		fmt.Println("aaaaa")
		return
	}
	d := a[:1]
	amount, _ := strconv.Atoi(a[1:])
	
	if d == "E" || d == "N" || d == "W" || d == "S" {
		dist = amount
		dir = d
		return
	}
	if d == "F" {
		dir = pdir
		dist = amount
		return
	}


	enws := []string{"E", "N", "W", "S"}
	i:= 0
	for j,elt := range enws {
		if elt == pdir {
			i = j
			break
		}
	}
	if d == "L" {
		amount = amount/90
		i = (i + amount) % len(enws)
	} else if d == "R" {
		i = ((16+i)-(amount/90)) % 4 //  -2 mod 4 = -2 so I bumped the minuend
	}
	dir = enws[i]
	dist = 0
	change = true
	return
}


func sail(actions[]string) {
	x, y, face := 0,0,"E"
	for _, action := range actions {
		dir, dist, change := parseAction(action, face) 
		x,y = move(x,y,dir,dist) 
		if change{
			face = dir
		}
	}
	fmt.Println("x: ", x, ", y: ", y)
}


func main() {
	sail(utils.FileToStringsBy("samp.txt", "\n"))
	sail(utils.FileToStringsBy("day12.txt", "\n"))
}
