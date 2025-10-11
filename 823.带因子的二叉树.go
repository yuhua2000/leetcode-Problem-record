package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=823 lang=golang
 *
 * [823] 带因子的二叉树
 */

// @lc code=start
func numFactoredBinaryTrees(arr []int) int {
	sort.Ints(arr)
	result, mod := int(0), int(1e9+7)
	dp := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		dp[i] = 1
		for left, right := 0, i-1; left <= right; left++ {
			for left <= right && arr[left]*arr[right] > arr[i] {
				right--
			}
			if left <= right && arr[left]*arr[right] == arr[i] {
				if left == right {
					dp[i] = (dp[i] + dp[left]*dp[right]) % mod
				} else {
					dp[i] = (dp[i] + dp[left]*dp[right]*2) % mod
				}
			}
		}
		result = (result + dp[i]) % mod
	}
	return result
}

// @lc code=end
