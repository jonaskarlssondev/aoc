import Foundation

let fileURL = URL(fileURLWithPath: "input.txt")
let contents = try String(contentsOf: fileURL)
let data = Array(contents)

var chunks = (data.count / 2)
var last: Int = data.count % 2 == 0 ? data.count - 2 : data.count - 1

var nums = 0
var a = 0
while a<data.count {
	nums += Int(String(data[a])) ?? 0
	a += 2
}

var gaps = 0
var b = 1
while b<data.count {
	gaps += Int(String(data[b])) ?? 0
	b += 2
}

func run(c: Int, l: Int, two: Bool) {

	if !two {
		var sum = 0
		var chunks = c
		var last = l

		var i = 1
		var j = 1

		var index = Int(String(data[0])) ?? 0

		var empty = Int(String(data[i])) ?? 0
		var numBlocks = Int(String(data[last])) ?? 0

		while i < last {
			while empty > 0 && numBlocks > 0 {
				sum += index * chunks
				index += 1

				empty -= 1
				numBlocks -= 1
			}
			
			if empty == 0 {
				let next = Int(String(data[i+1])) ?? 0
				for _ in 0..<next {
					if index >= nums {
						break
					}

					sum += index*j
					index += 1
				}
				i += 2
				j += 1
				empty = Int(String(data[i])) ?? 0
			}

			if numBlocks == 0 {
				last -= 2
				chunks -= 1
				numBlocks = Int(String(data[last])) ?? 0
			}
		}

		print(sum)
		return
	}
	
	//2333133121414131402
	var sum = 0
	var chunks = c

	var i: Int = 1
	var j = 1

	var index = Int(String(data[0])) ?? 0
	var numEmpty: Int = Int(String(data[i])) ?? 0

	var moved = [Int]()

	var deltaMove: Int = 0
	while chunks - deltaMove > 0 {
		var deltadeltaMove = deltaMove
		while numEmpty > 0 && (last-deltadeltaMove*2) > i {
			let deltaBlocks = Int(String(data[last-deltadeltaMove*2])) ?? 0
			if deltaBlocks <= numEmpty && !moved.contains((chunks - deltadeltaMove)) {
				moved.append((chunks - deltadeltaMove))
				for _ in 0..<deltaBlocks {
					sum += index * (chunks - deltadeltaMove)
					index += 1
					numEmpty -= 1
				}

				if deltaMove == deltadeltaMove {
					deltaMove += 1
				}

			}
			deltadeltaMove += 1
		}

		if numEmpty != 0 {
			index += numEmpty
		}

		i += 2
		if i >= data.count {
			break
		}
		numEmpty = Int(String(data[i])) ?? 0
		let x = Int(String(data[i-1])) ?? 0

		for _ in 0..<x {
			if !moved.contains(j) {
				sum += index * j
			}

			index += 1
		}
				
		j += 1
	}

	print(sum)
}

run(c: chunks, l: last, two: false)
run(c: chunks, l: last, two: true)