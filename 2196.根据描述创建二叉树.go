package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func createBinaryTree(descriptions [][]int) *TreeNode {
	treeNodeGraph := make(map[int]*TreeNode)
	isChild := make(map[int]bool)
	for _, descriptions := range descriptions {
		parent, child, isLeft := descriptions[0], descriptions[1], descriptions[2] == 1

		parentNode, ok := treeNodeGraph[parent]
		if !ok {
			parentNode = &TreeNode{Val: parent}
			treeNodeGraph[parent] = parentNode
		}

		if _, ok := isChild[parent]; !ok {
			isChild[parent] = false
		}

		childNode, ok := treeNodeGraph[child]
		if !ok {
			childNode = &TreeNode{Val: child}
			treeNodeGraph[child] = childNode
		}

		isChild[child] = true

		if isLeft {
			parentNode.Left = childNode
		} else {
			parentNode.Right = childNode
		}
	}

	for k, v := range isChild {
		if !v {
			return treeNodeGraph[k]
		}
	}
	return nil
}
