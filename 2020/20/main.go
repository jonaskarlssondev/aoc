package main

import (
	"aoc/util"
	"fmt"
	"math"
	"os"
)

func main() {
	input, err := util.ReadLines("input.txt")
	if err != nil {
		exit(err)
	}

	data := ParseInput(input)
	fmt.Printf("Part 1: %d\n", One(data))
	fmt.Printf("Part 2: %d\n", Two(data))
}

type image struct {
	id                       int
	data                     [][]rune
	top, right, bottom, left string
}

func One(data []image) int64 {
	var product uint64 = 1

	size := int(math.Sqrt(float64(len(data))))
	puzzle := make([][]*image, size)
	for i := 0; i < size; i++ {
		puzzle[i] = make([]*image, size)
	}

	picture := layPuzzle(puzzle, data, size, map[int]bool{})

	product *= uint64(picture[0][0].id)
	product *= uint64(picture[size-1][0].id)
	product *= uint64(picture[0][size-1].id)
	product *= uint64(picture[size-1][size-1].id)

	return int64(product)
}

func layPuzzle(puzzle [][]*image, images []image, size int, used map[int]bool) [][]*image {
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if puzzle[i][j] == nil {
				for idx, img := range images {
					if !used[idx] {
						for _, alt := range getAllSides(img) {
							if i != 0 {
								if puzzle[i-1][j].bottom != alt.top {
									continue
								}
							}

							if j != 0 {
								if puzzle[i][j-1].right != alt.left {
									continue
								}
							}

							puzzle[i][j] = &alt
							used[idx] = true

							res := layPuzzle(puzzle, images, size, used)
							if res != nil {
								return res
							}

							puzzle[i][j] = nil
							used[idx] = false
						}
					}
				}
				// No image matched
				if puzzle[i][j] == nil {
					return nil
				}
			}
		}
	}

	return puzzle
}

func getAllSides(img image) []image {
	imgs := make([]image, 0)

	imgs = append(imgs, img)

	//left-right flip
	imgs = append(imgs, image{
		id:     img.id,
		top:    reverse(img.top),
		right:  img.left,
		bottom: reverse(img.bottom),
		left:   img.right,
		data:   reverseMatrix(copyMatrix(img.data)),
	})

	//left-right flip -> 90 degree rotation
	imgs = append(imgs, image{
		id:     img.id,
		top:    reverse(img.right),
		right:  reverse(img.top),
		bottom: reverse(img.left),
		left:   reverse(img.bottom),
		data:   rotate90(reverseMatrix(copyMatrix(img.data))),
	})

	//Up-down flip
	imgs = append(imgs, image{
		id:     img.id,
		top:    img.bottom,
		right:  reverse(img.right),
		bottom: img.top,
		left:   reverse(img.left),
		data:   flip(copyMatrix(img.data)),
	})

	// Up-down flip -> 90 degree rotation
	imgs = append(imgs, image{
		id:     img.id,
		top:    img.left,
		right:  img.bottom,
		bottom: img.right,
		left:   img.top,
		data:   rotate90(flip(copyMatrix(img.data))),
	})

	// 90 degree rotation
	imgs = append(imgs, image{
		id:     img.id,
		top:    reverse(img.left),
		right:  img.top,
		bottom: reverse(img.right),
		left:   reverse(img.bottom),
		data:   rotate90(copyMatrix(img.data)),
	})

	// 180 degree rotation
	imgs = append(imgs, image{
		id:     img.id,
		top:    reverse(img.bottom),
		right:  reverse(img.left),
		bottom: reverse(img.top),
		left:   reverse(img.right),
		data:   rotate90(rotate90(copyMatrix(img.data))),
	})

	// 270 degree rotation
	imgs = append(imgs, image{
		id:     img.id,
		top:    img.right,
		right:  reverse(img.bottom),
		bottom: img.left,
		left:   reverse(img.top),
		data:   rotate90(rotate90(rotate90(copyMatrix(img.data)))),
	})

	return imgs
}

func reverse(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, rune := range s {
		n--
		runes[n] = rune
	}
	return string(runes[n:])
}

