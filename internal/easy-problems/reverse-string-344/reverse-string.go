package reverse_string_344

// https://leetcode.com/problems/reverse-string/
func reverseString(s []byte) {
	for i := 0; i < len(s)/2; i++ {
		temp := s[i]
		s[i], s[len(s)-1-i] = s[len(s)-1-i], temp
	}
}

func reverseString1(str string) string {
	strAsRuneArr := []rune(str)
	var reversedString []rune

	for i := len(strAsRuneArr) - 1; i >= 0; i-- {
		reversedString = append(reversedString, []rune(str)[i])
	}

	return string(reversedString)
}

func reverseString2(str string) string {
	strRuneArray := []rune(str)

	for i := 0; i <= len(strRuneArray)/2; i++ {
		temp := strRuneArray[i]
		strRuneArray[i] = strRuneArray[len(strRuneArray)-1-i]
		strRuneArray[len(strRuneArray)-1-i] = temp
	}

	return string(strRuneArray)
}
