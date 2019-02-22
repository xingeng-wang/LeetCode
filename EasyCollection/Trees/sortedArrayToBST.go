package Trees

// Given an array where elements are sorted in ascending order, convert it to a height balanced BST.
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/94/trees/631/

// Definition for a binary tree node..
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	length := len(nums)
	if length == 0 {
		return nil
	} else if length == 1 {
		return &TreeNode{Val: nums[0]}
	} else if length == 2 {
		right := &TreeNode{Val: nums[1]}
		return &TreeNode{Val: nums[0], Right: right}
	} else if length == 3 {
		left := &TreeNode{Val: nums[0], Left: nil, Right: nil}
		right := &TreeNode{Val: nums[2], Left: nil, Right: nil}
		return &TreeNode{Val: nums[1], Left: left, Right: right}
	}
	middle := int(length / 2)
	left := sortedArrayToBST(nums[0:middle])
	right := sortedArrayToBST(nums[middle+1:])
	return &TreeNode{Val: nums[middle], Left: left, Right: right}
}
