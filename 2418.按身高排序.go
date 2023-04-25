/*
 * @lc app=leetcode.cn id=2418 lang=golang
 *
 * [2418] 按身高排序
 */

// @lc code=start
func sortPeople(names []string, heights []int) []string {
	n := len(names)
	indices := make([]int, n)
	for i := 0; i < n; i++ {
		indices[i] = i
	}
	sort.Slice(indices, func(i, j int) bool {
		return heights[indices[j]] < heights[indices[i]]
	})
	res := make([]string, 0, n)
	for i := 0; i < n; i++ {
		res = append(res, names[indices[i]])
	}
	return res
}

// @lc code=end
