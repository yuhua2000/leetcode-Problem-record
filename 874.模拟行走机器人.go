/*
 * @lc app=leetcode.cn id=874 lang=golang
 *
 * [874] 模拟行走机器人
 */

// @lc code=start
func robotSim(commands []int, obstacles [][]int) int {
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	px, py, d := 0, 0, 1
	m := make(map[int]bool, len(obstacles))
	for _, obstacles := range obstacles {
		m[obstacles[0]*60001+obstacles[1]] = true
	}
	res := 0
	for _, c := range commands {
		if c < 0 {
			if c == -1 {
				d = (d + 1) % 4
			} else {
				d = (d + 3) % 4
			}
		} else {
			for i := 0; i < c; i++ {
				if m[(px+dirs[d][0])*60001+py+dirs[d][1]] {
					break
				}
				px += dirs[d][0]
				py += dirs[d][1]
				res = max(res, px*px+py*py)
			}
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
