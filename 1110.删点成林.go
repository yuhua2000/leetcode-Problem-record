/*
 * @lc app=leetcode.cn id=1110 lang=golang
 *
 * [1110] 删点成林
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

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	toDeleteSet := make(map[int]bool)
	for _, i := range to_delete {
		toDeleteSet[i] = true
	}
	var roots []*TreeNode
	var dfs func(*TreeNode, bool) *TreeNode
	dfs = func(root *TreeNode, isRoot bool) *TreeNode {
		if root == nil {
			return nil
		}
		deleted := toDeleteSet[root.Val]
		root.Left = dfs(root.Left, deleted)
		root.Right = dfs(root.Right, deleted)
		if deleted {
			return nil
		} else {
			if isRoot {
				roots = append(roots, root)
			}
			return root
		}

	}
	dfs(root, true)
	return roots
}

// @lc code=end
