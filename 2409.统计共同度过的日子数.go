package leetcode

import "strconv"

/*
 * @lc app=leetcode.cn id=2409 lang=golang
 *
 * [2409] 统计共同度过的日子数
 */

// @lc code=start
func countDaysTogether(arriveAlice string, leaveAlice string, arriveBob string, leaveBob string) int {
	datesOfMonths := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	prefixSum := make([]int, 1)
	for _, date := range datesOfMonths {
		prefixSum = append(prefixSum, prefixSum[len(prefixSum)-1]+date)
	}
	if arriveAlice > arriveBob {
		arriveBob = arriveAlice
	}
	if leaveAlice < leaveBob {
		leaveBob = leaveAlice
	}
	ans := calculateDayOfYear(leaveBob, prefixSum) - calculateDayOfYear(arriveBob, prefixSum) + 1
	if ans < 0 {
		ans = 0
	}
	return ans
}

func calculateDayOfYear(day string, prefixSum []int) int {
	month, _ := strconv.Atoi(day[:2])
	date, _ := strconv.Atoi(day[3:])
	return prefixSum[month-1] + date
}

// @lc code=end
