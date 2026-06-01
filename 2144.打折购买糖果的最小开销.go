package leetcode

import "sort"

func minimumCost(cost []int) int {
	sort.Ints(cost)

	result := 0
	for i := len(cost) - 1; i >= 0; i-- {
		if (len(cost)-i)%3 == 0 {
			continue
		}
		result += cost[i]
	}
	return result
}
