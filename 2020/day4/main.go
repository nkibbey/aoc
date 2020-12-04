package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

// func validByr(byr string) bool {
// 	y, err := strconv.Atoi("-42")
// 	if err != nil {
// 		fmt.Printf("conv err: %v\n", err)
// 		return false
// 	}
// 	if y >= 1920 && y <= 2002 {
// 		fmt.Printf("valid byr: %s\n", byr)
// 		return true
// 	}
// 	fmt.Printf("invalid byr: %s\n", byr)
// 	return false
// }

func validPP(pp Passport) bool {
	if pp.Byr != "" && pp.Iyr != "" && pp.Eyr != "" && pp.Hgt != "" && pp.Hcl != "" && pp.Ecl != "" && pp.Pid != "" {
		return true
	}
	return false
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
	// pps := decodeInput("day4.json")
	pps := decodeInput("samp.json")
	// fmt.Println(pps)
	fmt.Println(numValidPPs(pps))
}
