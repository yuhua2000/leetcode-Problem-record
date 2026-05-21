package leetcode

import (
	"strconv"
)

type TrieNode struct {
	Children [10]*TrieNode
}

func longestCommonPrefix(arr1 []int, arr2 []int) int {
	root1 := &TrieNode{}
	root2 := &TrieNode{}

	var buildTrie func(node *TrieNode, num string, index int)
	buildTrie = func(node *TrieNode, num string, index int) {
		if index >= len(num) {
			return
		}

		digit := int(num[index] - '0')
		if node.Children[digit] == nil {
			node.Children[digit] = &TrieNode{}
		}

		buildTrie(node.Children[digit], num, index+1)
	}

	for _, num := range arr1 {
		buildTrie(root1, strconv.Itoa(num), 0)
	}

	for _, num := range arr2 {
		buildTrie(root2, strconv.Itoa(num), 0)
	}

	var calcMaxPrefix func(node1, node2 *TrieNode) int
	calcMaxPrefix = func(node1, node2 *TrieNode) int {
		if node1 == nil || node2 == nil {
			return 0
		}
		maxLen := 0
		for i := 0; i < 10; i++ {
			maxLen = max(calcMaxPrefix(node1.Children[i], node2.Children[i]), maxLen)
		}
		return maxLen + 1
	}

	return calcMaxPrefix(root1, root2) - 1
}
