package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
 * @lc app=leetcode.cn id=979 lang=golang
 *
 * [979] 在二叉树中分配硬币
 */

// @lc code=start

// Definition for a binary tree node.

func distributeCoins(root *TreeNode) int {
	move := 0
	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		moveLeft := 0
		moveRight := 0
		if root == nil {
			return 0
		}
		if root.Left != nil {
			moveLeft = dfs(root.Left)
		}
		if root.Right != nil {
			moveRight = dfs(root.Right)
		}
		move += ads(moveLeft) + ads(moveRight)
		return moveLeft + moveRight + root.Val - 1

	}
	dfs(root)
	return move
}

func ads(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// @lc code=end
