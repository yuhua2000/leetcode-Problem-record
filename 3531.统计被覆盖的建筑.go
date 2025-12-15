package leetcode

import "slices"

func countCoveredBuildings1(n int, buildings [][]int) int {
	type point struct {
		l, r, u, d bool
		x          bool
	}

	graph := make([][]point, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]point, n)
	}

	for _, building := range buildings {
		graph[building[0]-1][building[1]-1].x = true
	}

	for i := 0; i < n; i++ {
		graph[i][0].l = graph[i][0].x
		graph[i][n-1].r = graph[i][n-1].x

		graph[0][i].u = graph[0][i].x
		graph[n-1][i].d = graph[n-1][i].x
	}

	for i := 1; i < n-1; i++ {
		for j := 1; j < n-1; j++ {
			graph[i][j].l = graph[i][j-1].l || graph[i][j].x
			graph[i][n-1-j].r = graph[i][n-1-j+1].r || graph[i][n-1-j].x

			graph[j][i].u = graph[j-1][i].u || graph[j][i].x
			graph[n-1-j][i].d = graph[n-1-j+1][i].d || graph[n-1-j][i].x
		}
	}

	result := 0
	for _, building := range buildings {
		x := building[0] - 1
		y := building[1] - 1
		if x == 0 || x == n-1 || y == 0 || y == n-1 {
			continue
		}

		if graph[x][y-1].l && graph[x][y+1].r && graph[x-1][y].u && graph[x+1][y].d {
			result++
		}

	}

	return result
}

func countCoveredBuildings(n int, buildings [][]int) int {
	xToY := make(map[int][]int, n)
	yToX := make(map[int][]int, n)
	for _, building := range buildings {
		x, y := building[0], building[1]
		xToY[x] = append(xToY[x], y)
		yToX[y] = append(yToX[y], x)
	}

	for _, v := range xToY {
		slices.Sort(v)
	}

	for _, v := range yToX {
		slices.Sort(v)
	}

	result := 0
	for _, building := range buildings {
		x, y := building[0], building[1]
		yList := xToY[x]
		xList := yToX[y]

		if len(yList) < 3 || len(xList) < 3 {
			continue
		}

		if yList[0] >= y || yList[len(yList)-1] <= y {
			continue
		}

		if xList[0] >= x || xList[len(xList)-1] <= x {
			continue
		}
		result++
	}

	return result
}
