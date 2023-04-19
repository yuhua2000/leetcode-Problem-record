package leetcode

/*
 * @lc app=leetcode.cn id=1026 lang=golang
 *
 * [1026] 节点与其祖先之间的最大差值
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
// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func maxAncestorDiff(root *TreeNode) (result int) {
	var dfs func(*TreeNode)
	queue := make([]int, 0)
	dfs = func(node *TreeNode) {
		defer func() {
			queue = queue[:len(queue)-1]
		}()
		for _, v := range queue {
			ans := v - node.Val
			if ans < 0 {
				ans = -ans
			}
			if ans > result {
				result = ans
			}
		}
		queue = append(queue, node.Val)

		if node.Left != nil {
			dfs(node.Left)
		}
		if node.Right != nil {
			dfs(node.Right)
		}
	}
	dfs(root)
	return result
}

// @lc code=end
