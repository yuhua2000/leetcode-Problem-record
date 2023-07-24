/*
 * @lc app=leetcode.cn id=860 lang=golang
 *
 * [860] 柠檬水找零
 */

// @lc code=start
func lemonadeChange(bills []int) bool {
	var five, ten = 0, 0
	for _, b := range bills {
		if b == 5 {
			five++
		} else if b == 10 {
			if five < 1 {
				return false
			}
			five--
			ten++
		} else {
			if five == 0 || five*5+ten*10 < 15 {
				return false
			}
			if ten > 0 {
				five--
				ten--
			} else {
				five -= 3
			}
		}
	}
	return true
}

// @lc code=end
