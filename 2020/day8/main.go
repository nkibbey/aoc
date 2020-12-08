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
		fmt.Println("ahh")
		return next, errors.New("-1")
	}
	op := h.code[:3]
	arg, err := strconv.Atoi(h.code[4:])
	if err != nil {
		fmt.Println("ahhhh")
		return next, errors.New("-2")
	}
	// fmt.Println("op: ", op, ", arg: ", arg)
	if op == "nop" {
		return hh{h.index+1, h.accum, codes[h.index+1]}, nil
	} else if op == "jmp" {
		return hh{h.index+arg, h.accum,  codes[h.index+arg]}, nil
	} else if op == "acc" {
		return hh{h.index+1, h.accum+arg, codes[h.index+1]}, nil
	} 
	return next, errors.New("hmm")
}

func alreadyProcessed(curr hh, processed []hh) bool {
	for _, elt := range processed {
		if curr.index == elt.index {
			fmt.Println("loop")
			return true
		}
	}
	return false
}

func processCodes(codes []string) {
	prev := hh{0, 0, codes[0]}
	var proccesed []hh
	var curr hh
	var err error
	for err == nil {
		curr, err = prev.next(codes)
		if alreadyProcessed(curr, proccesed) {
			fmt.Println(curr)
			return
		}
		proccesed = append(proccesed, curr)
		prev = curr
	}
	fmt.Println("hit err ", err)
}

func main() {
	processCodes(utils.GetStringSlice("samp.txt"))
	processCodes(utils.GetStringSlice("day8.txt"))
}
