package main

import (
	"aoc/util"
	"fmt"
	"os"
	"sort"
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

	nums, err := util.ToIntArray(input)
	if err != nil {
		exit(err)
	}

	sort.Ints(nums)
	oneDiff := 0
	threeDiff := 1 // For diff to computer adapter

	previous := 0
	for _, x := range nums {
		if x-previous == 1 {
			oneDiff++
		} else if x-previous == 3 {
			threeDiff++
		}

		previous = x
	}

	fmt.Printf("Product of diffs: %d\n", oneDiff*threeDiff)
}

func two() {
	input, err := util.ReadLines("input.txt")
	if err != nil {
		exit(err)
	}

	nums, err := util.ToIntArray(input)
	if err != nil {
		exit(err)
	}

	nums = append(nums, 0)
	sort.Ints(nums)
	nums = append(nums, nums[len(nums)-1]+3) // Append internal joltage adapter

	// The options of paths to the current adapter is the sum of possible options to previous adapter that
	// are reachable from this adapter.
	// With the example input of [0, 1, 4, 5, 6, 7], then:
	// 1 can only be reached from 0 -> 1 option (key: 1, value: 1)
	// 4 can only be reached from 1 -> 1 option (key: 4, value 1)
	// 5 can only be reached from 4 -> 1 option (key: 5, value 1)
	// 6 can be reached from both 4 and 5. Which means that the number of options to reach 6 is the sum of options to reach 4, and 5 -> 2 options (key: 6, value 2)
	// 7 can be reached from 4, 5, and 6. Sum of options is memo[4] + memo[5] + memo[6] = 4. (key: 7, value: 4).
	memo := make(map[int]int)
	for i, v := range nums {
		possible := 0
		for _, j := range []int{1, 2, 3} {
			if i-j >= 0 && v-nums[i-j] <= 3 {
				if val, ok := memo[nums[i-j]]; ok {
					possible += val
				}
			}
		}

		if possible == 0 {
			possible = 1
		}
		memo[v] = possible
	}
	fmt.Printf("Found %d paths", memo[nums[len(nums)-1]])
}

func exit(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err)
	os.Exit(1)
}
