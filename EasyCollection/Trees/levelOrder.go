package Trees

// Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/94/trees/628/

// Definition for a binary tree node..
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	result := make([][]int, 1)
	result[0] = append(result[0], root.Val)
	result = updateResult(root.Left, 1, result)
	result = updateResult(root.Right, 1, result)
	return result
}

func updateResult(root *TreeNode, level int, result [][]int) [][]int {
	if root == nil {
		return result
	}
	if len(result) < level+1 {
		result = append(result, []int{root.Val})
	} else {
		result[level] = append(result[level], root.Val)
	}
	result = updateResult(root.Left, level+1, result)
	result = updateResult(root.Right, level+1, result)
	return result
}
