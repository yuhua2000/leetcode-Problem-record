package leetcode

import (
	"sort"
	"strings"
)

/*
 * @lc app=leetcode.cn id=1233 lang=golang
 *
 * [1233] 删除子文件夹
 */

// @lc code=start
type Tree struct {
	children map[string]*Tree
	isEnd    bool
}

func NewTree() *Tree {
	return &Tree{
		children: map[string]*Tree{},
	}
}
func (t *Tree) insert(w string) {
	node := t
	for _, p := range strings.Split(w, "/")[1:] {
		if _, ok := node.children[p]; !ok {
			node.children[p] = NewTree()
		}
		node, _ = node.children[p]
	}
	node.isEnd = true
}

func (t *Tree) search(w string) bool {
	node := t
	for _, p := range strings.Split(w, "/")[1:] {
		if _, ok := node.children[p]; !ok {
			return false
		}
		node, _ = node.children[p]
		if node.isEnd {
			return true
		}
	}
	return false
}

func removeSubfolders(folder []string) []string {
	sort.Strings(folder)
	trie := NewTree()
	var ans []string
	for _, v := range folder {
		if !trie.search(v) {
			trie.insert(v)
			ans = append(ans, v)
		}
	}
	return ans

}

// @lc code=end
