package String

// Write a function that takes a string as input and returns the string reversed.
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/127/strings/879/
func reverseString(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}
