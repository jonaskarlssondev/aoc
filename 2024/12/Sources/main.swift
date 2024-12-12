import Foundation

let fileURL = URL(fileURLWithPath: "input.txt")
let contents = try String(contentsOf: fileURL)
let data = contents.components(separatedBy: "\n")

var plots: [[String]] = [[String]]()

for i in data.indices {
	let arr = Array(data[i])
	plots.append([])
	for y in arr {
		plots[i].append(String(y))
	}
}

var sum = 0
var sum2 = 0
var checked = [String:Bool]()
for a in plots.indices {
	for b in plots[a].indices {
		if checked["\(a):\(b)"] == nil {
			let x = fill(m: plots, i: a, j: b, x: plots[a][b])
			sum += x.0 * x.1
			sum2 += x.0 * x.2
		}
	}
}

print(sum)
print(sum2)


func fill(m: [[String]], i: Int, j:Int, x: String) -> (Int, Int, Int) {
	if i < 0 || j < 0 || i >= m.count || j >= m[i].count {
		return (0,0,0)
	}

	if checked["\(i):\(j)"] != nil {
		return (0,0,0)
	}

	if m[i][j] != x {
		return (0,0,0)
	}

	checked["\(i):\(j)"] = true

	var p = 0
	var s = 1
	var c = 0

	if inside(m:m, i: i-1, j:j) && m[i-1][j] == x {
		let a = fill(m:m, i: i-1, j: j, x: x)
		s += a.0
		p += a.1
		c += a.2
	} else {
		p += 1
	}

	if inside(m:m, i: i+1, j:j) && m[i+1][j] == x {
		let a = fill(m:m, i: i+1, j: j, x: x)
		s += a.0
		p += a.1
		c += a.2
	} else {
		p += 1
	}

	if inside(m:m, i: i, j:j-1) && m[i][j-1] == x {
		let a = fill(m:m, i: i, j: j-1, x: x)
		s += a.0
		p += a.1
		c += a.2
	} else {
		p += 1
	}

	if inside(m:m, i: i, j:j+1) && m[i][j+1] == x {
		let a =  fill(m:m, i: i, j: j+1, x: x)
		s += a.0
		p += a.1
		c += a.2
	} else {
		p += 1
	}
	
	let optN = other(m:m, i: i-1, j: j, x: x)
	let optNW = other(m:m, i: i-1, j: j-1, x: x)
	let optW = other(m:m, i: i, j: j-1, x: x)
	let optSW = other(m:m, i: i+1, j: j-1, x: x)
	let optS = other(m:m, i: i+1, j: j, x: x)
	let optSE = other(m:m, i: i+1, j: j+1, x: x)
	let optE = other(m:m, i: i, j: j+1, x: x)
	let optNE = other(m:m, i: i-1, j: j+1, x: x)

	// Edge corner
	if optW && optN {
		c += 1
	} 
	if optN && optE {
		c += 1
	} 
	if optE && optS {
		c += 1
	} 
	if optS && optW {
		c += 1
	} 
	
	// Inner corner
	if !optW && optNW && !optN {
		c += 1
	}
	if !optN && optNE && !optE {
		c += 1
	}	
	if !optE && optSE && !optS {
		c += 1
	}
	if !optS && optSW && !optW {
		c += 1
	}

	return (s,p,c)
}

func other(m: [[String]], i: Int, j:Int, x: String) -> Bool {
	return !inside(m: m, i: i, j: j) || m[i][j] != x
}

func inside(m: [[String]], i: Int, j:Int) -> Bool {
	if i < 0 || j < 0 || i >= m.count || j >= m[i].count {
		return false
	}

	return true
}