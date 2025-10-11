package leetcode

func minNumber(nums1 []int, nums2 []int) int {
	nums1Min := 10
	nums2Min := 10
	sameNumMin := 10
	num1Map := map[int]bool{}

	for _, num := range nums1 {
		num1Map[num] = true
		if num < nums1Min {
			nums1Min = num
		}
	}

	for _, num := range nums2 {
		if num1Map[num] && num < sameNumMin {
			sameNumMin = num
		}
		if num < nums2Min {
			nums2Min = num
		}
	}

	if sameNumMin < 10 {
		return sameNumMin
	}

	if nums1Min < nums2Min {
		return nums1Min*10 + nums2Min
	} else {
		return nums2Min*10 + nums1Min
	}
}
