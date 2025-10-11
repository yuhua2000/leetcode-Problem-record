package leetcode

/*
 * @lc app=leetcode.cn id=1904 lang=golang
 *
 * [1904] 你完成的完整对局数
 */

// @lc code=start
func numberOfRounds(loginTime string, logoutTime string) int {
	parseTime := func(time string) int {
		h := int(time[0]-'0')*10 + int(time[1]-'0')
		m := int(time[3]-'0')*10 + int(time[4]-'0')
		return h*60 + m
	}

	login := parseTime(loginTime)
	logout := parseTime(logoutTime)
	if login > logout {
		logout += 60 * 24
	}

	if (login % 15) != 0 {
		login += 15 - (login % 15)
	}

	logout = logout - logout%15

	if ret := (logout - login) / 15; ret <= 0 {
		return 0
	} else {
		return ret
	}
}

// @lc code=end
