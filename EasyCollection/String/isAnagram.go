package String

// Given two strings s and t , write a function to determine if t is an anagram of s.
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/127/strings/882/
func isAnagram(s string, t string) bool {
	sLength := len(s)
	if sLength != len(t) {
		return false
	}
	if sLength == 0 {
		return true
	}
	var sNumber, tNumber rune
	var xor rune
	for _, c := range s {
		sNumber = sNumber + c*c
		xor = xor ^ c
	}

	for _, c := range t {
		tNumber = tNumber + c*c
		xor = xor ^ c
	}
	return tNumber == sNumber && xor == 0
}
