package leetcode

import (
	"bytes"
	"slices"
)

/*
 * @lc app=leetcode.cn id=2273 lang=golang
 *
 * [2273] 移除字母异位词后的结果数组
 */

// @lc code=start
func removeAnagrams(words []string) []string {
	prev := []byte(words[0])
	slices.Sort(prev)
	result := []string{words[0]}
	for i := 1; i < len(words); i++ {
		word := []byte(words[i])
		slices.Sort(word)
		if bytes.Equal(prev, word) {
			continue
		}
		prev = word
		result = append(result, words[i])

	}

	return result
}

// @lc code=end
