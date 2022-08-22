package add_binary

// AddBinary https://leetcode.com/problems/add-binary/
func addBinary(a, b string) string {
	result := ""
	longestBinStrLength := 0
	stringIndexer := 1
	leastSigDigitA := ""
	leastSigDigitB := ""
	carryOver := "0"
	sum := ""
	lenA := len(a)
	lenB := len(b)
	if lenA > lenB {
		longestBinStrLength = lenA
	} else {
		longestBinStrLength = lenB
	}
	for stringIndexer <= longestBinStrLength+1 {
		if stringIndexer <= lenA {
			leastSigDigitA = a[lenA-stringIndexer : lenA-stringIndexer+1]
		} else {
			leastSigDigitA = "0"
		}
		if stringIndexer <= lenB {
			leastSigDigitB = b[lenB-stringIndexer : lenB-stringIndexer+1]
		} else {
			leastSigDigitB = "0"
		}

		if leastSigDigitA == "1" && leastSigDigitB == "1" {
			if carryOver == "0" {
				sum = "0"
			} else {
				sum = "1"
			}
			carryOver = "1"
		} else if leastSigDigitA == "0" && leastSigDigitB == "0" {
			if carryOver == "0" {
				sum = "0"
			} else {
				sum = "1"
			}
			carryOver = "0"
		} else {
			if carryOver == "0" {
				sum = "1"
				carryOver = "0"
			} else {
				sum = "0"
				carryOver = "1"
			}
		}
		if stringIndexer == longestBinStrLength+1 && sum == "0" {
			break
		}
		result = sum + result
		stringIndexer++
	}
	return result
}
