package leetcode

func minZeroArray(nums []int, queries [][]int) int {
	n := len(nums)
	deltaArray := make([]int, n+1)
	operations := 0
	k := 0
	for i := 0; i < n; i++ {
		operations += deltaArray[i]
		for k < len(queries) && operations < nums[i] {
			l, r, v := queries[k][0], queries[k][1], queries[k][2]
			deltaArray[l] += v
			deltaArray[r] += v
			if i >= l && i <= r {
				operations += v
			}
			k++
		}
		if operations < nums[i] {
			return -1
		}
	}
	return k
}
