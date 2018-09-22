package Array

// Remove Duplicates from Sorted Array
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/727/
func removeDuplicates(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return l
	}
	for i := l - 1; i > 0; i-- {
		if nums[i] == nums[i-1] {
			nums = append(nums[:i], nums[i+1:]...)
		}
	}
	return len(nums)
}
