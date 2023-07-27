/*
 * @lc app=leetcode.cn id=599 lang=golang
 *
 * [599] 两个列表的最小索引总和
 */

// @lc code=start
func findRestaurant(list1 []string, list2 []string) (result []string) {
	indexMap := make(map[string]int)
	for i := 0; i < len(list1); i++ {
		indexMap[list1[i]] = i
	}
	minIndex := 2000
	for i := 0; i < len(list2); i++ {
		if index, ok := indexMap[list2[i]]; ok {
			if index+i < minIndex {
				result = []string{list2[i]}
				minIndex = index + i
			} else if index+i == minIndex {
				result = append(result, list2[i])
			}
		}
	}
	return
}

// @lc code=end
