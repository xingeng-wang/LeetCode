package Trees

// Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/94/trees/627/

// Definition for a binary tree node..
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(rootOne *TreeNode, rootTwo *TreeNode) bool {
	if rootOne == nil && rootTwo == nil {
		return true
	} else if rootOne == nil || rootTwo == nil {
		return false
	}
	if rootOne.Val == rootTwo.Val {
		return isMirror(rootOne.Left, rootTwo.Right) && isMirror(rootTwo.Left, rootOne.Right)
	}
	return false
}
