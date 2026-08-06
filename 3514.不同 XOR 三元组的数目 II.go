package leetcode

import "slices"

func uniqueXorTriplets(nums []int) int {
	mx := slices.Max(nums) << 1
	st := make([]bool, mx)
	for i, num := range nums {
		for j := i; j < len(nums); j++ {
			st[num^nums[j]] = true
		}
	}

	s := make([]int, mx)
	for i := 0; i < len(st); i++ {
		if st[i] {
			for _, num := range nums {
				s[num^i] = 1
			}
		}
	}

	ans := 0

	for _, v := range s {
		ans += v
	}
	return ans
}
