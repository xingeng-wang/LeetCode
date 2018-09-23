package Array

// Given an array of integers, return indices of the two numbers such that they add up to a specific target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/546/
func twoSum(nums []int, target int) []int {
	for i, a := range nums {
		for j, b := range nums {
			if i == j {
				continue
			}
			if a+b == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
