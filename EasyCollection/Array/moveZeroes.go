package Array

// Move Zeroes, Given an array nums, write a function to move all 0's to the end
// of it while maintaining the relative order of the non-zero elements.
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/567/
func moveZeroes(nums []int) {
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, 0)
		}
	}
}
