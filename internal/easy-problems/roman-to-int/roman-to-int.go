package roman_to_int

// https://leetcode.com/problems/roman-to-integer/
func romanToInt(s string) int {
	result := 0
	for i := range s {
		switch string(s[i]) {
			case `I`:
				result += 1
			case `V`:
				result += 5
				if i != 0 && string(s[i-1]) == `I` {
					result -= 2
				}
			case `X`:
				result += 10
				if i != 0 && string(s[i-1]) == `I` {
					result -= 2
				}
			case `C`:
				result += 100
				if i != 0 && string(s[i-1]) == `X` {
					result -= 20
				}
			case `L`:
				result += 50
				if i != 0 && string(s[i-1]) == `X` {
					result -= 20
				}
			case `D`:
				result += 500
				if i != 0 && string(s[i-1]) == `C` {
					result -= 200
				}
			case `M`:
				result += 1000
				if i != 0 && string(s[i-1]) == `C` {
					result -= 200
				}
		}
	}
	return result
}
