package main

import (
	"aoc/util"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, err := util.ReadLines("input.txt")
	if err != nil {
		exit(err)
	}

	data := ParseInput(input, false)
	fmt.Printf("Part 1: %d\n", One(data))
	data = ParseInput(input, true)
	fmt.Printf("Part 2: %d\n", Two(data))
}

func One(ops []operation) int {
	sum := 0
	for _, op := range ops {
		sum += op.compute()
	}
	return sum
}

func Two(ops []operation) int {
	sum := 0
	for _, op := range ops {
		sum += op.compute()
	}
	return sum
}

type operation struct {
	rpn string
}

func (op operation) compute() int {
	stack := stack[int]{}

	for _, c := range op.rpn {
		char := string(c)
		val, err := strconv.Atoi(char)
		if err != nil {
			var v1, v2 int
			stack, v1 = stack.pop()
			stack, v2 = stack.pop()

			if char == "*" {
				stack = stack.push(v1 * v2)
			} else {
				stack = stack.push(v1 + v2)
			}
		} else {
			stack = stack.push(val)
		}
	}

	return stack.peek()
}

// PareInput returns an array of operations where each operation is an rpn created using the shunted yard algorithm
func ParseInput(s []string, part2 bool) []operation {
	ops := []operation{}

	for _, line := range s {
		op := parseOp(line, part2)
		ops = append(ops, op)
	}

	return ops
}

func parseOp(s string, part2 bool) operation {
	op := operation{""}

	precedence := map[string]int{
		"+": 2,
		"*": 1,
	}

	stack := stack[string]{}
	for _, c := range s {
		char := string(c)
		if char == " " {
			continue
		}
		_, err := strconv.Atoi(char)

		if err != nil {
			// Its an operator

			// If a parenthesis match has been found
			// Pop and push to stack until first match
			// And discard parenthesis
			if char == ")" {
				for stack.peek() != "(" {
					var v string
					stack, v = stack.pop()
					op.rpn += v
				}
				stack, _ = stack.pop()
			} else {
				if len(stack) > 0 && stack.peek() != "(" && char != "(" {
					if part2 {
						for len(stack) > 0 && precedence[stack.peek()] >= precedence[char] {
							var v string
							stack, v = stack.pop()
							op.rpn += v
						}
					} else {
						var v string
						stack, v = stack.pop()
						op.rpn += v
					}
				}

				stack = stack.push(char)
			}
		} else {
			// Its a value
			op.rpn += char
		}
	}

	for len(stack) > 0 {
		var v string
		stack, v = stack.pop()
		op.rpn += v
	}
	return op
}

func exit(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err)
	os.Exit(1)
}

type stack[T any] []T

func (s stack[T]) push(v T) stack[T] {
	return append(s, v)
}

func (s stack[T]) pop() (stack[T], T) {
	l := len(s)
	return s[:l-1], s[l-1]
}

func (s stack[T]) peek() T {
	return s[len(s)-1]
}
