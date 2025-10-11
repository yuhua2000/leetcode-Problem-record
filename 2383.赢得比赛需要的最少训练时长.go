package leetcode

/*
 * @lc app=leetcode.cn id=2383 lang=golang
 *
 * [2383] 赢得比赛需要的最少训练时长
 */

// @lc code=start
func minNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	ans := 0
	n := len(energy)
	for i := 0; i < n; {
		if initialEnergy > energy[i] && initialExperience > experience[i] {
			initialEnergy -= energy[i]
			initialExperience += experience[i]
			i++
		} else if initialEnergy <= energy[i] {
			temp := energy[i] - initialEnergy + 1
			ans += temp
			initialEnergy += temp
		} else if initialExperience <= experience[i] {
			temp := experience[i] - initialExperience + 1
			ans += temp
			initialExperience += temp
		}
	}
	return ans
}

// @lc code=end
