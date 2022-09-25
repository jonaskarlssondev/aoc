package main

import (
	"aoc/util"
	"fmt"
	"os"
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

	for i := 25; i < len(nums)-25; i++ {
		sum := nums[i]
		_, _, err := util.TwoSum(sum, nums[i-25:i])
		if err != nil {
			fmt.Printf("Could not find a sum for: %d\n", sum)
			break
		}
	}
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

	find := 57195069

	for i := 0; i < len(nums); i++ {
		sum := 0
		j := i
		for sum < find {
			j++
			sum += nums[j]
		}

		if sum == find {
			fmt.Printf("Found range %v\n", nums[i:j])

			min, _ := util.Min(nums[i:j])
			max, _ := util.Max(nums[i:j])
			fmt.Printf("Found min %d\n", min)
			fmt.Printf("Found max %d\n", max)

			fmt.Printf("Sum %d in sliding window.\n", min+max)
			break
		}
	}
}

func exit(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err)
	os.Exit(1)
}
