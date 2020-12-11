package main

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	samp = `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`

	samp2 = `


`

	bnameReg   = regexp.MustCompile(`[0-9]\s[a-z]+\s[a-z]+`)
	twoWordReg = regexp.MustCompile(`[a-z]+\s[a-z]+`)
)

func inputToRules(txt string) map[string][]string {
	rules := make(map[string][]string)
	lines := strings.Split(txt, "\n")
	for _, line := range lines {
		c := strings.Split(line, "bag")
		if len(c) < 3 {
			fmt.Printf("aah: %s %s\n", line, c)
			// return rules
			continue
		}
		key := strings.TrimSpace(c[0])
		vals := bnameReg.FindAllString(line, -1)
		rules[key] = vals
	}

	return rules
}

func refine(rules map[string][]string) {
	for base, inner := range rules {
		// fmt.Println("Base, ", base, "\tinner, ", inner, "\tlen inner, ", len(inner))
		for i, in := range inner {
			if len(in) < 2 {
				fmt.Println("ahhhhhhhhhhhhh, Base, ", base, "\tinner, ", inner, "\tlen inner, ", len(inner))
			}
			inner[i] = strings.TrimSpace(in[2:])
		}
	}
}

type tree struct {
	name     string
	num      int
	children []tree
}

// func resolveList(target string, colors map[string][]string) {
// 	// for _, val := range colors {
// 	// 	for _, c := range val {

// 	// 	}
// 	//}
// 	accum := ""
// 	for _, c := range colors[target] {
// 		for
// 	}
// }
func find(target string, rules map[string][]string, foundkeys map[string]int) {
	// if strings.Contains(foundkeys, target) {
	// 	return 0
	// }
	// r := 0
	for key, vals := range rules {
		// fmt.Println("hmm, ", strings.Contains(foundkeys, key))
		for _, v := range vals {
			if target == v {
				// r++
				foundkeys[key]++
				// fmt.Println("found target: ", target, ", at key: ", key, " vals: ", vals, " v: ", v, " fkeys: ", foundkeys)

				find(key, rules, foundkeys)
			}
		}
	}
	// return foundkeys
}

// func addToTree(rules map[string][]string) tree {
// 	result := tree{name: "shiny gold", num: 0}
// 	target := "shiny gold"
// 	childRules := rules[target];
// 	for len(childRules) != 0 {
// 		fmt.Println(childRules)
// 		for _, rchild := range childRules {

// 		}
// 		// if len(rchild) < 3 {
// 		// 	fmt.Println("ahhhh")
// 		// 	return tree{}
// 		// }
// 		// child := strings.TrimSpace(rchild[2:])

// 	}
// 	return result
// }

func totalTarget(target string, rules map[string][]string) {
	for base, inner := range rules {
		fmt.Println("Base, ", base, "\tinner, ", inner, "\tlen inner, ", len(inner))

	}

}

func main() {
	// x := inputToRules(samp)
	// // fmt.Println(x)
	// refine(x)
	// // fmt.Println(x)
	// y := make(map[string]int)
	// find("shiny gold", x, y)

	// fmt.Println(len(y))

	// input, _ := ioutil.ReadFile("day7.txt")
	// // x := inputToRules(string(input))
	// z := inputToRules(string(input))
	// addToTree(z)

	// fmt.Println(z)
	// refine(x)
	// y := make(map[string]int)
	// fmt.Println(y)
	// find("shiny gold", x, y)
	// fmt.Println((y))

}
