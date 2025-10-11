package leetcode

/*
 * @lc app=leetcode.cn id=211 lang=golang
 *
 * [211] 添加与搜索单词 - 数据结构设计
 */

// @lc code=start

type tree struct {
	children [26]*tree
	isEnd    bool
}

func (t *tree) insert(w string) {
	node := t
	for i := 0; i < len(w); i++ {
		index := w[i] - 'a'
		if node.children[index] == nil {
			node.children[index] = &tree{}
		}
		node = node.children[index]
	}
	node.isEnd = true
}

type WordDictionary struct {
	*tree
}

func Constructor() WordDictionary {
	return WordDictionary{
		tree: &tree{},
	}
}

func (w *WordDictionary) AddWord(word string) {
	w.insert(word)
}

func (w *WordDictionary) Search(word string) bool {
	var dfs func(int, *tree) bool
	dfs = func(index int, node *tree) bool {
		if index == len(word) {
			return node.isEnd
		}
		var ch = word[index]
		if ch != '.' {
			child := node.children[ch-'a']
			if child != nil && dfs(index+1, child) {
				return true
			}
		} else {
			for i := 0; i < len(node.children); i++ {
				child := node.children[i]
				if child != nil && dfs(index+1, child) {
					return true
				}
			}
		}
		return false
	}
	return dfs(0, w.tree)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
// @lc code=end
