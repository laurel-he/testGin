package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	/*
		判断树的深度，可以采用递归，但是当深度达到一定程度会溢出，所以采用迭代的方法解决。
	*/
	if root == nil {
		return 0
	}
	maxDepthNum := 1
	for {
		if root.Left == nil || root.Right == nil {
			maxDepthNum++

		}
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func main() {

}
