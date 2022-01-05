package is_one_bit_char

//https://leetcode.com/problems/1-bit-and-2-bit-characters/
func isOneBitCharacter(bits []int) bool {
	for i := 0; i <= len(bits)-1; i++ {
		if bits[i] == 1 {
			i++
			if i == len(bits)-1 {
				return false
			}
		}
	}
	return true
}
