package leetcode

/*
 * @lc app=leetcode.cn id=1946 lang=golang
 *
 * [1946] 子字符串突变后可能得到的最大整数
 */

// @lc code=start
func maximumNumber(num string, change []int) string {
	ans := []byte(num)
	for i := 0; i < len(ans); i++ {
		if int(ans[i]-'0') < change[ans[i]-'0'] {
			for j := i; j < len(ans) && int(ans[j]-'0') <= change[ans[j]-'0']; j++ {
				ans[j] = byte(change[ans[j]-'0']) + '0'
			}
			break
		}
	}
	return string(ans)
}

// @lc code=end
