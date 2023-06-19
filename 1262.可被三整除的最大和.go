/*
 * @lc app=leetcode.cn id=1262 lang=golang
 *
 * [1262] 可被三整除的最大和
 */

// @lc code=start
func accumulate(v []int) int {
	ans := 0
	for _, x := range v {
		ans += x
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func maxSumDivThree(nums []int) int {
	v := make([][]int, 3)
	for _, num := range nums {
		v[num%3] = append(v[num%3], num)
	}
	sort.Slice(v[1], func(i, j int) bool {
		return v[1][i] > v[1][j]
	})
	sort.Slice(v[2], func(i, j int) bool {
		return v[2][i] > v[2][j]
	})
	ans, lb, lc := 0, len(v[1]), len(v[2])
	for i := max(lb-2, 0); i <= lb; i++ {
		for j := max(lc-2, 0); j <= lc; j++ {
			if (i-j)%3 == 0 {
				ans = max(ans, accumulate(v[1][:i])+accumulate(v[2][:j]))
			}
		}
	}
	return ans + accumulate(v[0])
}

// @lc code=end
