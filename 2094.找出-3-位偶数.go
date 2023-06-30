/*
 * @lc app=leetcode.cn id=2094 lang=golang
 *
 * [2094] 找出 3 位偶数
 */

// @lc code=start
func findEvenNumbers(digits []int) (result []int) {
	m := map[int]struct{}{}
	for i := 0; i < len(digits); i++ {
		for j := 0; j < len(digits); j++ {
			for k := 0; k < len(digits); k++ {
				if i == j || i == k || j == k || digits[i] == 0 || digits[k]%2 == 1 {
					continue
				}

				nums := digits[i]*100 + digits[j]*10 + digits[k]
				m[nums] = struct{}{}
			}
		}
	}
	for k := range m {
		result = append(result, k)
	}
	sort.Ints(result)
	return result
}

// @lc code=end
