package leetcode

/*
 * @lc app=leetcode.cn id=771 lang=golang
 *
 * [771] 宝石与石头
 */

// @lc code=start
func numJewelsInStones(jewels string, stones string) (result int) {
	jewelMap := make(map[byte]bool, len(jewels))
	for i := range jewels {
		jewelMap[jewels[i]] = true
	}

	for i := range stones {
		if jewelMap[stones[i]] {
			result++
		}
	}

	return
}

// @lc code=end
