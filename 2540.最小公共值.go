package leetcode

func getCommon(nums1 []int, nums2 []int) int {
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		num1, num2 := nums1[i], nums2[j]
		if num1 == num2 {
			return num1
		}

		if num1 < num2 {
			i++
		} else {
			j++
		}
	}
	return -1
}
