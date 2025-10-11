package leetcode

import (
	"sort"
	"strings"
)

/*
 * @lc app=leetcode.cn id=833 lang=golang
 *
 * [833] 字符串中的查找与替换
 */

// @lc code=start
func findReplaceString(s string, indices []int, sources []string, targets []string) string {
	n, m := len(s), len(indices)
	ops := make([]int, len(indices))
	for i := range ops {
		ops[i] = i
	}
	sort.Slice(ops, func(i, j int) bool {
		return indices[ops[i]] < indices[ops[j]]
	})

	ans := strings.Builder{}
	pt := 0
	for i := 0; i < n; {
		for pt < m && indices[ops[pt]] < i {
			pt++
		}
		succeed := false
		for pt < m && indices[ops[pt]] == i {
			if s[i:i+min(len(sources[ops[pt]]), n-i)] == sources[ops[pt]] {
				succeed = true
				break
			}
			pt++
		}
		if succeed {
			ans.WriteString(targets[ops[pt]])
			i += len(sources[ops[pt]])
		} else {
			ans.WriteByte(s[i])
			i++
		}
	}
	return ans.String()

}

// @lc code=end
