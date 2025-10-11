package leetcode

import "fmt"

/*
 * @lc app=leetcode.cn id=2352 lang=golang
 *
 * [2352] 相等行列对
 */

// @lc code=start
func equalPairs(grid [][]int) int {
	n := len(grid)
	cnt := make(map[string]int)
	for i := 0; i < n; i++ {
		cnt[fmt.Sprint(grid[i])]++
	}
	res := 0
	for i := 0; i < n; i++ {
		var arr = make([]int, 0, n)
		for j := 0; j < n; j++ {
			arr = append(arr, grid[j][i])
		}
		res += cnt[fmt.Sprint(arr)]
	}
	return res
}

// @lc code=end
