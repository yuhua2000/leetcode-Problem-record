package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=1090 lang=golang
 *
 * [1090] 受标签影响的最大值
 */

// @lc code=start
func largestValsFromLabels(values []int, labels []int, numWanted int, useLimit int) int {
	n := len(values)
	idx := make([]int, n)
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(i, j int) bool {
		return values[idx[i]] > values[idx[j]]
	})
	ans, choose := 0, 0
	cnt := make(map[int]int)
	for i := range idx {
		label := labels[idx[i]]
		if cnt[label] >= useLimit {
			continue
		}
		choose++
		ans += values[idx[i]]
		cnt[label]++
		if choose == numWanted {
			break
		}
	}
	return ans
}

// @lc code=end
