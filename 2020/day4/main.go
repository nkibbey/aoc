package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	//byr (Birth Year)
	Byr string `json:"byr"`
	//iyr (Issue Year)
	Iyr string `json:"iyr"`
	//eyr (Expiration Year)
	Eyr string `json:"eyr"`
	//hgt (Height)
	Hgt string `json:"hgt"`
	//(Hair Color)
	Hcl string `json:"hcl"`
	//(Eye Color)
	Ecl string `json:"ecl"`
	//pid (Passport ID)
	Pid string `json:"pid"`
	//cid (Country ID)
	Cid string `json:"cid"`
}

func decodeInput(fn string) []Passport {
	var p []Passport
	f, _ := ioutil.ReadFile(fn)
	json.Unmarshal(f, &p)
	return p
}

func lenCheck(s string, el int) bool {
	if s == "" {
		return false
	} else if len(s) != el {
		fmt.Printf("Suspicious len %s expected len %v\n", s, el)
		return false
	}
	return true
}

func between(x, min, max int) bool {
	return x >= min && x <= max
}

func validYear(ys string, start, end int) bool {
	if !lenCheck(ys, 4) {
		return false
	}

	y, err := strconv.Atoi(ys)
	if err != nil {
		fmt.Printf("conv err: %v\n", err)
		return false
	}
	return between(y, start, end)
}

func validByr(byr string) bool {
	return validYear(byr, 1920, 2002)
}

func validIyr(iyr string) bool {
	return validYear(iyr, 2010, 2020)
}

func validEyr(eyr string) bool {
	return validYear(eyr, 2020, 2030)
}

func validHgt(hgt string) bool {
	if hgt == "" {
		return false
	} else if len(hgt) <= 2 {
		return false
	}
	// assign height requirements
	hs := strings.TrimSuffix(hgt, "in")
	max, min := -1, -1
	if len(hs) < len(hgt) {
		min, max = 59, 76
	} else {
		hs = strings.TrimSuffix(hgt, "cm")
		if len(hs) < len(hgt) {
			min, max = 150, 193
		}
	}

	// conv to height
	h, err := strconv.Atoi(hs)
	if err != nil {
		fmt.Printf("conv err: %v\n", err)
		return false
	}

	return between(h, min, max)
}

func validHcl(hcl string) bool {
	if !lenCheck(hcl, 7) {
		return false
	}
	r, _ := regexp.Match(`#[a-f0-9]{6}`, []byte(hcl))
	return r
}

func validEcl(ecl string) bool {
	if !lenCheck(ecl, 3) {
		return false
	}
	colors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, c := range colors {
		if c == ecl {
			return true
		}
	}
	return false
}

func validPid(pid string) bool {
	r, _ := regexp.Match(`[0-9]{9}`, []byte(pid))
	return r
}

func validPP(pp Passport) bool {
	return validByr(pp.Byr) && validIyr(pp.Iyr) && validEyr(pp.Eyr) && validHgt(pp.Hgt) && validHcl(pp.Hcl) && validEcl(pp.Ecl) && validPid(pp.Pid)
}

func numValidPPs(pps []Passport) int {
	vpps := 0
	for _, pp := range pps {
		if validPP(pp) {
			vpps++
		}
	}
	return vpps
}

func main() {
	pps := decodeInput("day4.json")
	// pps := decodeInput("samp.json")
	// fmt.Println(pps)
	fmt.Println(numValidPPs(pps))
}
