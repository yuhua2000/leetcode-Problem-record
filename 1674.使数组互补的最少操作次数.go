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

/*
 观察以上五种情况，可以发现对于给定的数对 (a,b) 和修改目标 c，其需要的操作次数在 [2,2×limit] 的各个连续子区间上是一定的，且分界点可以直接计算。这意味着每对数在不同目标值 c 下需要操作数可以使用差分数组来维护，其前缀和就是当前目标 c 下的总操作次数。维护该操作次数并取最小值即可。

  n := len(nums)
    diff := make([]int, 2*limit+2)

    for i := 0; i < n/2; i++ {
        a := min(nums[i], nums[n-1-i])
        b := max(nums[i], nums[n-1-i])

        diff[2] += 2
        diff[a+1] -= 1
        diff[a+b] -= 1
        diff[a+b+1] += 1
        diff[b+limit+1] += 1
    }


*/
