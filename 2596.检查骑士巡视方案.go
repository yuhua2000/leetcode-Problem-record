func checkValidGrid(grid [][]int) bool {
	if grid[0][0] != 0 {
		return false
	}
	n := len(grid)
	step := make([]int, n*n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			step[grid[i][j]] = i*n + j
		}
	}
	for i := 0; i < n*n-1; i++ {
		if !legal(step[i], step[i+1], n) {
			return false
		}
	}
	return true
}

func legal(x, y, n int) bool {
	x1, y1 := x/n, x%n
	x2, y2 := y/n, y%n

	diffX := x1 - x2
	if x < 0 {
		diffX = -diffX
	}
	diffY := y1 - y2
	if diffY < 0 {
		diffY = -diffY
	}
	return diffX != 0 && diffX+diffY == 2

}
