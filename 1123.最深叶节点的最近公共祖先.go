package leetcode

/*
 * @lc app=leetcode.cn id=1123 lang=golang
 *
 * [1123] 最深叶节点的最近公共祖先
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	var f func(*TreeNode) (int, *TreeNode)
	f = func(root *TreeNode) (int, *TreeNode) {
		if root == nil {
			return 0, root
		}
		leftHight, leftLoc := f(root.Left)
		rightHeight, rightLoc := f(root.Right)
		if leftHight > rightHeight {
			return leftHight + 1, leftLoc
		}
		if leftHight < rightHeight {
			return rightHeight + 1, rightLoc
		}
		return leftHight + 1, root
	}
	_, root = f(root)
	return root
}

// @lc code=end
