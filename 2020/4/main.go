package main

import (
	"aoc/util"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	one()
	two()
}

func one() {
	input, err := util.ReadLines("input.txt")
	if err != nil {
		exit(err)
	}

	valid := 0
	total := 0
	passport := make(map[string]string)
	for i, line := range input {
		if line != "" {
			entries := strings.Split(line, " ")
			if len(entries) >= 1 {
				for _, entry := range entries {
					key, value := split(entry, ":")
					passport[key] = value
				}
			}
		}

		if line == "" || i == len(input)-1 {
			_, byrOK := passport["byr"]
			_, iyrOK := passport["iyr"]
			_, eyrOK := passport["eyr"]
			_, hgtOK := passport["hgt"]
			_, hclOK := passport["hcl"]
			_, eclOK := passport["ecl"]
			_, pidOK := passport["pid"]

			if byrOK && iyrOK && eyrOK && hgtOK && hclOK && eclOK && pidOK {
				valid++
			}

			total++
			passport = make(map[string]string)
		}
	}

	fmt.Println("Valid: ", valid)
	fmt.Println("Total: ", total)
}

func two() {
	input, err := util.ReadLines("input.txt")
	if err != nil {
		exit(err)
	}

	valid := 0
	total := 0
	passport := make(map[string]string)
	for i, line := range input {
		if line != "" {
			entries := strings.Split(line, " ")
			if len(entries) >= 1 {
				for _, entry := range entries {
					key, value := split(entry, ":")
					passport[key] = value
				}
			}
		}

		if line == "" || i == len(input)-1 {
			byrStr, byrOK := passport["byr"]
			if byrOK {
				byr, err := strconv.Atoi(byrStr)
				if err != nil {
					byrOK = false
				}
				byrOK = byr >= 1920 && byr <= 2002
			}

			iyrStr, iyrOK := passport["iyr"]
			if iyrOK {
				iyr, err := strconv.Atoi(iyrStr)
				if err != nil {
					iyrOK = false
				}
				iyrOK = iyr >= 2010 && iyr <= 2020
			}

			eyrStr, eyrOK := passport["eyr"]
			if eyrOK {
				eyr, err := strconv.Atoi(eyrStr)
				if err != nil {
					eyrOK = false
				}
				eyrOK = eyr >= 2020 && eyr <= 2030
			}

			hgt, hgtOK := passport["hgt"]
			if hgtOK {
				hgtNum := hgt[:len(hgt)-2]
				hgtNumInt, err := strconv.Atoi(hgtNum)
				if err != nil {
					hgtOK = false
				}

				switch hgt[len(hgt)-2:] {
				case "cm":
					hgtOK = hgtNumInt >= 150 && hgtNumInt <= 193
					break
				case "in":
					hgtOK = hgtNumInt >= 59 && hgtNumInt <= 76
					break
				default:
					hgtOK = false
					break
				}
			}

			hclStr, hclOK := passport["hcl"]
			if hclOK {
				hclOK = hclStr[0] == '#' && len(hclStr) == 7
				regval, err := regexp.MatchString("[0-9a-f]", hclStr[1:])
				if err != nil || !regval {
					regval = false
				}
				hclOK = hclOK && regval
			}

			eclStr, eclOK := passport["ecl"]
			if eclOK {
				eclOK = eclStr == "amb" || eclStr == "blu" || eclStr == "brn" || eclStr == "gry" || eclStr == "grn" || eclStr == "hzl" || eclStr == "oth"
			}

			pidStr, pidOK := passport["pid"]
			if pidOK {
				_, err := strconv.Atoi(pidStr)
				if err != nil {
					pidOK = false
				}

				pidOK = pidOK && len(pidStr) == 9
			}

			if byrOK && iyrOK && eyrOK && hgtOK && hclOK && eclOK && pidOK {
				valid++
			}

			total++
			passport = make(map[string]string)
		}
	}

	fmt.Println("Valid: ", valid)
	fmt.Println("Total: ", total)
}

func split(s string, sep string) (string, string) {
	split := strings.Split(s, sep)
	return split[0], split[1]
}

func exit(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err)
	os.Exit(1)
}
