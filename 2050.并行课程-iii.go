/*
 * @lc app=leetcode.cn id=2050 lang=golang
 *
 * [2050] 并行课程 III
 */

// @lc code=start
func minimumTime(n int, relations [][]int, time []int) (result int) {
	prev := make([][]int, n+1) //构建一个先修课的表
	for i := 0; i < len(relations); i++ {
		prev[relations[i][0]] = append(prev[relations[i][0]], relations[i][1])
	}
	memo := make(map[int]int, 0) //用一个map存储课程的最晚结束时间
	var dp func(int) int
	dp = func(i int) int {
		if cur, ok := memo[i]; !ok { //如果已经计算 则直接获取之前计算的值返回即可
			for _, p := range prev[i] { //如果有先修课，计算其所有先修课中最晚结束的，从表中获取所有先修课进行计算
				cur = max(cur, dp(p))
			}
			cur += time[i-1] //加上本课程的上课时间
			memo[i] = cur    //存储下当前课程的时间 避免多次计算
			return cur
		} else {
			return cur
		}
	}
	for i := 1; i <= n; i++ { //求所有课程中最晚结束课 返回结果即可
		result = max(result, dp(i))
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
