package leetcode

/*
 * @lc app=leetcode.cn id=1935 lang=golang
 *
 * [1935] 可以输入的最大单词数
 */

// @lc code=start
func canBeTypedWords(text string, brokenLetters string) int {
	brokenLettersMap := make(map[byte]bool, len(brokenLetters))
	for i := 0; i < len(brokenLetters); i++ {
		brokenLettersMap[brokenLetters[i]] = true
	}

	result := 0
	isValidWord := true
	for i := range text {
		char := text[i]
		if char == ' ' {
			if isValidWord == true {
				result++
			}
			isValidWord = true
		} else if brokenLettersMap[char] {
			isValidWord = false

		}
	}

	if isValidWord {
		result++
	}
	return result
}

// @lc code=end
