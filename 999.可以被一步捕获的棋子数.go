package leetcode

func numRookCaptures(board [][]byte) int {
	var R int
	var B = map[int]bool{}
	var p = map[int]bool{}
	for i, row := range board {
		for j, c := range row {
			switch c {
			case 'R':
				R = i*8 + j
			case 'B':
				B[i*8+j] = true
			case 'p':
				p[i*8+j] = true
			}
		}
	}
	result := 0
	for _, dire := range [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {

		for x, y := R/8, R%8; x >= 0 && x < 8 && y >= 0 && y < 8; {
			if B[x*8+y] {
				break
			}
			if p[x*8+y] {
				result++
				break
			}

			x += dire[0]
			y += dire[1]
		}

	}
	return result
}

/*
给定一个 8 x 8 的棋盘，只有一个 白色的车，用字符 'R' 表示。棋盘上还可能存在白色的象 'B' 以及黑色的卒 'p'。空方块用字符 '.' 表示。

车可以按水平或竖直方向（上，下，左，右）移动任意个方格直到它遇到另一个棋子或棋盘的边界。如果它能够在一次移动中移动到棋子的方格，
则能够 吃掉 棋子。
*/
