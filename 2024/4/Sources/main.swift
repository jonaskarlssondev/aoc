import Foundation

let fileURL = URL(fileURLWithPath: "input.txt")
let contents = try String(contentsOf: fileURL)
let data = contents.components(separatedBy: "\n")

var arr = [[Character]]()

for row in data {
	arr.append(Array(row))
}

var targetWord = Array("MAS")

func search(m: [[Character]], x: Int, y: Int) -> Int {
	let directions = [-1,0,1]

	var s = 0
	for i in directions {
		for j in directions {
			if i != 0 || j != 0 {
				s += search2(m: m, x: x + i, dx: i, y: y + j, dy: j, target: 0)
			}
		}
	}

	return s
}

func search2(m: [[Character]], x: Int, dx: Int, y: Int, dy: Int, target: Int) -> Int {
	if x < 0 || y < 0 || x >= m.count || y >= m[x].count {
		return 0
	}

	if m[x][y] != targetWord[target] {
		return 0
	}

	if target+1 == targetWord.count {
		return 1
	}

	return search2(m:m, x:x+dx, dx: dx, y: y+dy, dy: dy, target: target+1)
}


var sum = 0

for i in arr.indices {
	for j in arr[i].indices {
		if arr[i][j] == "X" {
			sum += search(m:arr, x:i, y:j)
		}
	}
}

print(sum)

var sum2 = 0
for i in arr.indices {
	for j in arr[i].indices {
		if arr[i][j] == "A" {
			sum2 += search3(m:arr, x:i, y:j)
		}
	}
}

print(sum2)

func search3(m: [[Character]], x: Int, y: Int) -> Int {
	if x-1 < 0 || y-1 < 0 || x+1 >= m.count || y+1 >= m[x].count {
		return 0
	}

	var count = 0
	if m[x-1][y-1] == "M" && m[x+1][y+1] == "S" {
		count += 1
	}
	if m[x-1][y-1] == "S" && m[x+1][y+1] == "M" {
		count += 1
	}
	if m[x+1][y-1] == "M" && m[x-1][y+1] == "S" {
		count += 1
	}
	if m[x+1][y-1] == "S" && m[x-1][y+1] == "M" {
		count += 1
	}

	if count == 2 {
		return 1
	}

	return 0
}