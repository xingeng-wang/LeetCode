package String

import (
	"regexp"
	"strings"
)

// Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/127/strings/883/
func isPalindrome(s string) bool {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	newString := strings.ToLower(reg.ReplaceAllString(s, ""))
	chars := []rune(newString)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	reverse := string(chars)
	return reverse == newString
}
