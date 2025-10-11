package leetcode

/*
 * @lc app=leetcode.cn id=1156 lang=golang
 *
 * [1156] 单字符重复子串的最大长度
 */

// @lc code=start
func maxRepOpt1(text string) int {
	n := len(text)
	count := make(map[byte]int)
	for i := range text {
		count[text[i]]++
	}
	res := 0
	for i := 0; i < n; {
		j := i
		for j < n && text[j] == text[i] {
			j++
		}
		curCnt := j - i
		if curCnt < count[text[i]] && (i > 0 || j < n) {
			res = max(res, curCnt+1)
		}
		k := j + 1
		for k < n && text[k] == text[i] {
			k++
		}
		res = max(res, min(k-i, count[text[i]]))
		i = j
	}
	return res
}

// @lc code=end
