/*
 * @lc app=leetcode.cn id=670 lang=golang
 *
 * [670] 最大交换
 */

// @lc code=start
func maximumSwap(num int) int {
	s := []byte(strconv.Itoa(num))
	maxIdx, idx1, idx2 := len(s)-1, -1, -1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] > s[maxIdx] {
			maxIdx = i
		} else if s[i] < s[maxIdx] {
			idx1, idx2 = i, maxIdx
		}
	}
	if idx1 < 0 {
		return num
	}
	s[idx1], s[idx2] = s[idx2], s[idx1]
	v, _ := strconv.Atoi(string(s))
	return v
}

// @lc code=end
