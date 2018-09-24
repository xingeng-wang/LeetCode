package Array

// Given an array, rotate the array to the right by k steps, where k is non-negative.
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/646/

func rotate(nums []int, k int) {
	length := len(nums)
	if length < k {
		k = k % length
	}
	index := length - k
	copy(nums, append(nums[index:], nums[:index]...))
}
