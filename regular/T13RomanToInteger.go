package regular

func romanToInt(s string) int {
	dict := make(map[string]int)
	dict["IV"] = 4
	dict["IX"] = 9
	dict["XL"] = 40
	dict["XC"] = 90
	dict["CD"] = 400
	dict["CM"] = 900
	dict["I"] = 1
	dict["V"] = 5
	dict["X"] = 10
	dict["L"] = 50
	dict["C"] = 100
	dict["D"] = 500
	dict["M"] = 1000
	return solution(s, 0, dict)
}

func solution(remain string, result int, dict map[string]int) int {
	length := len(remain)
	if length >= 2 {
		val, prs := dict[remain[0:2]]
		if prs {
			result += val
			return solution(remain[2:], result, dict)
		}
	} else if length >= 1 {
		result += dict[remain[0:1]]
		return solution(remain[1:], result, dict)
	} else {
		return result
	}
	return result
}
