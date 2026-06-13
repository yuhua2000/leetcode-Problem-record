package leetcode

const MOD int64 = 1_000_000_007

func modPow(a, b int64) int64 {
	res := int64(1)
	for b > 0 {
		if b&1 == 1 {
			res = res * a % MOD
		}
		a = a * a % MOD
		b >>= 1
	}
	return res
}

func assignEdgeWeights1(edges [][]int, queries [][]int) []int {
	n := len(edges) + 1

	graph := make([][]int, n+1)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	const LOG = 18

	up := make([][LOG]int, n+1)
	q := []int{1}
	depth := make([]int, n+1)
	depth[1] = 1
	for len(q) > 0 {
		size := len(q)
		for _, node := range q[:size] {
			for _, nextNode := range graph[node] {
				if depth[nextNode] == 0 {
					up[nextNode][0] = node
					depth[nextNode] = depth[node] + 1
					q = append(q, nextNode)
				}
			}
		}
		q = q[size:]
	}

	for i := 1; i < LOG; i++ {
		for u := 1; u <= n; u++ {
			up[u][i] = up[up[u][i-1]][i-1]
		}
	}

	result := make([]int, 0, len(queries))
	for _, querie := range queries {
		u, v := querie[0], querie[1]

		if v == u {
			result = append(result, 0)
			continue
		}

		var distance int

		if depth[u] > depth[v] {
			diff := depth[u] - depth[v]
			for k := LOG - 1; k >= 0; k-- {
				if diff >= 1<<k && up[u][k] != 0 {
					distance += (1 << k)
					u = up[u][k]
					diff = depth[u] - depth[v]
				}
			}
		}

		if depth[v] > depth[u] {
			diff := depth[v] - depth[u]
			for k := LOG - 1; k >= 0; k-- {
				if diff >= 1<<k && up[v][k] != 0 {
					distance += (1 << k)
					v = up[v][k]
					diff = depth[v] - depth[u]
				}
			}
		}

		for k := LOG - 1; k >= 0 && u != v; k-- {
			if up[u][k] != up[v][k] {
				u = up[u][k]
				v = up[v][k]
				distance += 2 * (1 << k)
			}
		}

		if u != v {
			distance += 2
		}

		result = append(result, int(modPow(2, int64(distance-1))))
	}

	return result

}
func assignEdgeWeights(edges [][]int, queries [][]int) []int {
	n := len(edges) + 1

	graph := make([][]int, n+1)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	const LOG = 18

	up := make([][LOG]int, n+1)
	q := []int{1}
	depth := make([]int, n+1)
	depth[1] = 1
	for len(q) > 0 {
		size := len(q)
		for _, node := range q[:size] {
			for _, nextNode := range graph[node] {
				if depth[nextNode] == 0 {
					up[nextNode][0] = node
					depth[nextNode] = depth[node] + 1
					q = append(q, nextNode)
				}
			}
		}
		q = q[size:]
	}

	for i := 1; i < LOG; i++ {
		for u := 1; u <= n; u++ {
			up[u][i] = up[up[u][i-1]][i-1]
		}
	}

	result := make([]int, 0, len(queries))

	var lca func(u, v int) int

	lca = func(u, v int) int {
		if depth[u] > depth[v] {
			diff := depth[u] - depth[v]
			for k := LOG - 1; k >= 0; k-- {
				if diff&(1<<k) != 0 {
					u = up[u][k]
				}
			}
		}

		if depth[v] > depth[u] {
			diff := depth[v] - depth[u]
			for k := LOG - 1; k >= 0; k-- {
				if diff&(1<<k) != 0 {
					v = up[v][k]
				}
			}
		}

		if u == v {
			return u
		}

		for k := LOG - 1; k >= 0; k-- {
			if up[u][k] != up[v][k] {
				u = up[u][k]
				v = up[v][k]
			}
		}

		return up[u][0]
	}

	for _, querie := range queries {
		u, v := querie[0], querie[1]

		if v == u {
			result = append(result, 0)
			continue
		}

		distance := depth[u] + depth[v] - 2*depth[lca(u, v)]

		result = append(result, int(modPow(2, int64(distance-1))))
	}

	return result
}
