import Foundation

let fileURL = URL(fileURLWithPath: "input.txt")
let contents = try String(contentsOf: fileURL)
let data = contents.components(separatedBy: "\n")

var sum = 0

for row in data {
	if row == "" {
		break
	}
	let d = row.components(separatedBy:" ")
	
	let dir = (Int(d[1]) ?? 0) - (Int(d[0]) ?? 0)
	if dir == 0 {
		continue
	}
	
	var last = Int(d[0]) ?? 0
	var invalid = false
	for num in d.dropFirst() {
		let n = Int(num) ?? 0
		let diff = n - last

		if abs(diff) < 1 || abs(diff) > 3 {
			invalid = true
			break
		}

		if dir > 0 && diff <= 0 {
			invalid = true
			break
		}

		if dir < 0 && diff >= 0 {
			invalid = true
			break
		}

		last = n
	}

	if !invalid {
		sum += 1
	}
}

print(sum)

func safe(d: [Int]) -> Bool {
	var diff: Array<Int> = []

	for i in 1...d.count-1 {
		diff.append(d[i] - d[i-1])
	}

	return diff.allSatisfy({$0 >= 1 && $0 <= 3 }) || diff.allSatisfy({$0 <= -1 && $0 >= -3 })
}

var safe = 0

for row in data {
	if row == "" {
		break
	}
	let d = row.components(separatedBy:" ").map { Int($0) ?? 0 }

	if safe(d: d) {
		safe += 1
		continue
	}

	for idx in d.indices {
		var modified = d
		modified.remove(at: idx)

		if safe(d: modified) {
			safe += 1
			break
		}
	}
}

print(safe)

