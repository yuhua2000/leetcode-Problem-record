package leetcode

func missingInteger(nums []int) int {
	prefixSum := nums[0]
	prefixEnd := 0
	i := 1
	for ; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			prefixSum += nums[i]
			prefixEnd = i
		} else {
			break
		}
	}

	seen := make(map[int]bool)
	for ; i < len(nums); i++ {
		seen[nums[i]] = true
	}

	for i := prefixSum; ; i++ {
		if seen[i] {
			continue
		}

		if i >= nums[0] && i <= nums[prefixEnd] {
			continue
		}

		return i
	}
}
