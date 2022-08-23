package osprey

/*
take = (["apple", "pear", "lemon", "orange"], 2)
# ["apple", "pear"]
*/

/*
take = (["apple", "pear", "lemon", "orange"], 5)
# ["apple", "pear", "lemon", "orange", "apple"]
*/

// continue returning elements until you reach num

func nStrings(strArr []string, num int) (result []string) {
	if len(strArr) == 0 {
		return nil
	}
	for i := 0; i < num/len(strArr); i++ {
		result = append(result, strArr...)
	}
	return append(result, strArr[:num%len(strArr)]...)
}
