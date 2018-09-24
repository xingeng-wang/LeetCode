package Math

// Given an integer, write a function to determine if it is a power of three.
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/102/math/745/
func isPowerOfThree(n int) bool {
	if n == 1 {
		return true
	}
	if n < 3 {
		return false
	}
	for i := n; i > 0; {
		if i == 3 {
			return true
		} else if i%3 == 0 {
			i = i / 3
		} else {
			return false
		}
	}
	return true
}
