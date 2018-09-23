package Array

// Given a non-empty array of digits representing a non-negative integer, plus one to the integer.
// The digits are stored such that the most significant digit is at the head of the list, and each element in the array contain a single digit.
// You may assume the integer does not contain any leading zero, except the number 0 itself.
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/559/

import "strconv"

func plusOne(digits []int) []int {
	numInString := ""
	for _, digit := range digits {
		numInString = numInString + strconv.Itoa(digit)
	}
	num, _ := strconv.ParseInt(numInString, 10, 64)
	num += 1
	numInString = strconv.FormatInt(num, 10)
	digits = []int{}
	d := 0
	for _, n := range numInString {
		d, _ = strconv.Atoi(string(n))
		digits = append(digits, d)
	}
	return digits
}
