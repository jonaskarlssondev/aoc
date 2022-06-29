package main

import (
	"aoc/util"
	"errors"
	"fmt"
	"os"
)

func main() {
	one()
	two()
}

func one() {
	nums := input()

	x, y, err := twoSum(2020, nums)
	if err != nil {
		exit(err)
	}

	fmt.Printf("%d,%d\n", x, y)
	fmt.Printf("%d\n", x*y)
}

func two() {
	nums := input()

	var sum = 2020

	for i, num := range nums {
		target := sum - num
		excludingNum := append(nums[:i], nums[i+1:]...)
		x, y, err := twoSum(target, excludingNum)
		if err != nil {
			continue
		}

		fmt.Printf("%d,%d,%d\n", x, y, num)
		fmt.Printf("%d\n", x*y*num)
		os.Exit(1)
	}

	fmt.Printf("No solution found.")
}

func twoSum(sum int, a []int) (int, int, error) {
	found := make(map[int]bool)
	for _, num := range a {
		diff := sum - num
		_, ok := found[diff]

		if ok {
			return num, diff, nil
		}

		found[num] = true
	}

	return 0, 0, errors.New("no numbers found")
}

func input() []int {
	input, err := util.ReadLines("input.txt")
	if err != nil {
		exit(err)
	}

	nums, err := util.ToIntArray(input)
	if err != nil {
		exit(err)
	}

	return nums
}

func exit(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err)
	os.Exit(1)
}
