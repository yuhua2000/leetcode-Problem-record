package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=2363 lang=golang
 *
 * [2363] 合并相似的物品
 */

// @lc code=start
func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
	mp := make(map[int]int, len(items1))
	for _, a := range items1 {
		mp[a[0]] += a[1]
	}
	for _, a := range items2 {
		mp[a[0]] += a[1]
	}
	var ans = make([][]int, 0, len(mp))
	for a, b := range mp {
		ans = append(ans, []int{a, b})
	}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i][0] < ans[j][0]
	})
	return ans
}

// @lc code=end
