package main

import (
	"aoc/utils"
	"errors"
	"fmt"
	"strconv"
)

type hh struct {
	index int
	accum int
	code  string
}

func (h hh) next(codes []string) (hh, error) {
	var next hh
	if len(h.code) < 6 {
		fmt.Println("I think eof")
		return h, errors.New("-1")
	}
	op := h.code[:3]
	arg, err := strconv.Atoi(h.code[4:])
	if err != nil {
		fmt.Println("ahhhh")
		return next, errors.New("-2")
	}
	// fmt.Println("op: ", op, ", arg: ", arg)
	nextIndex := h.index
	nextAccum := h.accum
	if op == "nop" {
		nextIndex++
	} else if op == "jmp" {
		nextIndex += arg
	} else if op == "acc" {
		nextAccum += arg
		nextIndex++
	}
	if nextIndex >= 0 && nextIndex < len(codes) {
		return hh{nextIndex, nextAccum, codes[nextIndex]}, nil
	}
	fmt.Println("oob: accum: ", nextAccum)
	return next, errors.New("hmm")
}

func alreadyProcessed(curr hh, processed []hh) bool {
	for _, elt := range processed {
		if curr.index == elt.index {
			// fmt.Println("loop")
			return true
		}
	}
	return false
}

func processCodes(codes []string) bool {
	prev := hh{0, 0, codes[0]}
	var proccesed []hh
	var curr hh
	var err error
	for err == nil {
		curr, err = prev.next(codes)
		if alreadyProcessed(curr, proccesed) {
			fmt.Println(curr)
			return false
		}
		proccesed = append(proccesed, curr)
		prev = curr
	}
	fmt.Println(curr)
	return true
}

func changeIJumpToNop(i int, codes []string) []string {
	js := 0
	ncodes := []string{}
	for index, code := range codes {
		if len(code) == 0 {
			return ncodes
		}
		ncodes = append(ncodes, code)
		if code[:3] == "jmp" {
			if js == i {
				ncodes[index] = "nop -1"
			}
			js++
		}
	}
	return ncodes
}

func main() {
	samp := utils.GetStringSlice("samp.txt")
	t := processCodes(samp)
	// fmt.Println(t)
	// t = processCodes(utils.GetStringSlice("sampc.txt"))
	// fmt.Println(t)
	day8 := utils.GetStringSlice("day8.txt")
	t = processCodes(day8)
	fmt.Println(t)
	for i := 0; i < 229; i++ {
		nsamp := changeIJumpToNop(i, day8)
		if processCodes(nsamp) {
			fmt.Println(nsamp)
			return
		}
	}
}
