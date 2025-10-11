package leetcode

/*
 * @lc app=leetcode.cn id=2103 lang=golang
 *
 * [2103] 环和杆
 */

// @lc code=start
func countPoints(rings string) int {
	if len(rings) < 6 {
		return 0
	}
	const (
		R int8 = 1 << iota
		G
		B
	)
	var poles = [10]int8{}
	for i := 0; i < len(rings); i += 2 {
		var colour = R
		switch rings[i] {
		case 'G':
			colour = G
		case 'B':
			colour = B
		}
		poles[rings[i+1]-'0'] |= colour
	}
	result := 0
	for _, pole := range poles {
		if pole == 7 {
			result++
		}
	}
	return result
}

// @lc code=end
