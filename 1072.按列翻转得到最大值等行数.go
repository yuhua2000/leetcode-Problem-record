package leetcode

/*
 * @lc app=leetcode.cn id=1072 lang=golang
 *
 * [1072] 按列翻转得到最大值等行数
 */

// @lc code=start
func maxEqualRowsAfterFlips(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	mp := make(map[string]int)
	for i := 0; i < m; i++ {
		arr := make([]byte, n)
		for j := 0; j < n; j++ {
			if matrix[i][j]^matrix[i][0] == 0 {
				arr[j] = '0'
			} else {
				arr[j] = '1'
			}
		}
		mp[string(arr)]++
	}
	result := 0
	for _, v := range mp {
		if v > result {
			result = v
		}
	}
	return result
}

// @lc code=end
