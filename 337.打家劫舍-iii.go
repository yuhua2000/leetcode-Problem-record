//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * @lc app=leetcode.cn id=337 lang=golang
 *
 * [337] 打家劫舍 III
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
func rob(root *TreeNode) int {
	var f func(root *TreeNode) []int
	f = func(root *TreeNode) []int {
		if root == nil {
			return []int{0, 0}
		}

		l, r := f(root.Left), f(root.Right)
		selected := root.Val + l[1] + r[1]
		notSelected := max(l[0], l[1]) + max(r[0], r[1])
		return []int{selected, notSelected}
	}
	val := f(root)
	return max(val[0], val[1])
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
