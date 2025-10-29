package leetcode

import "sort"

func smallestNumber(n int) int {
	ans := []int{1, 1<<2 - 1, 1<<3 - 1, 1<<4 - 1, 1<<5 - 1, 1<<6 - 1, 1<<7 - 1, 1<<8 - 1, 1<<9 - 1, 1<<10 - 1}

	return ans[sort.Search(len(ans), func(i int) bool {
		return ans[i] > n
	})]
}
