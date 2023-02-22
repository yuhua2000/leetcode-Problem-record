/*
 * @lc app=leetcode.cn id=2347 lang=golang
 *
 * [2347] 最好的扑克手牌
 */

// @lc code=start
func bestHand(ranks []int, suits []byte) string {
	count := len(ranks)
	rankDic := make(map[int]int, count)
	suitDic := make(map[byte]int, count)
	maxNum := 0
	for i := 0; i < count; i++ {
		rankDic[ranks[i]]++
		suitDic[suits[i]]++
		if rankDic[ranks[i]] > maxNum {
			maxNum = rankDic[ranks[i]]
		}
	}
	if len(suitDic) == 1 {
		return "Flush"
	}

	if maxNum >= 3 {
		return "Three of a Kind"
	}
	if maxNum == 2 {
		return "Pair"
	}
	return "High Card"
}

// "Flush"：同花，五张相同花色的扑克牌。
// "Three of a Kind"：三条，有 3 张大小相同的扑克牌。
// "Pair"：对子，两张大小一样的扑克牌。
// "High Card"：高牌，五张大小互不相同的扑克牌。
// @lc code=end
