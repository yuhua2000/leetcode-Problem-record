package leetcode

/*
 * @lc app=leetcode.cn id=677 lang=golang
 *
 * [677] 键值映射
 */

// @lc code=start
type tree struct {
	children [26]*tree
	number   int
}

type MapSum struct {
	*tree
	cnt map[string]int
}

func Constructor() MapSum {
	return MapSum{&tree{}, make(map[string]int)}
}

func (m *MapSum) Insert(key string, val int) {
	delta := val
	if m.cnt[key] > 0 {
		delta -= m.cnt[key]
	}
	m.cnt[key] = val
	node := m.tree
	for i := 0; i < len(key); i++ {
		index := key[i] - 'a'
		if node.children[index] == nil {
			node.children[index] = &tree{}
		}
		node = node.children[index]
		node.number += delta
	}
}

func (m *MapSum) Sum(prefix string) int {
	node := m.tree
	for i := 0; i < len(prefix); i++ {
		index := prefix[i] - 'a'
		if node.children[index] == nil {
			return 0
		}
		node = node.children[index]
	}

	return node.number
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
// @lc code=end
