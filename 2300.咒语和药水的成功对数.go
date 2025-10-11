package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=2300 lang=golang
 *
 * [2300] 咒语和药水的成功对数
 */

// @lc code=start
func successfulPairs(spells []int, potions []int, success int64) (result []int) {
	sort.Ints(potions)
	n := len(potions)
	for _, spell := range spells {
		index := int(success) / spell
		if int(success)%spell != 0 {
			index++
		}
		result = append(result, n-sort.SearchInts(potions, index))
	}
	return result
}

// @lc code=end
