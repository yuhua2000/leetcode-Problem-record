package leetcode

/*
 * @lc app=leetcode.cn id=1367 lang=golang
 *
 * [1367] 二叉树中的链表
 */

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubPath(head *ListNode, root *TreeNode) bool {
	if head == nil {
		return true
	}

	if root == nil {
		return false
	}

	var dfs func(head *ListNode, root *TreeNode) bool

	dfs = func(head *ListNode, root *TreeNode) bool {
		if head == nil {
			return true
		}

		if root == nil {
			return false
		}
		if head.Val == root.Val {
			return dfs(head.Next, root.Left) || dfs(head.Next, root.Right)
		}

		return false
	}

	if head.Val == root.Val {
		if dfs(head.Next, root.Left) || dfs(head.Next, root.Right) {
			return true
		}
	}

	return isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

// @lc code=end
