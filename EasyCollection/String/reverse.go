package String

import (
	"math"
	"strconv"
)

//Given a 32-bit signed integer, reverse digits of an integer.
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/127/strings/880/
func reverse(x int) int {
	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}
	xString := strconv.Itoa(x)
	chars := []rune(xString)
	if x > 0 {
		for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
			chars[i], chars[j] = chars[j], chars[i]
		}
	} else {
		for i, j := 1, len(chars)-1; i < j; i, j = i+1, j-1 {
			chars[i], chars[j] = chars[j], chars[i]
		}
	}
	xString = string(chars)
	num, _ := strconv.Atoi(xString)
	if num > math.MaxInt32 || num < math.MinInt32 {
		return 0
	}
	return num
}
