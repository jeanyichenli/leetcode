package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	// root is nil
	if root == nil {
		return subRoot == nil
	}

	// Ensure root and subRoot are same trees
	if isSameTree(root, subRoot) {
		return true
	}

	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSameTree(a *TreeNode, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil || a.Val != b.Val {
		return false
	}

	return isSameTree(a.Left, b.Left) && isSameTree(a.Right, b.Right)
}
