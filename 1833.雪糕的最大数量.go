package leetcode

func maxIceCream(costs []int, coins int) int {
	const bucketLen int = 1e5 + 1
	bucket := make([]int, bucketLen) // 初始为0的数组
	for _, cost := range costs {
		if cost > coins {
			continue
		}
		bucket[cost]++
	}

	result := 0
	for i, v := range bucket {
		if v == 0 {
			continue
		}
		if i > coins {
			break
		}
		ans := min(coins/i, v)
		result += ans
		coins -= ans * i

	}
	return result
}
