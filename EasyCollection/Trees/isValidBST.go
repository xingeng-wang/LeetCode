package Trees

// Given a binary tree, determine if it is a valid binary search tree (BST).
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/94/trees/625/

// Definition for a binary tree node..
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func smallerThanSmallest(root *TreeNode, val int) bool {
	if root.Val <= val {
		return false
	} else {
		if root.Left != nil {
			return smallerThanSmallest(root.Left, val)
		} else {
			return true
		}
	}
}

func biggerThanBiggest(root *TreeNode, val int) bool {
	if root.Val >= val {
		return false
	} else {
		if root.Right != nil {
			return biggerThanBiggest(root.Right, val)
		} else {
			return true
		}
	}
}

func isValidBST(root *TreeNode) bool {
	var rightResult, leftResult bool
	var s, b bool
	if root == nil {
		return true
	}
	if root.Right == nil {
		rightResult = true
		s = true
	} else {
		if root.Right.Val > root.Val {
			s = smallerThanSmallest(root.Right, root.Val)
			rightResult = isValidBST(root.Right)
		} else {
			return false
		}
	}
	if root.Left == nil {
		leftResult = true
		b = true
	} else {
		if root.Left.Val < root.Val {
			b = biggerThanBiggest(root.Left, root.Val)
			leftResult = isValidBST(root.Left)
		} else {
			return false
		}
	}
	return rightResult && leftResult && s && b
}
