package leetcode

import "sort"

func minMoves(nums []int, limit int) int {
	n := len(nums)
	sumCount := make(map[int]int)
	minArr := make([]int, n/2)
	maxArr := make([]int, n/2)

	for i := 0; i < n/2; i++ {
		a := min(nums[i], nums[n-i-1])
		b := max(nums[i], nums[n-i-1])

		sumCount[a+b]++
		maxArr[i] = b
		minArr[i] = a
	}

	sort.Ints(maxArr)
	sort.Ints(minArr)

	minOps := n
	for c := 2; c <= 2*limit; c++ {
		addLeft := n/2 - lowerBound(minArr, c)
		addRight := lowerBound(maxArr, c-limit)
		currentOps := n/2 + addLeft + addRight - sumCount[c]

		minOps = min(minOps, currentOps)
	}

	return minOps
}

func lowerBound(arr []int, target int) int {
	left, right := 0, len(arr)
	for left < right {
		mid := (left + right) / 2
		if arr[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
