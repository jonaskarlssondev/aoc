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
	input, err := util.ReadLines("input.txt")
	if err != nil {
		exit(err)
	}

	rules, chars, messages := ParseInput(input)
	fmt.Printf("Part 1: %d\n", One(rules, chars, messages))
	fmt.Printf("Part 2: %d\n", Two(rules, chars, messages))
}

func One(rules map[int]rule, chars map[int]string, messages []string) int {
	sum := 0

	valid := replace(rules, chars, 0)

	for _, m := range messages {
		if contains(valid, m) {
			sum++
		}
	}
	return sum
}

func Two(rules map[int]rule, chars map[int]string, messages []string) int {
	sum := 0

	// 0: 8 11
	// 8: 42 | 42 8 -> Which means 1..n 42 matches
	// 11: 42 31 | 42 11 31 -> Which means 1..n 42 matches and same number of 31

	// This is now a regex problem..
	// I did google for the regex stuff, because a problem with a regexp solution means you have 2 problems..

	// Compute the base cases
	valid42 := replace(rules, chars, 42)
	valid31 := replace(rules, chars, 31)

	s42 := fmt.Sprintf("(%s)", strings.Join(valid42, "|"))
	s31 := fmt.Sprintf("(%s)", strings.Join(valid31, "|"))

	r8 := fmt.Sprintf("(%s)+", s42) // <-- + allows us to match one or more times

	reg := func(num int) *regexp.Regexp {
		// rule 11 is an equal number of 42 and 31 rules
		return regexp.MustCompile(fmt.Sprintf("^%s%s{%d}%s{%d}$", r8, s42, num, s31, num))
	}

	for _, m := range messages {
		for i := 1; i < 5; i++ { // Keep increasing until result stops increasing
			pattern := reg(i)
			if pattern.MatchString(m) {
				sum++
				break
			}
		}
	}

	return sum
}

type rule [][]int

func ParseInput(s []string) (map[int]rule, map[int]string, []string) {
	messages := []string{}
	rules := map[int]rule{}
	chars := map[int]string{}

	stage := 0
	for _, line := range s {
		if line == "" {
			stage++
			continue
		}

		if stage == 0 {
			l := strings.Split(line, ": ")
			num, _ := strconv.Atoi(l[0])
			rule := rule{}

			if strings.Contains(l[1], "|") {
				for _, part := range strings.Split(l[1], " | ") {
					ids := []int{}
					for _, id := range strings.Split(part, " ") {
						n, _ := strconv.Atoi(id)
						ids = append(ids, n)
					}

					rule = append(rule, ids)
				}
				rules[num] = rule
			} else {
				if strings.Contains(l[1], "\"") {
					char := strings.Replace(l[1], "\"", "", -1)
					chars[num] = char
				} else {
					ids := []int{}
					for _, id := range strings.Split(l[1], " ") {
						n, _ := strconv.Atoi(id)
						ids = append(ids, n)
					}
					rule = append(rule, ids)
					rules[num] = rule
				}

			}
		}

		if stage == 1 {
			messages = append(messages, line)
		}
	}

	return rules, chars, messages
}

func replace(rules map[int]rule, chars map[int]string, i int) []string {
	allRules := []string{}

	if rule, ok := rules[i]; ok {
		for _, nums := range rule {
			single := []string{""}
			for _, num := range nums {
				//fmt.Printf("Checking rule: %v, and num: %d\n", rule, num)
				replacement := replace(rules, chars, num)

				updated := []string{}
				for _, respE := range replacement {
					for _, resE := range single {
						updated = append(updated, resE+respE)
					}
				}
				//fmt.Printf("Got replacement: %v for rule: %v and num: %d and i: %d, which means updated is: %v\n", replacement, rule, num, i, updated)
				single = updated
			}
			allRules = append(allRules, single...)
		}
	} else {
		return []string{chars[i]}
	}

	return allRules
}

func contains(s []string, v string) bool {
	for _, i := range s {
		if i == v {
			return true
		}
	}
	return false
}

func exit(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err)
	os.Exit(1)
}
