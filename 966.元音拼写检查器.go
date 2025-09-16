package leetcode

import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=966 lang=golang
 *
 * [966] 元音拼写检查器
 */

// @lc code=start
func spellchecker(wordlist []string, queries []string) []string {
	vowels := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}

	formatWord := func(word string) string {
		wordBytes := make([]byte, 0, len(word))
		loweredWord := strings.ToLower(word)
		for i := range loweredWord {
			char := loweredWord[i]
			if vowels[char] {
				wordBytes = append(wordBytes, '*')
			} else {
				wordBytes = append(wordBytes, char)
			}
		}
		return string(wordBytes)
	}

	wordMap := make(map[string][]string, len(wordlist))
	for _, word := range wordlist {
		formattedWord := formatWord(word)
		wordMap[formattedWord] = append(wordMap[formattedWord], word)
	}

	result := make([]string, 0, len(queries))
	for _, q := range queries {
		formattedQuery := formatWord(q)
		matchedWord := ""
		for _, word := range wordMap[formattedQuery] {
			if word == q {
				matchedWord = word
				break
			}
		}

		if len(matchedWord) == 0 {
			for _, word := range wordMap[formattedQuery] {
				if strings.EqualFold(word, q) {
					matchedWord = word
					break
				}
			}
		}

		if len(matchedWord) == 0 && len(wordMap[formattedQuery]) > 0 {
			matchedWord = wordMap[formattedQuery][0]
		}

		result = append(result, matchedWord)
	}

	return result
}

// @lc code=end
