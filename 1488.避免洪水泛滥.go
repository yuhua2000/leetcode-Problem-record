/*
 * @lc app=leetcode.cn id=1488 lang=golang
 *
 * [1488] 避免洪水泛滥
 */

// @lc code=start
func avoidFlood(rains []int) []int {
	n := len(rains)
	ans := make([]int, n)
	st := make([]int, 0)
	mp := make(map[int]int)
	for i := 0; i < n; i++ {
		ans[i] = 1
	}
	for i, rain := range rains {
		if rain == 0 {
			st = append(st, i)
		} else {
			ans[i] = -1
			if day, ok := mp[rain]; ok {
				it := sort.SearchInts(st, day)
				if it == len(st) {
					return []int{}
				}
				ans[st[it]] = rain
				copy(st[it:len(st)-1], st[it+1:len(st)])
				st = st[:len(st)-1]

			}
			mp[rain] = i
		}
	}
	return ans
}

// @lc code=end
