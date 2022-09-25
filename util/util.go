package util

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

// ReadLines reads a file from the path and returns a slice containing each row read from the file.
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// Split separates a string into a slice of strings and returns the first and second entry.
func Split(s string, sep string) (string, string) {
	split := strings.Split(s, sep)
	return split[0], split[1]
}

// ToIntArray takes a string slice and attemps to convert each string to an int.
func ToIntArray(s []string) ([]int, error) {
	arr := make([]int, len(s))
	for i, str := range s {
		j, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}

		arr[i] = j
	}

	return arr, nil
}

// TwoSum tries to find two numbers in a, that add up to sum.
func TwoSum(sum int, a []int) (int, int, error) {
	found := make(map[int]bool)
	for _, num := range a {
		diff := sum - num

		if found[diff] {
			return num, diff, nil
		}

		found[num] = true
	}

	return 0, 0, errors.New("no numbers found")
}

func Min(a []int) (int, error) {
	if len(a) == 0 {
		return 0, errors.New("length 0 array")
	}

	v := a[0]

	for _, x := range a {
		if x < v {
			v = x
		}
	}

	return v, nil
}

func Max(a []int) (int, error) {
	if len(a) == 0 {
		return 0, errors.New("length 0 array")
	}

	v := a[0]

	for _, x := range a {
		if x > v {
			v = x
		}
	}

	return v, nil
}
