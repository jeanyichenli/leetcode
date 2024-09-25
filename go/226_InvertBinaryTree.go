package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func invertTree(root *TreeNode) *TreeNode {
	exchange(root)
	return root
}

func exchange(root *TreeNode) {
	if root == nil {
		return
	}

	left, right := root.Left, root.Right

	root.Left = right
	root.Right = left

	exchange(root.Left)
	exchange(root.Right)

}
