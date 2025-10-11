package leetcode

/*
 * @lc app=leetcode.cn id=1222 lang=golang
 *
 * [1222] 可以攻击国王的皇后
 */

// @lc code=start
func queensAttacktheKing(queens [][]int, king []int) (result [][]int) {
	queensMap := map[int]bool{}
	for _, queen := range queens {
		queensMap[queen[0]*8+queen[1]] = true
	}

	x, y := king[0], king[1]
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			kx, ky := x+dx, y+dy
			for kx >= 0 && kx < 8 && ky >= 0 && ky < 8 {
				if queensMap[kx*8+ky] {
					result = append(result, []int{kx, ky})
					break
				}

				kx += dx
				ky += dy
			}
		}
	}
	return result

}

// @lc code=end