func Two(data []image) int {
	size := int(math.Sqrt(float64(len(data))))
	puzzle := make([][]*image, size)
	for i := 0; i < size; i++ {
		puzzle[i] = make([]*image, size)
	}

	picture := layPuzzle(puzzle, data, size, map[int]bool{})

	joinedPic := make([][]rune, size*8)
	for i := range joinedPic {
		joinedPic[i] = make([]rune, size*8)
	}

	// Copy each picture into the joined pic "chunk" by chunk.
	for i := 0; i < len(picture); i++ {
		for j := 0; j < len(picture); j++ {
			for k := 0; k < len(picture[0][0].data); k++ {
				// slice low:high is [low:high)
				copy(joinedPic[i*8+k][j*8:j*8+8], picture[i][j].data[k])
			}
		}
	}

	// search for pattern
	// Move 1 by 1 and look for 20x3 pattern
	total := 0
	for i := 0; i < len(joinedPic); i++ {
		for j := 0; j < len(joinedPic[0]); j++ {
			if joinedPic[i][j] == '#' {
				total++
			}
		}
	}

	img := image{
		id:     0,
		top:    "..",
		right:  "..",
		bottom: "..",
		left:   "..",
		data:   joinedPic,
	}

	monsters := 0
	for _, alt := range getAllSides(img) {
		for i := 0; i < len(joinedPic)-3; i++ {
			r1 := alt.data[i]
			r2 := alt.data[i+1]
			r3 := alt.data[i+2]

			for j := 0; j < len(joinedPic[i])-20; j++ {
				check := [][]rune{
					r1[j : j+20],
					r2[j : j+20],
					r3[j : j+20],
				}

				if findPattern(check) {
					monsters++
				}
			}
		}
	}

	fmt.Printf("Monsters: %d\n", monsters)

	return total - monsters*15
}

func findPattern(m [][]rune) bool {
	return m[0][18] == '#' &&
		m[1][0] == '#' && m[1][5] == '#' && m[1][6] == '#' && m[1][11] == '#' && m[1][12] == '#' && m[1][17] == '#' && m[1][18] == '#' && m[1][19] == '#' &&
		m[2][1] == '#' && m[2][4] == '#' && m[2][7] == '#' && m[2][10] == '#' && m[2][13] == '#' && m[2][16] == '#'
}

func ParseInput(s []string) []image {
	imgs := []image{}

	for i := 0; i < len(s); i += 12 {
		img := image{}

		fmt.Sscanf(s[i], "Tile %d:", &img.id)

		img.data = make([][]rune, 0)

		img.top = s[i+1]
		for j := i + 1; j < i+11; j++ {
			if j != i+1 && j != i+10 {
				rowXborder := s[j][1 : len(s[j])-1]
				img.data = append(img.data, []rune(rowXborder))
			}

			img.left += string(s[j][0])
			img.right += string(s[j][9])
		}
		img.bottom = s[i+10]

		imgs = append(imgs, img)
	}

	return imgs
}

func rotate90[T any](m [][]T) [][]T {
	m = transpose(m)
	m = reverseMatrix(m)

	return m
}

func transpose[T any](m [][]T) [][]T {
	new := make([][]T, len(m))
	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[0]); j++ {
			new[j] = append(new[j], m[i][j])
		}
	}
	return new
}

func flip[T any](m [][]T) [][]T {
	for i, j := 0, len(m)-1; i < j; i, j = i+1, j-1 {
		m[i], m[j] = m[j], m[i]
	}

	return m
}

func reverseMatrix[T any](m [][]T) [][]T {
	for i := 0; i < len(m); i++ {
		for j, k := 0, len(m)-1; j < k; j, k = j+1, k-1 {
			m[i][j], m[i][k] = m[i][k], m[i][j]
		}
	}

	return m
}

func copyMatrix[T any](m [][]T) [][]T {
	c := make([][]T, len(m))
	for i := range m {
		c[i] = make([]T, len(m[i]))
		copy(c[i], m[i])
	}

	return c
}

func exit(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err)
	os.Exit(1)
}
