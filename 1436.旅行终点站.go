package leetcode

/*
 * @lc app=leetcode.cn id=1436 lang=golang
 *
 * [1436] 旅行终点站
 */

// @lc code=start
func destCity(paths [][]string) string {
	outDegree := make(map[string]int)
	for _, path := range paths {
		outDegree[path[0]]++
		if _, ok := outDegree[path[1]]; !ok {
			outDegree[path[1]] = 0
		}
	}
	for city, degree := range outDegree {
		if degree == 0 {
			return city
		}
	}
	return ""
}

// @lc code=end
