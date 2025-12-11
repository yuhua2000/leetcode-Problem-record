package leetcode

import (
	"slices"
)

func specialTriplets(nums []int) int {
	numIndex := make(map[int][]int)
	result := 0
	for i, num := range nums {
		numIndex[num] = append(numIndex[num], i)
	}

	for i := 1; i < len(nums)-1; i++ {
		target := nums[i] * 2
		if targetPos, ok := numIndex[target]; ok && len(targetPos) > 1 && targetPos[0] < i {
			mid, _ := slices.BinarySearch(targetPos, i)
			l := mid
			r := len(targetPos) - l
			if nums[i] == 0 {
				r--
			}
			result = (result + l*r) % 1000000007
		}
	}

	return result
}
