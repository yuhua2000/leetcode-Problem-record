package leetcode

func canReach(arr []int, start int) bool {
	visit := make([]bool, len(arr))
	var dfs func(int) bool
	dfs = func(pos int) bool {
		if pos < 0 || pos >= len(arr) {
			return false
		}

		if arr[pos] == 0 {
			return true
		}

		if visit[pos] {
			return false
		}
		visit[pos] = true

		return dfs(pos-arr[pos]) || dfs(pos+arr[pos])
	}

	return dfs(start)
}
