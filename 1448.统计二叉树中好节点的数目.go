/*
 * @lc app=leetcode.cn id=1448 lang=golang
 *
 * [1448] 统计二叉树中好节点的数目
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
func goodNodes(root *TreeNode) (result int) {
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root.Left != nil {
			if root.Left.Val >= root.Val {
				result++
			} else {
				root.Left.Val = root.Val
			}
			dfs(root.Left)
		}
		if root.Right != nil {
			if root.Right.Val >= root.Val {
				result++
			} else {
				root.Right.Val = root.Val
			}
			dfs(root.Right)
		}
	}
	dfs(root)
	return result + 1
}

// @lc code=end
