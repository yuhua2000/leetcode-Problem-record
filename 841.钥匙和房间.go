/*
 * @lc app=leetcode.cn id=841 lang=golang
 *
 * [841] 钥匙和房间
 */

// @lc code=start
func canVisitAllRooms(rooms [][]int) bool {
	mr := make(map[int]bool)
	mr[0] = true
	km := rooms[0]
	for len(km) > 0 {
		key := km[0]
		if mr[key] {
			km = km[1:]
			continue
		}
		mr[key] = true
		km = append(km[1:], rooms[key]...)
	}

	return len(mr) == len(rooms)
}

// @lc code=end
