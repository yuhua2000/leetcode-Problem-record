/*
 * @lc app=leetcode.cn id=1092 lang=golang
 *
 * [1092] 最短公共超序列
 */

// @lc code=start
func shortestCommonSupersequence(str1 string, str2 string) string {
	n, m := len(str1), len(str2)
	f := make([][]int, n+1)
	for i := range f {
		f[i] = make([]int, m+1)
	}
	for j := 0; j < m; j++ {
		f[0][j] = j
	}
	for i := 0; i < n; i++ {
		f[i][0] = i
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if str1[i] == str2[j] {
				f[i+1][j+1] = f[i][j] + 1
			} else {
				f[i+1][j+1] = min(f[i][j+1], f[i+1][j]) + 1
			}
		}
	}
	ans := []byte{}
	i, j := n-1, m-1
	for i >= 0 && j >= 0 {
		if str1[i] == str2[j] {
			ans = append(ans, str1[i])
			j--
			i--
		} else if f[i+1][j+1] == f[i][j+1]+1 {
			ans = append(ans, str1[i])
			i--
		} else {
			ans = append(ans, str2[j])
			j--
		}
	}
	for i, n := 0, len(ans); i < n/2; i++ {
		ans[i], ans[n-i-1] = ans[n-i-1], ans[i]
	}
	return str1[:i+1] + str2[:j+1] + string(ans)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end
