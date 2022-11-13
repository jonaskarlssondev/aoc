package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := "598162734"
	fmt.Printf("Part 1: %s\n", One(data, 100))
	fmt.Printf("Part 2: %d\n", Two(data, 10000000))
}

func One(data string, moves int) string {
	cIndex := 0
	highest := 9

	for i := 0; i < moves; i++ {
		current := string(data[cIndex])

		pick1 := (cIndex + 1) % len(data)
		pick2 := (cIndex + 2) % len(data)
		pick3 := (cIndex + 3) % len(data)
		c1, c2, c3 := data[pick1], data[pick2], data[pick3]
		data = strings.Replace(data, string(c1), "", 1)
		data = strings.Replace(data, string(c2), "", 1)
		data = strings.Replace(data, string(c3), "", 1)

		find, _ := strconv.Atoi(current)
		newLocation := -1
		for newLocation == -1 {
			find--
			if find < 1 {
				find = highest
			}
			newLocation = strings.Index(data, fmt.Sprint(find))
		}

		if newLocation == len(data)-1 {
			data += string(c1) + string(c2) + string(c3)
		} else {
			pre, post := data[:newLocation+1], data[newLocation+1:]
			data = pre + string(c1) + string(c2) + string(c3) + post
		}

		cIndex = (strings.Index(data, fmt.Sprint(current)) + 1) % len(data)
	}

	res := ""
	start := strings.Index(data, "1")
	i := (start + 1) % len(data)
	for i != start {
		res += string(data[i])
		i = (i + 1) % len(data)
	}

	return res
}

type node struct {
	val        int
	prev, next *node
}

func Two(data string, moves int) int {
	val, _ := strconv.Atoi(string(data[0]))
	current := &node{val: val}

	// Cache
	cache := map[int]*node{}
	cache[val] = current

	cup := current
	for _, r := range data[1:] {
		val, _ := strconv.Atoi(string(r))
		cup.next = &node{
			val:  val,
			prev: cup,
		}
		cup = cup.next
		cache[val] = cup
	}

	for i := 10; i <= 1000000; i++ {
		cup.next = &node{
			val:  i,
			prev: cup,
		}
		cup = cup.next
		cache[i] = cup
	}

	// Close double linked list
	cup.next = current
	current.prev = cup

	for i := 0; i < moves; i++ {
		find := current.val - 1

		first := current.next
		last := current.next.next.next

		for find == first.val || find == first.next.val || find == first.next.next.val || find <= 0 {
			find--
			if find <= 0 {
				find = 1000000
			}
		}

		// Remove first -> last from LL
		current.next = last.next
		current.next.prev = current

		found := cache[find]
		last.next = found.next
		found.next.prev = last
		found.next = first
		first.prev = found

		current = current.next
	}

	return cache[1].next.val * cache[1].next.next.val
}
