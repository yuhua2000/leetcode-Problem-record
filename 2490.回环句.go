/*
 * @lc app=leetcode.cn id=2490 lang=golang
 *
 * [2490] 回环句
 */

// @lc code=start

func isCircularSentence(sentence string) bool {
	for i := 0; i < len(sentence); i++ {
		if sentence[i] == ' ' {
			if sentence[i-1] != sentence[i+1] {
				return false
			}
		}
	}
	return sentence[0] == sentence[len(sentence)-1]
}

// @lc code=end
