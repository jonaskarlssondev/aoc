import Foundation

let fileURL = URL(fileURLWithPath: "input.txt")
let contents = try String(contentsOf: fileURL)
let data = contents.components(separatedBy: "\n")


struct Machine {
	var a: (Int, Int)
	var b: (Int, Int)
	var t: (Int64, Int64)
}

func run(two: Bool) throws {
	if #available(macOS 14.0.0, *) {
		let buttonRegex: Regex<AnyRegexOutput> = try Regex("(?:A|B): X\\+(\\d+), Y\\+(\\d+)")
		let targetRegex = try Regex("X\\=(\\d+), Y\\=(\\d+)")

		var i=0

		var machines = [Machine]()

		while i < data.count {
			var a = (0,0)
			if let f = data[i].firstMatch(of: buttonRegex) {
				let f1: String = String(describing: f[1].value!)
				let f2: String = String(describing: f[2].value!)
				a = (Int(f1) ?? 0, Int(f2) ?? 0)
			}

			var b = (0,0)
			if let f = data[i+1].firstMatch(of: buttonRegex) {
				let f1: String = String(describing: f[1].value!)
				let f2: String = String(describing: f[2].value!)
				b = (Int(f1) ?? 0, Int(f2) ?? 0)
			}

			var t:(Int64, Int64) = (0,0)
			if let f = data[i+2].firstMatch(of: targetRegex) {
				let f1: String = String(describing: f[1].value!)
				let f2: String = String(describing: f[2].value!)
				t = (Int64(f1) ?? 0, Int64(f2) ?? 0)
				if two {
					t = (10000000000000 + t.0, 10000000000000 + t.1)
				}
			}

			machines.append(Machine(a: a, b: b, t: t))

			i += 4
		}

		var sum = 0
		for m in machines {
			// See line - line intersection on wikipedia
			let pressB: Double = (Double(m.t.1) * Double(m.a.0) - Double(m.t.0) * Double(m.a.1)) / Double(Double(m.b.1) * Double(m.a.0) - Double(m.b.0) * Double(m.a.1))
			let pressA: Double = Double(Double(m.t.0) - pressB * Double(m.b.0)) / Double(m.a.0)

			if !two && (pressA > 100 || pressB > 100) {
				continue
			}

			if pressA < 0 || pressB < 0{
				continue
			}
			if pressA.truncatingRemainder(dividingBy: 1) != 0 || pressB.truncatingRemainder(dividingBy: 1) != 0 {
				continue
			}

			sum += Int(3*pressA + pressB)
		}

		print(sum)
	}
}

try run(two: false)
try run(two: true)

