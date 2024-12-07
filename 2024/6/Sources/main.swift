import Foundation

let fileURL = URL(fileURLWithPath: "input.txt")
let contents = try String(contentsOf: fileURL)
let data = contents.components(separatedBy: "\n")

var walkable: [[Bool]] = [[Bool]]()
var visited = [[Bool]]()

var sx: Int = 0
var sy: Int = 0

for i in data.indices {
	let arr = Array(data[i])
	walkable.append([])
	visited.append([])
	for j in arr.indices {
		if arr[j] == "^" {
			sx = j
			sy = i
			visited[i].append(false)
			walkable[i].append(true)
			continue
		}

		visited[i].append(false)
		walkable[i].append(arr[j] == ".")
	}
}

func walk(w: [[Bool]], vis: [[Bool]], startX: Int, startY: Int) -> Int {
	var v = vis
	var x = startX
	var y = startY
	var dx = 0
	var dy = -1

	var seen = [String:Bool]()

	while true {
		let nextOutside = (x+dx) < 0 || (x+dx) >= w.count || (y+dy) < 0 || (y+dy) >= w[0].count
		if nextOutside {
			v[x][y] = true
			break
		}

		let key: String = "\(x):\(y):\(dx):\(dy)"
		if seen[key] != nil {
			return 1
		} else {
			seen[key] = true
		}

		if w[y+dy][x+dx] {
			v[x][y] = true

			x = x+dx
			y = y+dy
		} else {
			if dx == 0 {
				if dy == -1 {
					dx = 1
					dy = 0
				} else {
					dx = -1
					dy = 0
				}
				continue
			}

			if dy == 0 {
				if dx == -1 {
					dx = 0
					dy = -1
				} else {
					dx = 0
					dy = 1
				}
			}
		}
	}

	var sum = 0
	for x in v {
		for y in x {
			if y {
				sum += 1
			}
		}
	}

	return sum
}

print(walk(w: walkable, vis: visited, startX: sx, startY: sy))

var sum2 = 0
for i: Range<Array<[Bool]>.Index>.Element in walkable.indices {
	for j in walkable[0].indices {
		if walkable[i][j] {
			var copy = walkable
			copy[i][j] = false

			let res = walk(w: copy, vis: visited, startX: sx, startY: sy)
			if res == 1 {
				sum2 += 1
			}
		}
	}
}

print(sum2)
