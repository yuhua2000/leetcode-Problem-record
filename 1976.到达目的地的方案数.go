package leetcode

import "math"

/*
 * @lc app=leetcode.cn id=1976 lang=golang
 *
 * [1976] 到达目的地的方案数
 */

// @lc code=start
const mod = 1000000007

func countPaths(n int, roads [][]int) int {
	dist := make([][]int64, n)
	for i := range dist {
		dist[i] = make([]int64, n)
		for j := range dist[i] {
			if i == j {
				dist[i][j] = 0
			} else {
				dist[i][j] = math.MaxInt64 / 2
			}
		}
	}
	for _, road := range roads {
		x, y, z := road[0], road[1], int64(road[2])
		dist[x][y] = z
		dist[y][x] = z
	}

	// Floyd 算法求解最短路
	// 完成后，dist[0][i] 即为正文部分的 dist[i]
	// for k := 0; k < n; k++ {
	//     for i := 0; i < n; i++ {
	//         for j := 0; j < n; j++ {
	//             if dist[i][k]+dist[k][j] < dist[i][j] {
	//                 dist[i][j] = dist[i][k] + dist[k][j]
	//             }
	//         }
	//     }
	// }

	// Dijkstra 算法求解最短路
	// 完成后，dist[0][i] 即为正文部分的 dist[i]
	const inf = math.MaxInt64 / 2
	used := make([]bool, n)

	for i := 0; i < n; i++ {
		u := -1
		for j := 0; j < n; j++ {
			if !used[j] && (u == -1 || dist[0][j] < dist[0][u]) {
				u = j
			}
		}
		used[u] = true
		for v := 0; v < n; v++ {
			if dist[0][u]+dist[u][v] < dist[0][v] {
				dist[0][v] = dist[0][u] + dist[u][v]
			}
		}
	}

	// 构造图 G
	g := make([][]int, n)
	for _, road := range roads {
		x, y, z := road[0], road[1], int64(road[2])
		if dist[0][y]-dist[0][x] == z {
			g[x] = append(g[x], y)
		} else if dist[0][x]-dist[0][y] == z {
			g[y] = append(g[y], x)
		}
	}

	f := make([]int, n)
	for i := range f {
		f[i] = -1
	}
	var dfs func(int) int
	dfs = func(u int) int {
		if u == n-1 {
			return 1
		}
		if f[u] != -1 {
			return f[u]
		}

		f[u] = 0
		for _, v := range g[u] {
			f[u] += dfs(v)
			if f[u] >= mod {
				f[u] -= mod
			}
		}
		return f[u]
	}
	return dfs(0)
}

// @lc code=end
