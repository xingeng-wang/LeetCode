package Array

import "sort"

// Given an array of integers, find if the array contains any duplicates.
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/578/
func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	l := len(nums)
	for i, num := range nums {
		if i == l-1 {
			return false
		}
		if num == nums[i+1] {
			return true
		}
	}
	return false
}
