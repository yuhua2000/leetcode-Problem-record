/*
 * @lc app=leetcode.cn id=1761 lang=golang
 *
 * [1761] 一个图中连通三元组的最小度数
 */

// @lc code=start
func minTrioDegree1(n int, edges [][]int) int {
	g := make([][]int, n)
	degree := make([]int, n)
	for i := 0; i < n; i++ {
		g[i] = make([]int, n)
	}
	for _, edge := range edges {
		x, y := edge[0]-1, edge[1]-1
		g[x][y] = 1
		g[y][x] = 1
		degree[x]++
		degree[y]++
	}
	ans := 0x3f3f3f3f
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if g[i][j] != 1 {
				continue
			}
			for k := j + 1; k < n; k++ {
				if g[i][k] == 1 && g[j][k] == 1 {
					ans = min(degree[i]+degree[j]+degree[k]-6, ans)
				}
			}
		}
	}
	if ans == 0x3f3f3f3f {
		return -1
	}
	return ans
}

func minTrioDegree(n int, edges [][]int) int {
	g := make([][]int, n)
	h := make([][]int, n)
	degree := make([]int, n)
	for i := 0; i < n; i++ {
		g[i] = make([]int, n)
	}
	for _, edge := range edges {
		x, y := edge[0]-1, edge[1]-1
		g[x][y] = 1
		g[y][x] = 1
		degree[x]++
		degree[y]++

		if x < y {
			h[x] = append(h[x], y)
		} else {
			h[y] = append(h[y], x)
		}
	}

	// for _, edge := range edges {
	// 	x, y := edge[0]-1, edge[1]-1
	// 	if degree[x] < degree[y] || (degree[x] == degree[y] && x < y) {
	// 		h[x] = append(h[x], y)
	// 	} else {
	// 		h[y] = append(h[y], x)
	// 	}
	// }

	ans := 0x3f3f3f3f
	for i := 0; i < n; i++ {
		for _, j := range h[i] {
			for _, k := range h[j] {
				if g[i][k] == 1 {
					ans = min(degree[i]+degree[j]+degree[k]-6, ans)
				}
			}
		}
	}
	if ans == 0x3f3f3f3f {
		return -1
	}
	return ans
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end
