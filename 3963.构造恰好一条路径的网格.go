package leetcode

func createGrid(m int, n int) []string {
	ans := make([][]byte, m)
	for i := 0; i < m; i++ {
		ans[i] = make([]byte, n)
		ans[i][0] = '.'
	}

	for i := 0; i < n; i++ {
		ans[m-1][i] = '.'
	}

	var result []string

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if ans[i][j] == 0 {
				ans[i][j] = '#'
			}
		}

		result = append(result, string(ans[i]))
	}

	return result

}
