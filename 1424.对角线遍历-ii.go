/*
 * @lc app=leetcode.cn id=1424 lang=golang
 *
 * [1424] 对角线遍历 II
 */

// @lc code=start
func findDiagonalOrder1(nums [][]int) (reslut []int) {
	n := len(nums)
	m := len(nums[0])
	for i := 0; i < len(nums); i++ {
		if len(nums[i]) > m {
			m = len(nums[i])
		}
	}
	for i := 0; i < n; i++ {
		for j, k := i, 0; j >= 0 && k < m; j-- {
			if len(nums[j]) > k {
				reslut = append(reslut, nums[j][k])
			}
			k++
		}
	}
	for i := 1; i < m; i++ {
		for j, k := i, 1; j < m && k <= n; j++ {
			if len(nums[n-k]) > j {
				reslut = append(reslut, nums[n-k][j])
			}
			k++
		}
	}
	return
}

func findDiagonalOrder(nums [][]int) (reslut []int) {
	n := len(nums)
	path := [][]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < len(nums[i]); j++ {
			if i+j < len(path) {
				path[i+j] = append(path[i+j], nums[i][j])
			} else {
				path = append(path, []int{nums[i][j]})
			}
		}
	}
	for i := 0; i < len(path); i++ {
		for j := len(path[i]) - 1; j >= 0; j-- {
			reslut = append(reslut, path[i][j])
		}
	}
	return
}

// @lc code=end
