import Foundation

let fileURL = URL(fileURLWithPath: "input.txt")
let contents = try String(contentsOf: fileURL)
let data = contents.components(separatedBy: "\n")

func solvable(x: [Int], res: Int) -> Bool {
	if x.count < 2 {
		return false
	}

	var q = x

	let f = q.removeFirst()
	let s = q.removeFirst()

	let concat = Int(String(f) + String(s)) ?? 0

	if q.count == 0 && (f+s == res || f*s == res || concat == res) {
		return true 
	}

	var a = q
	a.insert(f+s, at: 0)

	var b = q
	b.insert(f*s, at: 0)

	var c = q
	c.insert(concat, at: 0)

	return solvable(x: a, res: res) || solvable(x:b, res: res) || solvable(x:c, res: res)
}


var sum = 0

for row in data {
	let a = row.components(separatedBy: ": ")
	let b = a[1].components(separatedBy: " ").map({Int($0) ?? 0})

	let res = Int(a[0]) ?? 0

	if solvable(x: b, res: res) {
		sum += res
	}
}

print(sum)