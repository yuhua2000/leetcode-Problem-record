func mostPoints1(questions [][]int) int64 {
	var dfs func(i int) int64
	cache := map[int]int64{}
	dfs = func(i int) int64 {
		if i >= len(questions) {
			return 0
		}
		if ret, ok := cache[i]; ok {
			return ret
		}

		result := max(dfs(i+1), int64(questions[i][0])+dfs(i+questions[i][1]+1))
		cache[i] = result
		return result
	}

	return dfs(0)
}

func mostPoints(questions [][]int) int64 {
	n := len(questions)
	dp := make([]int64, n+1)

	for i := n - 1; i >= 0; i-- {
		dp[i] = max(dp[i+1], int64(questions[i][0])+dp[min(n, i+1+questions[i][1])])
	}

	return dp[0]
}
