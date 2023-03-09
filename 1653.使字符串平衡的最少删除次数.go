
/*
 * @lc app=leetcode.cn id=1653 lang=golang
 *
 * [1653] 使字符串平衡的最少删除次数
 */

// @lc code=start
func minimumDeletions1(s string) int {
	n := len(s)
	ans := make([][2]int, n)
	a, b := 0, 0
	for i, c := range s {
		if c == 'b' {
			b++
		}
		ans[i][1] = b
		if s[n-1-i] == 'a' {
			a++
		}
		ans[n-1-i][0] = a
	}
	min := math.MaxInt
	for _, a := range ans {
		if a[0]+a[1]-1 < min {
			min = a[0] + a[1] - 1
		}
	}
	return min
}

func minimumDeletions(s string) int {
	n := len(s)
	ans := make([][2]int, n)
	a, b := 0, 0
	for i, c := range s {
		if c == 'a' {
			a++
		}
	}
	min = a
	for _, c := range s {
		if c == 'a' {
			a--
		} else {
			b++
		}
		if a+b < min {
			min = a + b
		}
	}

	return min
}

// @lc code=end
