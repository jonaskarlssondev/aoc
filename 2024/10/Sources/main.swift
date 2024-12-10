import Foundation

let fileURL = URL(fileURLWithPath: "input.txt")
let contents = try String(contentsOf: fileURL)
let data = contents.components(separatedBy: "\n")

var arr = [[Character]]()

for row in data {
	arr.append(Array(row))
}

var targetWord = Array("0123456789")

var seen = [String: Bool]()

func search(m: [[Character]], ogx: Int, ogy:Int, x: Int, y: Int, t: Int) -> Int {
	if x < 0 || y < 0 || x >= m.count || y >= m[x].count {
		return 0
	}

	if m[x][y] != targetWord[t] {
		return 0
	}

	if t+1 == targetWord.count {
		seen["\(ogy):\(ogx):\(y):\(x)"] = true
		return 1
	}

	let directions: [Int] = [-1,0,1]

	var s = 0
	for i in directions {
		for j in directions {
			if i != j && (i == 0 || j == 0) {
				s += search(m: m, ogx: ogx, ogy: ogy, x: x + i, y: y + j, t: t+1)
			}
		}
	}

	return s
}


var sum = 0

for i in arr.indices {
	for j: Range<Array<Character>.Index>.Element in arr[i].indices {
		if arr[i][j] == "0" {
			sum += search(m:arr, ogx: i, ogy: j, x:i, y:j, t: 0)
		}
	}
}

print(seen.count)
print(sum)
