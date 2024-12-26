package leetcode

import "math/bits"

func numberOfSpecialChars1(word string) int {
	var charFlags [26]int8
	for i := 0; i < len(word); i++ {
		if word[i] >= 'a' && word[i] <= 'z' {
			charFlags[word[i]-'a'] |= 1
		} else {
			charFlags[word[i]-'A'] |= 2
		}
	}

	result := 0
	for i := 0; i < len(charFlags); i++ {
		if charFlags[i] == 3 {
			result++
		}
	}
	return result
}

func numberOfSpecialChars(word string) int {
	var charFlags [2]int
	for _, c := range word {
		charFlags[c>>5&1] |= 1 << (c & 31)
	}

	return bits.OnesCount(uint(charFlags[0] & charFlags[1]))
}
