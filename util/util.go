package util

import (
	"bufio"
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
