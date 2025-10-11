package leetcode

/*
 * @lc app=leetcode.cn id=2446 lang=golang
 *
 * [2446] 判断两个事件是否存在冲突
 */

// @lc code=start
func haveConflict(event1 []string, event2 []string) bool {
	return !(event2[1] < event1[0] || event2[0] > event1[1])
	startTime11 := (int(event1[0][0]-'0')*10+int(event1[0][1]-'0'))*60 + (int(event1[0][3]-'0')*10 + int(event1[0][4]-'0'))
	startTime12 := (int(event1[1][0]-'0')*10+int(event1[1][1]-'0'))*60 + (int(event1[1][3]-'0')*10 + int(event1[1][4]-'0'))
	startTime21 := (int(event2[0][0]-'0')*10+int(event2[0][1]-'0'))*60 + (int(event2[0][3]-'0')*10 + int(event2[0][4]-'0'))
	startTime22 := (int(event2[1][0]-'0')*10+int(event2[1][1]-'0'))*60 + (int(event2[1][3]-'0')*10 + int(event2[1][4]-'0'))
	return !(startTime22 < startTime11 || startTime21 > startTime12)
}

// @lc code=end
