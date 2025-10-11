package leetcode

import "sort"

/*
 * @lc app=leetcode.cn id=1333 lang=golang
 *
 * [1333] 餐厅过滤器
 */

// @lc code=start
func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {
	eligibleRestaurants := make([][]int, 0)
	for i := 0; i < len(restaurants); i++ {
		if restaurants[i][2] >= veganFriendly && restaurants[i][3] <= maxPrice && restaurants[i][4] <= maxDistance {
			eligibleRestaurants = append(eligibleRestaurants, restaurants[i])
		}
	}
	sort.Slice(eligibleRestaurants, func(i, j int) bool {
		if eligibleRestaurants[i][1] == eligibleRestaurants[j][1] {
			return eligibleRestaurants[i][0] > eligibleRestaurants[j][0]
		} else {
			return eligibleRestaurants[i][1] > eligibleRestaurants[j][1]
		}
	})
	result := make([]int, len(eligibleRestaurants))
	for i := 0; i < len(eligibleRestaurants); i++ {
		result[i] = eligibleRestaurants[i][0]
	}
	return result
}

// @lc code=end
