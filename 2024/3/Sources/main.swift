import Foundation

let fileURL = URL(fileURLWithPath: "input.txt")
var data = try String(contentsOf: fileURL)

extension String {
	func indexOf(_ input: String,
                 options: String.CompareOptions = .literal) -> String.Index? {
        return self.range(of: input, options: options)?.lowerBound
    }

    func lastIndexOf(_ input: String) -> String.Index? {
        return indexOf(input, options: .backwards)
    }
}

var sum = 0
var enabledSum = 0
var enabled = true

var scanner = Scanner(string: data)
while scanner.scanLocation < data.count {
	var scrap: NSString?
	scanner.scanUpTo("mul(", into: &scrap)
	if scanner.scanLocation == data.count {
		break
	}

	if scrap != nil {
		let doIdx = String(scrap!).lastIndexOf("do()")
		let dontIndex = String(scrap!).lastIndexOf("don't()")

		if doIdx != nil || dontIndex != nil {
			if doIdx != nil && dontIndex == nil {
				if !enabled {
					enabled = true
				}
			} else if doIdx == nil && dontIndex != nil {
				if enabled {
					enabled = false
				}
			} else {
				if doIdx! > dontIndex! && !enabled {
					enabled = true
				}
				if doIdx! < dontIndex! && enabled {
					enabled = false
				}
			}
		}
	}
	if enabled && scrap != nil && scrap!.contains("don't()") {
		enabled = false
	}

	var first = ""
	var ffound = false
	for i in 0...3 {
		let idx = data.index(data.startIndex, offsetBy: scanner.scanLocation + i + 4) // 4 = "mul("
		let str = String(data[idx])
		if str == "," {
			ffound = true
			break
		}
		if Int(str) == nil {
			break
		}
		first += str
	}

	if first != "" {
		scanner.scanLocation += first.count + 4
	}

	if first != "" && ffound {
		scanner.scanLocation += 1 //1 for ','

		var second = ""
		var sfound = false
		for i in 0...3 {
			let idx = data.index(data.startIndex, offsetBy: scanner.scanLocation + i)
			let str = String(data[idx])
			if str == ")" {
				sfound = true
				break
			}
			if Int(str) == nil {
				break
			}
			second += str
		}

		scanner.scanLocation += second.count
		if second != "" && sfound{
			let f = Int(first)
			let s = Int(second)

			sum += f!*s!

			if enabled {
				enabledSum += f!*s!
			}
		}
	}

	let i = data.index(data.startIndex, offsetBy: scanner.scanLocation)
	data = String(data[i...])
	scanner = Scanner(string: data)
}

print(sum)
print(enabledSum)