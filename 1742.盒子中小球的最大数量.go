/*
 * @lc app=leetcode.cn id=1742 lang=golang
 *
 * [1742] 盒子中小球的最大数量
 */

// @lc code=start
func countBalls(lowLimit int, highLimit int) int {
	var box = make(map[int]int)
	for i := lowLimit; i <= highLimit; i++ {
		ball := i
		var index int
		for ball > 0 {
			index += ball % 10
			ball /= 10
		}
		box[index]++
	}

	var result int
	for v := range maps.Values(box) {
		result = max(result, v)
	}

	return result
}

// @lc code=end
