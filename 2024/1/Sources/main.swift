import Foundation

let fileURL = URL(fileURLWithPath: "input.txt")
let contents = try String(contentsOf: fileURL)
let data = contents.components(separatedBy: "\n")

var left: Array<Int> = []
var right: Array<Int> = []

for row in data {
	if row == "" {
		break
	}
	let d = row.components(separatedBy:"   ")
	left.append(Int(d[0]) ?? -1)
	right.append(Int(d[1]) ?? -1)
}

left.sort()
right.sort()

var sum = 0

for (index, el) in left.enumerated() {
	let val = el - right[index]
	sum += abs(val)
}

print(sum)


var similarity = 0
for l in left {
	let n = right.filter( { $0 == l}).count
	similarity += l * n
}

print(similarity)