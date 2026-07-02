package leetcode

func findSafeWalk(grid [][]int, health int) bool {
	m, n := len(grid), len(grid[0])
	healthGrid := make([][]int, m)
	for i := range m {
		healthGrid[i] = make([]int, n)
	}

	healthGrid[0][0] = health - grid[0][0]

	dq := [][2]int{{0, 0}}
	dirs := [][2]int{{-1, 0}, {+1, 0}, {0, -1}, {0, +1}}

	for len(dq) > 0 {
		point := dq[0]
		dq = dq[1:]

		x, y := point[0], point[1]
		for _, dir := range dirs {
			nx, ny := dir[0]+x, dir[1]+y
			if nx >= 0 && nx < m && ny >= 0 && ny < n {
				nd := healthGrid[x][y] - grid[nx][ny]
				if nd > healthGrid[nx][ny] && nd >= 0 {
					healthGrid[nx][ny] = nd

					if grid[nx][ny] == 0 {
						dq = append([][2]int{{nx, ny}}, dq...)
					} else {
						dq = append(dq, [2]int{nx, ny})
					}
				}
			}
		}
	}

	return healthGrid[m-1][n-1] > 0
}
