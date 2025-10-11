package leetcode

/*
 * @lc app=leetcode.cn id=1922 lang=golang
 *
 * [1922] 统计好数字的数目
 */

// @lc code=start
const mod = 1e9 + 7

func countGoodNumbers(n int64) int {
	return (pow(5, (n+1)/2) * pow(4, n/2)) % mod
}

func pow(x, y int64) int {
	var res int64 = 1
	for y > 0 {
		if y%2 == 1 {
			res = res * x % mod
		}
		x = x * x % mod
		y /= 2
	}
	return int(res)
}

// @lc code=end
