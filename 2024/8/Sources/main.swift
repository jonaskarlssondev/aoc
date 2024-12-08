import Foundation

let fileURL = URL(fileURLWithPath: "input.txt")
let contents = try String(contentsOf: fileURL)
let data = contents.components(separatedBy: "\n")

var antennas: [[String]] = [[String]]()

for i in data.indices {
	let arr = Array(data[i])
	antennas.append([])
	for y in arr {
		antennas[i].append(String(y))
	}
}

func run(a: [[String]], two: Bool) {
	var freq = [String: Bool]()

	for j in antennas.indices {
		for i in antennas[j].indices {
			let char = antennas[j][i]
			if char != "." && char != "#" {
				for y in antennas.indices {
					for x in antennas[y].indices {
						let char2 = antennas[y][x]
						if char == char2 && (y > j || (x > i && y == j)) {
							let dx = abs(i - x)
							let dy = abs(j - y)

							if two {
								freq["\(i):\(j)"] = true
								freq["\(x):\(y)"] = true
							}

							var multf = 1
							while true {
								var fx = 0
								var fy = 0
								if i >= x && j <= y {
									fx = max(i, x) + dx * multf
									fy = min(j, y) - dy * multf
								} else {
									fx = min(i, x) - dx * multf
									fy = min(j, y) - dy * multf
								}

								if fx >= 0 && fx < antennas.count && fy >= 0 && fy < antennas[0].count {
									freq["\(fx):\(fy)"] = true
								} else {
									break
								}

								multf += 1

								if !two {
									break
								}
							}

							var mults = 1
							while true {
								var sx = 0
								var sy = 0
								if i >= x && j <= y {
									sx = min(i, x) - dx * mults
									sy = max(j, y) + dy * mults
								} else {
									sx = max(i, x) + dx * mults
									sy = max(j, y) + dy * mults
								}

								if sx >= 0 && sx < antennas.count && sy >= 0 && sy < antennas[0].count {
									freq["\(sx):\(sy)"] = true
								} else {
									break
								}

								mults += 1

								if !two {
									break
								}
							}
						}
					}
				}
			}
		}
	}


	var antennas2 = antennas

	for j in antennas.indices {
		for i in antennas[j].indices {
			if freq["\(j):\(i)"] != nil {
				antennas2[i][j] = "#"
			}
		}
	}

	printM(antennas2)
	print(freq.count)
}

func printM(_ m: [[String]]) {
	for x: [String] in m {
		print(x.joined())
	}
}

run(a: antennas, two: false)
run(a: antennas, two: true)