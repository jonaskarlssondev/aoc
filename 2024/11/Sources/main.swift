import Foundation

let fileURL = URL(fileURLWithPath: "input.txt")
let contents = try String(contentsOf: fileURL)
var data = contents.components(separatedBy: " ")

var cache = [String: Int]()

func run(i: String, n: Int, max: Int) -> Int {
	if n == max {
		return 1
	}

	let mem = cache["\(i):\(n)"]
	if mem != nil {
		return mem!
	}
	
	var len = 0

	if i == "0" {
		len += run(i:"1", n: n+1, max: max)
	} else if Array(i).count % 2 == 0 {
		let arr = Array(i)
		let first = String(arr[0..<(arr.count/2)])
		let second = arr[(arr.count / 2)..<arr.count]
		let formattedSecond = String(Int(String(second)) ?? 0)
		len += run(i: first, n: n+1, max: max) + run(i: formattedSecond, n: n+1, max: max)
	} else {
		let num = Int(i) ?? 0
		len += run(i: String(num*2024), n: n+1, max: max)
	}

	cache["\(i):\(n)"] = len
	return len
}


func tot(data: [String], max: Int) -> Int {
	var s = 0
	for x in data {
		s += run(i: x, n: 0, max: max)
	}
	return s
}
print(tot(data: data, max: 25))
cache = [String:Int]()
print(tot(data: data, max: 75))