package main

import (
	"fmt"
)

func main() {
	one()
	two()
}

func one() {
	input := []int{9, 19, 1, 6, 0, 5, 4}

	m := make(map[int][]int, 0)
	for i, v := range input {
		m[v] = append(m[v], i+1)
	}

	target := 2020
	turn := len(input)
	last := 4

	for turn < target {
		turn++
		last = output(turn, last, m)
	}

	fmt.Printf("Number is: %d\n", last)
}

func two() {
	input := []int{9, 19, 1, 6, 0, 5, 4}

	m := make(map[int][]int, 0)
	for i, v := range input {
		m[v] = append(m[v], i+1)
	}

	target := 30000000
	turn := len(input)
	last := 4

	for turn < target {
		turn++
		last = output(turn, last, m)
	}

	fmt.Printf("Number is: %d\n", last)
}

func output(turn, last int, m map[int][]int) int {
	val, ok := m[last]
	if ok {
		l := len(m[last])
		if l > 1 {
			last = val[l-1] - val[l-2]
		} else {
			last = 0
		}
		m[last] = append(m[last], turn)
	} else {
		last = 0
		m[last] = []int{turn}
	}

	return last
}
