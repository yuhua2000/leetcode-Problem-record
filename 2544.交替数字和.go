/*
 * @lc app=leetcode.cn id=2544 lang=golang
 *
 * [2544] 交替数字和
 */

// @lc code=start

func alternateDigitSum1(n int) int {
	sum := 0
	sign := 1
	ans := true
	for i := (int)(math.Pow(10, 9)); i > 0; i /= 10 {
		if i > n && ans {
			continue
		}
		ans = false
		num := n / i
		sum += sign * num
		sign = -sign
		n -= i * num
	}
	return sum
}
func alternateDigitSum(n int) int {
	res, sign := 0, 1
	for n > 0 {
		res += sign * n % 10
		n /= 10
		sign = -sign
	}
	return -sign * res
}

// @lc code=end
