import Foundation

let fileURL = URL(fileURLWithPath: "input.txt")
let contents = try String(contentsOf: fileURL)
let data = contents.components(separatedBy: "\n\n")

var rules = data[0].components(separatedBy: "\n")
var pageSet = data[1].components(separatedBy: "\n")

// Works because all of the rules are in there
func sort(first: String, second: String) -> Bool {
	for rule in rules {
		let r = rule.components(separatedBy: "|")
		if r[0] == first && r[1] == second {
			return true
		}
	}
	return false
}

func isSorted(x: [String], f: (_: String, _: String) -> Bool) -> Bool {
	for i in x.indices.dropLast() {
		var j = i + 1
		while j < x.count {
			if !f(x[i], x[j]) {
				return false
			}
			j += 1
		}
	}

	return true
}

var sum = 0
var sum2 = 0
for book in pageSet {
	var pages = book.components(separatedBy: ",")
	let res = isSorted(x: pages, f: sort)

	if res {
		sum += Int(pages[pages.count / 2]) ?? 0
	} else {
		pages.sort(by: sort)
		sum2 += Int(pages[pages.count / 2]) ?? 0
	}
}

print(sum)
print(sum2)