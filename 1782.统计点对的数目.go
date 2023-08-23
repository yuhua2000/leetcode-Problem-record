/*
 * @lc app=leetcode.cn id=1782 lang=golang
 *
 * [1782] 统计点对的数目
 */

// @lc code=start
func countPairs(n int, edges [][]int, queries []int) []int {
	degree := make([]int, n)
	cnt := make(map[int]int)
	for _, edge := range edges {
		x, y := edge[0]-1, edge[1]-1
		if x > y {
			x, y = y, x
		}
		degree[x]++
		degree[y]++
		cnt[x*n+y]++
	}
	arr := make([]int, n)
	copy(arr, degree)
	sort.Ints(arr)
	ans := make([]int, 0, len(queries))
	for _, bound := range queries {
		total := 0
		for i := 0; i < n; i++ {
			j := sort.SearchInts(arr, bound-arr[i]+1)
			total += n - max(i+1, j)
		}
		for val, freq := range cnt {
			x, y := val/n, val%n
			if degree[x]+degree[y] > bound && degree[x]+degree[y]-freq <= bound {
				total--
			}
		}
		ans = append(ans, total)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
