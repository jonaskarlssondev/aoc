package main

import (
	"aoc/util"
	"fmt"
	"os"
	"strconv"
)

type queue struct {
	data []int
}

func (q *queue) enqueue(v int) {
	q.data = append(q.data, v)
}

func (q *queue) dequeue() int {
	val := q.data[0]
	q.data = q.data[1:]
	return val
}

func main() {
	input, err := util.ReadLines("input.txt")
	if err != nil {
		exit(err)
	}

	d1, d2 := parseInput(input)
	fmt.Printf("Part 1: %d\n", One(d1, d2))
	fmt.Printf("Part 2: %d\n", Two(d1, d2))
}

func One(d1, d2 queue) int {
	for len(d1.data) > 0 && len(d2.data) > 0 {
		c1 := d1.dequeue()
		c2 := d2.dequeue()

		if c1 > c2 {
			d1.enqueue(c1)
			d1.enqueue(c2)
		} else {
			d2.enqueue(c2)
			d2.enqueue(c1)
		}
	}

	var d queue
	if len(d1.data) > 0 {
		d = d1
	} else {
		d = d2
	}

	sum := 0
	for x, i := 1, len(d.data)-1; i >= 0; x, i = x+1, i-1 {
		sum += d.data[i] * x
	}

	return sum
}

func Two(d1, d2 queue) int {
	res, _ := recursive(d1, d2, true)
	return res
}

func recursive(d1, d2 queue, game1 bool) (int, bool) {
	prev1 := map[string]bool{}
	prev2 := map[string]bool{}

	var p1win bool
	for len(d1.data) > 0 && len(d2.data) > 0 {
		rep1 := fmt.Sprintf("%v", d1.data)
		rep2 := fmt.Sprintf("%v", d2.data)

		c1 := d1.dequeue()
		c2 := d2.dequeue()

		if prev1[rep1] || prev2[rep2] {
			p1win = true
		} else {
			prev1[rep1] = true
			prev2[rep2] = true

			if len(d1.data) >= c1 && len(d2.data) >= c2 {
				d1d := make([]int, c1)
				copy(d1d, d1.data[:c1])
				d1r := queue{d1d}

				d2d := make([]int, c2)
				copy(d2d, d2.data[:c2])
				d2r := queue{d2d}

				_, p1win = recursive(d1r, d2r, false)
			} else {
				p1win = c1 > c2
			}
		}

		if p1win {
			d1.enqueue(c1)
			d1.enqueue(c2)
		} else {
			d2.enqueue(c2)
			d2.enqueue(c1)
		}
	}

	if !game1 {
		return -1, p1win
	}

	var d queue
	if len(d1.data) > 0 {
		d = d1
	} else {
		d = d2
	}

	sum := 0
	for x, i := 1, len(d.data)-1; i >= 0; x, i = x+1, i-1 {
		sum += d.data[i] * x
	}

	return sum, false
}

func parseInput(s []string) (queue, queue) {
	var sep int

	for i, r := range s {
		if r == "" {
			sep = i
		}
	}

	d1 := parse(s[:sep])
	d2 := parse(s[sep+1:])

	return d1, d2
}

func parse(s []string) queue {
	d := queue{}

	for _, r := range s[1:] {
		val, _ := strconv.Atoi(r)
		d.enqueue(val)
	}

	return d
}

func exit(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err)
	os.Exit(1)
}
