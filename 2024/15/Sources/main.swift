import Foundation

let fileURL = URL(fileURLWithPath: "input.txt")
let contents = try String(contentsOf: fileURL)
let mapInp = contents.components(separatedBy: "\n\n")[0]
let moveInp = contents.components(separatedBy: "\n\n")[1]

var map: [[String]] = [[String]]()

var m2 = mapInp.components(separatedBy: "\n")
for i in m2.indices {
	let arr = Array(m2[i])
	map.append([])
	for y in arr {
		if y == "#" {
			map[i].append("#")
			map[i].append("#")
		}
		if y == "O" {
			map[i].append("[")
			map[i].append("]")
		}
		if y == "." {
			map[i].append(".")
			map[i].append(".")
		}
		if y == "@" {
			map[i].append("@")
			map[i].append(".")
		}
	}
}

var r: (Int, Int) = (0, 0)

for j in map.indices {
	for i in map[j].indices {
		if map[j][i] == "@" {
			r = (i, j)
		}
	}
}

for c in Array(moveInp) {
	let res = walkTwo(c: String(c))

	if res {
		print(c)
		for y in map {
			print(y.joined())
		}
		print("----------")
	}
}

func walkOne(c: String) {
	var newX = r.0
	var newY = r.1
	var dx = 0
	var dy = 0
	switch(c) {
		case "<":
			newY = r.1 - 1
			dy = -1
			break
		case ">":
			newY = r.1 + 1
			dy = 1
			break
		case "^":
			newX = r.0 - 1
			dx = -1
			break
		case "v":
			newX = r.0 + 1
			dx = 1
			break
		default:
			print("Unknown char", c)
			return
	}
	let atNew = map[newX][newY]
	if atNew == "#" {
		return
	}

	if atNew == "." {
		map[r.0][r.1] = "."
		r = (newX, newY)
		map[newX][newY] = "@"
		return
	}

	// We know its a box
	var lx = newX + dx
	var ly = newY + dy
	var c2: String = map[lx][ly]
	while c2 == "O" {
		lx += dx
		ly += dy
		c2 = map[lx][ly]
	}
	if c2 == "." {
		map[lx][ly] = "O"
		map[newX][newY] = "@"
		r = (newX, newY)
		map[newX-dx][newY-dy] = "."
	}
}

func walkTwo(c: String) -> Bool {
	var newX = r.0
	var newY = r.1
	var dx = 0
	var dy = 0
	//print("i", newX, newY, dy, dx)
	switch(c) {
		case "<":
			newX = r.0 - 1
			dx = -1
			break
		case ">":
			newX = r.0 + 1
			dx = 1
			break
		case "^":
			newY = r.1 - 1
			dy = -1
			break
		case "v":
			newY = r.1 + 1
			dy = 1
			break
		default:
			print("Unknown char", c)
			return false
	}
	//print("a", newX, newY, dx, dy)
	let atNew = map[newY][newX]
	if atNew == "#" {
		return false
	}

	if atNew == "." {
		map[r.1][r.0] = "."
		r = (newX, newY)
		map[newY][newX] = "@"
		return false
	}

	// We know its a box
	var lx = newX + dx
	var ly = newY + dy
	var c2: String = map[ly][lx]
	//print("1", newX, newY, lx, ly)
	if dx != 0 {
		while c2 == "[" || c2 == "]" {
			lx += dx
			ly += dy
			c2 = map[ly][lx]
		}
		if c2 == "#" {
			return false
		}

		map[r.1][r.0] = "."
		map[newY][newX] = "@"
		r = (newX, newY)

		for i: Int in 1...abs(lx-newX) {
			let replace = map[newY][newX + (i*dx)]
			//print("t2", i, newX, newY, i*dx, newX + i*dx, replace)
			if replace == "[" {
				map[newY][newX + (i*dx)] = "]"
			} else if replace == "]" {
				map[newY][newX + (i*dx)] = "["
			} else if replace == "." {
				map[newY][newX + (i*dx)] = dx < 0 ? "[" : "]"
			}
		}

		return true
	} else {
		let t = moveIfPossible(c: map[newY][newX], x: r.0, y: r.1, dy: dy)
		if t {
			map[r.1][r.0] = "."
			map[newY][newX] = "@"
			r = (newX, newY)
		}

		return t
	}
}

var copy: [[String]] = map
func moveIfPossible(c: String, x: Int, y: Int, dy: Int) -> Bool {
	copy = map
	
	let d = c == "]" ? -1 : 1
	let p = canpush(d: d, x: x, y: y, dy: dy, dop: false)

	if p {
		print("can move")
		let _ = canpush(d: d, x: x, y: y, dy: dy, dop: true)
		map = copy
	}

	return p
}

func canpush(d: Int, x: Int, y: Int, dy: Int, dop: Bool) -> Bool {
	//print("CanP", d, x, y, dy)
	if map[y+2*dy][x] == "." && map[y+2*dy][x+d] == "." {
		print("two dots")
		if dop {
			print("2d", d, y+2*dy, x, x+d)
			copy[y+2*dy][x] = d == -1 ? "]" : "["
			copy[y+2*dy][x+d] = d == 1 ? "]" : "["
			copy[y+dy][x] = "."
			copy[y+dy][x+d] = "."
		}
		return true
	}

	if map[y+2*dy][x] == "#" || map[y+2*dy][x+d] == "#" {
		print("sharp")
		return false
	}

	if d == -1 && map[y+2*dy][x+d] == "[" && map[y+2*dy][x] == "]" {
		print("same block, standing right")
		let t = canpush(d: d, x: x, y: y+dy, dy: dy, dop: dop)
		if t && dop {
			print("sb, sr", d, y+2*dy, x, x+d)
			copy[y+2*dy][x] = d == -1 ? "]" : "["
			copy[y+2*dy][x+d] = d == 1 ? "]" : "["
			copy[y+dy][x] = "."
			copy[y+dy][x+d] = "."
		}
		return t
	}

	if d == 1 && map[y+2*dy][x] == "[" && map[y+2*dy][x+d] == "]" {
		print("same block, standing left")
		let t = canpush(d: d, x: x, y: y+dy, dy: dy, dop: dop)
		if t && dop {
			print("sb, sl", d, y+2*dy, x, x+d)
			copy[y+2*dy][x] = d == -1 ? "]" : "["
			copy[y+2*dy][x+d] = d == 1 ? "]" : "["
			copy[y+dy][x] = "."
			copy[y+dy][x+d] = "."
		}

		return t
	}

	print("..")
	// One or two other blocks are blocking
	var a = false
	var b = false
	if map[y+2*dy][x+d] == "." {
		a = true
	} else {
		a = canpush(d: d, x: x+d, y: y+dy, dy: dy, dop: dop)
	}

	if map[y+2*dy][x] == "." {
		b = true
	} else {
		b = canpush(d: -d, x: x, y: y+dy, dy: dy, dop: dop)
	}

	if a && b && dop {
			print("..", d, y+2*dy, x, x+d)
			copy[y+2*dy][x] = d == -1 ? "]" : "["
			copy[y+2*dy][x+d] = d == 1 ? "]" : "["
			copy[y+dy][x] = "."
			copy[y+dy][x+d] = "."
	}

	return a && b
}

var sum = 0
for j: Range<Array<[String]>.Index>.Element in map.indices {
	for i in map[j].indices {
		if map[j][i] == "[" {
			sum += 100 * j + i
		}
	}
}

print(sum)