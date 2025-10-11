package leetcode

/*
 * @lc app=leetcode.cn id=1043 lang=golang
 *
 * [1043] 分隔数组以得到最大和
 */

// @lc code=start

func maxSumAfterPartitioning(arr []int, k int) int {
	n := len(arr)
	d := make([]int, n+1)
	for i := 1; i <= n; i++ {
		j := i - 1
		maxValue := arr[j]
		for ; j >= max(0, i-k); j-- {
			if arr[j] > maxValue {
				maxValue = arr[j]
			}
			d[i] = max(d[i], d[j]+(maxValue*(i-j)))
		}
	}
	return d[n]
}

// @lc code=end
