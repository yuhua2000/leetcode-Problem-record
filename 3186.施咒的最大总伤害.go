package leetcode

import (
	"sort"
)

func maximumTotalDamage(power []int) int64 {
	sort.Ints(power)

	dp := make([][2]int64, len(power))
	dp[0][0] = int64(power[0])

	bestPrev := func(i int) int64 {
		j := i - 1
		for j >= 0 && (power[j] == power[i]-1 || power[j] == power[i]-2) {
			j--
		}

		if j >= 0 {
			if power[j] == power[i] {
				return dp[j][0]
			}
			return max(dp[j][0], dp[j][1])
		}

		return 0
	}

	for i := 1; i < len(power); i++ {
		dp[i][1] = max(dp[i-1][1], dp[i-1][0])
		dp[i][0] = bestPrev(i) + int64(power[i])
	}

	return max(dp[len(power)-1][0], dp[len(power)-1][1])
}

/*
一个魔法师有许多不同的咒语。

给你一个数组 power ，其中每个元素表示一个咒语的伤害值，可能会有多个咒语有相同的伤害值。

已知魔法师使用伤害值为 power[i] 的咒语时，他们就 不能 使用伤害为 power[i] - 2 ，power[i] - 1 ，power[i] + 1 或者 power[i] + 2 的咒语。

每个咒语最多只能被使用 一次 。

请你返回这个魔法师可以达到的伤害值之和的 最大值 。



示例 1：

输入：power = [1,1,3,4]

输出：6

解释：

可以使用咒语 0，1，3，伤害值分别为 1，1，4，总伤害值为 6 。

示例 2：

输入：power = [7,1,6,6]

输出：13

解释：

可以使用咒语 1，2，3，伤害值分别为 1，6，6，总伤害值为 13 。



提示：

1 <= power.length <= 105
1 <= power[i] <= 109
*/
