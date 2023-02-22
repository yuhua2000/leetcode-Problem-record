/*
 * @lc app=leetcode.cn id=1223 lang=golang
 *
 * [1223] 掷骰子模拟
 */

// @lc code=start
func dieSimulator(n int, rollMax []int) int {
	const MOD int = 1000000007
	d := make([][][]int, n+1)
	for i := 0; i <= n; i++ {
		d[i] = make([][]int, 6)
		for j := 0; j < 6; j++ {
			d[i][j] = make([]int, 16)
		}
	}
	for j := 0; j < 6; j++ {
		d[1][j][1] = 1
	}
	for i := 2; i <= n; i++ {
		for j := 0; j < 6; j++ {
			for k := 1; k <= rollMax[j]; k++ {
				for p := 0; p < 6; p++ {
					if p != j {
						d[i][p][1] = (d[i][p][1] + d[i-1][j][k]) % MOD
					} else if k+1 <= rollMax[j] {
						d[i][p][k+1] = (d[i][p][k+1] + d[i-1][j][k]) % MOD
					}
				}
			}
		}
	}
	var res = 0
	for j := 0; j < 6; j++ {
		for k := 1; k <= rollMax[j]; k++ {
			res = (res + d[n][j][k]) % MOD
		}
	}
	return res

}

// @lc code=end
