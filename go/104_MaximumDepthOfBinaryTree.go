package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
	depth := 0

	return calculateMaxDepth(root, depth)
}

func calculateMaxDepth(root *TreeNode, depth int) int {

	if root == nil {
		return depth
	}

	depth++

	left_depth := calculateMaxDepth(root.Left, depth)
	right_depth := calculateMaxDepth(root.Right, depth)

	return max(left_depth, right_depth)
}
