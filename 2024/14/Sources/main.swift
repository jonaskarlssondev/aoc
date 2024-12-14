import Foundation

let fileURL = URL(fileURLWithPath: "input.txt")
let contents = try String(contentsOf: fileURL)
let data = contents.components(separatedBy: "\n")

struct Robot {
    var x: Int
    var y: Int
    var dx: Int
    var dy: Int
}

var robots = [Robot]()
for s in data {
    var pos = s.components(separatedBy: " ")[0]
    pos = String(Array(pos)[2...])
    let p = pos.components(separatedBy: ",")

		var delta = s.components(separatedBy: " ")[1]
    delta = String(Array(delta)[2...])
    let d = delta.components(separatedBy: ",")

    robots.append(Robot(x: Int(p[0]) ?? 0, y: Int(p[1]) ?? 0, dx: Int(d[0]) ?? 0, dy: Int(d[1]) ?? 0))
}

func walk(r: Robot, w: Int, h: Int) -> (Int, Int) {
	var nx = r.x + r.dx
	var ny = r.y + r.dy

	if nx >= w {
		nx = nx - w
	}
	if nx < 0 {
		nx = w + nx
	}
	if ny >= h {
		ny = ny - h
	}
	if ny < 0 {
		ny = h + ny
	}

	return (nx, ny)
}

let width = 101
let height = 103

let part1Range = 100
let part2Range = 10000
for n in 1...part2Range {
	for i in robots.indices {
		let res = walk(r: robots[i], w: width, h: height)
		robots[i].x = res.0
		robots[i].y = res.1
	}

	if nOnSameLine(r: robots, w: width, h: height, n: 10) {
		printM(r: robots, w: width, h: height)
		print(n)
		break
	}
}

var s1 = 0
var s2 = 0
var s3 = 0
var s4 = 0
for r in robots {
	if r.x > width / 2 {
		if r.y > height / 2 {
			s4 += 1
		}
		if r.y < height / 2 {
			s3 += 1
		}
	}
	if r.x < width / 2 {
		if r.y > height / 2 {
			s2 += 1
		}
		if r.y < height / 2 {
			s1 += 1
		}
	}
}

print(s1*s2*s3*s4)

func printM(r: [Robot], w: Int, h: Int) {
	for y in 0..<w {
		for x in 0..<h {
			var found = false
			for a in r {
				if a.x == x && a.y == y {
					found = true
				}
			}
			if found {
				print("*", terminator: "")
			} else {
				print(" ", terminator: "")
			}
		}
		print("")
	}
}

func nOnSameLine(r: [Robot], w: Int, h: Int, n: Int) -> Bool {
	var m = [[Int]]()
	for i in 0..<w {
		m.append([Int]())
		for _ in 0..<h {
			m[i].append(0)
		}
	}

	for x in r {
		m[x.x][x.y] += 1
	}

	for j in m.indices {
		for i in m[j].indices {
			if m[j][i] > 0 && nOnSame(x: m[j], s: i) > n {
				return true
			}
		}
	}

	return false
}

func nOnSame(x: [Int], s: Int) -> Int {
	var i = s
	while i < x.count && x[i] > 0 {
		i += 1
	}

	return i - s
}