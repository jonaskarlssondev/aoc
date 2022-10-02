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
	input, err := util.ReadLines("input.txt")
	if err != nil {
		exit(err)
	}

	mem := make(map[int]int)
	mask := "X"
	for _, v := range input {
		vs := strings.Split(v, " = ")
		pre, suf := vs[0], vs[1]
		if pre == "mask" {
			mask = reverse(suf)
		} else {
			var addr int
			fmt.Sscanf(pre, "mem[%d]", &addr)

			val, _ := strconv.Atoi(suf)
			rbin := toReverseBinary(val)

			value := applyMask(mask, rbin)

			mem[addr] = value
		}
	}

	sum := 0
	for _, v := range mem {
		sum += v
	}

	fmt.Printf("Sum is: %d\n", sum)
}

func two() {
	input, err := util.ReadLines("input.txt")
	if err != nil {
		exit(err)
	}

	mem := make(map[int]int)
	mask := "X"
	for _, v := range input {
		vs := strings.Split(v, " = ")
		pre, suf := vs[0], vs[1]
		if pre == "mask" {
			mask = reverse(suf)
		} else {
			var addr int
			fmt.Sscanf(pre, "mem[%d]", &addr)
			rbAddr := toReverseBinary(addr)

			maskedAddr := applyAddressMask(mask, rbAddr)
			addresses := expandAddr(maskedAddr)

			val, _ := strconv.Atoi(suf)

			for _, a := range addresses {
				mem[a] = val
			}
		}
	}

	sum := 0
	for _, v := range mem {
		sum += v
	}

	fmt.Printf("Sum is: %d\n", sum)

}

func expandAddr(addr string) (addresses []int) {
	temps := []string{addr}

	for _, v := range addr {
		char := string(v)
		if char == "X" {
			temp := make([]string, 0)
			for _, a := range temps {
				zero := strings.Replace(a, "X", "0", 1)
				temp = append(temp, zero)

				one := strings.Replace(a, "X", "1", 1)
				temp = append(temp, one)
			}

			temps = temp
		}
	}

	for _, v := range temps {
		addr := toValue(v)
		addresses = append(addresses, addr)
	}

	return
}

func applyAddressMask(mask, binary string) string {
	value := ""
	for i, runeMask := range mask {
		m := string(runeMask)

		b := "0"
		if i < len(binary) {
			b = string(binary[i])
		}

		if m == "X" {
			value += "X"
		} else {
			if m == "1" || b == "1" {
				value += "1"
			} else {
				value += "0"
			}
		}
	}

	return value
}

func toValue(binary string) int {
	value := 0
	for i, r := range binary {
		v := string(r)

		if v == "1" {
			value += 1 << i
		}
	}

	return value
}

func applyMask(mask, binary string) int {
	value := 0
	for i, runeMask := range mask {
		m := string(runeMask)

		b := "0"
		if i < len(binary) {
			b = string(binary[i])
		}

		if m == "1" {
			value += 1 << i
		} else if m == "X" {
			if b == "1" {
				value += 1 << i
			}
		}
	}

	return value
}

func toReverseBinary(v int) string {
	bin := strconv.FormatInt(int64(v), 2)
	return reverse(bin)
}

func reverse(s string) string {
	var res string
	for _, v := range s {
		res = string(v) + res
	}
	return res
}

func exit(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err)
	os.Exit(1)
}
