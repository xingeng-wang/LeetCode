package Trees

// Given a binary tree, find its maximum depth.
// https://leetcode.com/explore/interview/card/top-interview-questions-easy/94/trees/555/
func maxDepth(root *TreeNode) int {
	var l, r int
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left != nil {
		l = 1 + maxDepth(root.Left)
	}

	if root.Right != nil {
		r = 1 + maxDepth(root.Right)
	}

	if l > r {
		return l
	}
	return r
}
