package Array

// Given a non-empty array of integers, every element appears twice except for one. Find that single one.
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/549/
func singleNumber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	result := 0
	for _, num := range nums {
		result = result ^ num
	}
	return result
}
