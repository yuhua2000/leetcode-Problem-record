/*
 * @lc app=leetcode.cn id=1299 lang=golang
 *
 * [1299] 将每个元素替换为右侧最大元素
 */

// @lc code=start
func replaceElements(arr []int) (result []int) {
	maxNum := make([]int, len(arr))
	maxNum[len(arr)-1] = arr[len(arr)-1]

	for i := len(arr) - 2; i >= 0; i-- {
		maxNum[i] = max(arr[i], maxNum[i+1])
	}

	for i := 0; i < len(arr)-1; i++ {
		result = append(result, maxNum[i+1])
	}

	result = append(result, -1)
	return
}

// @lc code=end
