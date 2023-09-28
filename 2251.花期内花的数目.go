/*
 * @lc app=leetcode.cn id=2251 lang=golang
 *
 * [2251] 花期内花的数目
 */

// @lc code=start
func fullBloomFlowers(flowers [][]int, people []int) []int {
	var cnt = make(map[int]int)
	for _, f := range flowers {
		cnt[f[0]]++
		cnt[f[1]+1]--
	}

	var arr = make([][]int, 0, len(cnt))
	for key, val := range cnt {
		arr = append(arr, []int{key, val})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})

	var indices = make([][]int, len(people))
	for i, p := range people {
		indices[i] = []int{p, i}
	}
	sort.Slice(indices, func(i, j int) bool {
		return indices[i][0] < indices[j][0]
	})

	j, curr := 0, 0
	var result = make([]int, len(people))
	for _, p := range indices {
		for j < len(arr) && arr[j][0] <= p[0] {
			curr += arr[j][1]
			j++
		}
		result[p[1]] = curr
	}

	return result
}

// @lc code=end
