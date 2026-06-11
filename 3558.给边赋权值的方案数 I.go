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

func assignEdgeWeights(edges [][]int) int {
	graph := make(map[int][]int)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	q := []int{1}
	visited := make(map[int]bool)
	visited[1] = true
	depth := -1
	for len(q) > 0 {
		depth++
		size := len(q)
		for _, node := range q[:size] {
			for _, nextNode := range graph[node] {
				if !visited[nextNode] {
					visited[nextNode] = true
					q = append(q, nextNode)
				}
			}
		}
		q = q[size:]
	}

	return int(modPow(2, int64(depth-1)))

}
