
/*
 * @lc app=leetcode.cn id=1096 lang=golang
 *
 * [1096] 花括号展开 II
 */

// @lc code=start
func braceExpansionII(expression string) []string {
	s := make(map[string]bool)
	var dnf func(string)
	dnf = func(exp string) {
		if j := strings.Index(exp, "}"); j < 0 {
			s[exp] = true
			return
		} else {
			i := strings.LastIndex(exp[:j], "{")
			a, c := exp[:i], exp[j+1:]
			for _, b := range strings.Split(exp[i+1:j], ",") {
				dnf(a + b + c)
			}
		}
	}
	dnf(expression)
	ans := make([]string, 0, len(s))
	for k := range s {
		ans = append(ans, k)
	}
	sort.Strings(ans)
	return ans
}

// @lc code=end
