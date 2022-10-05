package main

import (
	"aoc/util"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	one()
	two()
}

func one() {
	in, err := util.ReadLines("input.txt")
	if err != nil {
		exit(err)
	}

	input := formatInput(in)
	sum := 0
	for _, ticket := range input.Others {
		for _, num := range ticket.data {
			pass := false
			for _, check := range input.Validity {
				if check(num) {
					pass = true
					break
				}
			}

			if !pass {
				sum += num
			}
		}
	}

	fmt.Printf("Sum is: %d\n", sum)
}

func two() {
	in, err := util.ReadLines("input.txt")
	if err != nil {
		exit(err)
	}

	input := formatInput(in)

	valid := make([]ticket, 0)
	for _, ticket := range input.Others {
		isValid := true
		for _, num := range ticket.data {
			pass := false
			for _, check := range input.Validity {
				if check(num) {
					pass = true
					break
				}
			}

			if !pass {
				isValid = false
				break
			}
		}

		if isValid {
			valid = append(valid, ticket)
		}
	}

	skip := make(map[int]bool)
	correct := make(map[int]int)

	for len(skip) < 20 {
		possible := make([][]int, 0)

		for i, check := range input.Validity {
			possible = append(possible, make([]int, 0))

			for j := 0; j < 20; j++ {
				val, ok := skip[j]
				if ok && val {
					continue
				}
				pass := true
				for _, ticket := range valid {
					if !check(ticket.data[j]) {
						pass = false
						break
					}
				}

				if pass {
					possible[i] = append(possible[i], j)
				}
			}

			if len(possible[i]) == 1 {
				col := possible[i][0]
				skip[col] = true
				correct[i] = col
			}
		}
	}

	product := 1
	for i := 0; i < 6; i++ {
		col := correct[i]
		product *= input.Own.data[col]
	}

	fmt.Printf("Product is: %d\n", product)
}

type input struct {
	Validity []func(int) bool
	Own      ticket
	Others   []ticket
}

type ticket struct {
	data []int
}

func formatInput(in []string) input {
	data := input{
		Validity: make([]func(int) bool, 0),
		Others:   make([]ticket, 0),
	}

	phase := 0
	for i := 0; i < len(in); i++ {
		v := in[i]
		if v == "" {
			i++
			phase++
			continue
		}

		if phase == 0 {
			var i1, i2, i3, i4 int
			val := strings.Split(v, ": ")
			fmt.Sscanf(val[1], "%d-%d or %d-%d", &i1, &i2, &i3, &i4)
			data.Validity = append(data.Validity, func(v int) bool { return (v >= i1 && v <= i2) || (v >= i3 && v <= i4) })
		}

		if phase == 1 {
			data.Own = ticket{
				data: make([]int, 0),
			}

			nums := strings.Split(v, ",")
			for _, c := range nums {
				val, _ := strconv.Atoi(c)
				data.Own.data = append(data.Own.data, val)
			}
		}

		if phase == 2 {
			t := ticket{
				data: make([]int, 0),
			}

			nums := strings.Split(v, ",")
			for _, c := range nums {
				val, _ := strconv.Atoi(c)
				t.data = append(t.data, val)
			}

			data.Others = append(data.Others, t)
		}
	}

	return data
}

func exit(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err)
	os.Exit(1)
}
