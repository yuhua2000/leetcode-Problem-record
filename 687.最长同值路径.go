/*
 * @lc app=leetcode.cn id=687 lang=golang
 *
 * [687] 最长同值路径
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
func longestUnivaluePath(root *TreeNode) int {
	var ans = 0
	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		left1, right1 := 0, 0
		if root.Left != nil && root.Val == root.Left.Val {
			left1 = left + 1
		}
		if root.Right != nil && root.Val == root.Right.Val {
			right1 = right + 1
		}
		if left1+right1 > ans {
			ans = left1 + right1
		}
		if left1 > right1 {
			return left1
		}
		return right1
	}
	dfs(root)
	return ans
}

// @lc code=end

