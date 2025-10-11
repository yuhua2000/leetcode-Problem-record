package leetcode

import "reflect"

/*
 * @lc app=leetcode.cn id=2451 lang=golang
 *
 * [2451] 差值数组不同的字符串
 */

// @lc code=start
func oddString(words []string) string {
	diff0 := get(words[0])
	diff1 := get(words[1])
	if reflect.DeepEqual(diff0, diff1) {
		for i := 2; i < len(words); i++ {
			if !reflect.DeepEqual(diff0, get(words[i])) {
				return words[i]
			}
		}
	}
	if reflect.DeepEqual(diff0, get(words[2])) {
		return words[1]
	}
	return words[0]
}

func get(word string) []int {
	diff := make([]int, len(word)-1)
	for i := 0; i < len(word)-1; i++ {
		diff[i] = int(word[i+1] - word[i])
	}
	return diff
}

// @lc code=end
