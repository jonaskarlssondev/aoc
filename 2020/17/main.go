package main

import (
	"aoc/util"
	"fmt"
	"os"
)

func main() {
	input, err := util.ReadLines("input.txt")
	if err != nil {
		exit(err)
	}

	data := ParseInput(input)
	fmt.Printf("Part 1: %d\n", One(data, 6))
	fmt.Printf("Part 2: %d\n", Two(data, 6))
}

func One(m map[spot]bool, runs int) int {
	for current := 0; current < runs; current++ {
		m = generate(m, false)
	}

	return len(m)
}

func Two(m map[spot]bool, runs int) int {
	for current := 0; current < runs; current++ {
		m = generate(m, true)
	}

	return len(m)
}

func generate(m map[spot]bool, part2 bool) map[spot]bool {
	delta := []int{-1, 0, 1}

	dw := []int{0}
	if part2 {
		dw = []int{-1, 0, 1}
	}

	spots := make(map[spot]bool, 0)
	for s := range m {
		for _, x := range delta {
			for _, y := range delta {
				for _, z := range delta {
					for _, w := range dw {
						// Cheaper to overwrite than to check
						spots[spot{s.x + x, s.y + y, s.z + z, s.w + w}] = false
					}
				}
			}
		}
	}

	new := make(map[spot]bool, 0)
	for s := range spots {
		n := 0
		for _, x := range delta {
			for _, y := range delta {
				for _, z := range delta {
					for _, w := range dw {
						if x == 0 && y == 0 && z == 0 && w == 0 {
							continue
						}

						loc := spot{s.x + x, s.y + y, s.z + z, s.w + w}
						if val, ok := m[loc]; ok && val {
							n++
						}
					}
				}
			}
		}

		val, ok := m[s]
		if ok && val && (n == 2 || n == 3) {
			new[s] = true
		} else if n == 3 {
			new[s] = true
		}
	}

	return new
}

type spot struct {
	x, y, z, w int
}

func ParseInput(s []string) map[spot]bool {
	m := make(map[spot]bool, 0)

	for y, r := range s {
		for x, v := range r {
			if v == '.' {
				m[spot{x, y, 0, 0}] = false
			} else {
				m[spot{x, y, 0, 0}] = true
			}
		}
	}

	return m
}

func exit(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err)
	os.Exit(1)
}
