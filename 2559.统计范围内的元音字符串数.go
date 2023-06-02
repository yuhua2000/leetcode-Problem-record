/*
 * @lc app=leetcode.cn id=2559 lang=golang
 *
 * [2559] 统计范围内的元音字符串数
 */

// @lc code=start
func vowelStrings(words []string, queries [][]int) []int {
	ans := make([]int, len(words)+1)
	sum := 0
	for i, str := range words {
		if check(str) {
			sum++
		}
		ans[i+1] = sum
	}
	result := make([]int, len(queries))
	for i, q := range queries {
		result[i] = ans[q[1]+1] - ans[q[0]]
	}
	return result
}

func check(str string) bool {
	//aeiou
	start, end := str[0], str[len(str)-1]
	if start == 'a' || start == 'e' || start == 'i' || start == 'o' || start == 'u' {
		if end == 'a' || end == 'e' || end == 'i' || end == 'o' || end == 'u' {
			return true
		}
	}
	return false
}

// @lc code=end
