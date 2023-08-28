/*
 * @lc app=leetcode.cn id=228 lang=golang
 *
 * [228] 汇总区间
 */

// @lc code=start
func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	left := 0
	str := strconv.Itoa(nums[0])
	strs := []string{}
	for i := 1; i < len(nums); i++ {
		if nums[i]-nums[i-1] != 1 {
			if i-1 == left {
				strs = append(strs, str)
				str = strconv.Itoa(nums[i])
				left = i
			} else {
				str += "->" + strconv.Itoa(nums[i-1])
				strs = append(strs, str)
				str = strconv.Itoa(nums[i])
				left = i
			}
		}
	}
	if left != len(nums)-1 {
		str += "->" + strconv.Itoa(nums[len(nums)-1])
	}
	strs = append(strs, str)
	return strs
}

// @lc code=end
