package leetcode

func maximumLength(nums []int) int {

	numCount := make(map[int]int, 0)
	for _, n := range nums {
		numCount[n]++
	}

	oneCnt := numCount[1]
	ans := oneCnt
	if oneCnt%2 == 0 {
		ans--
	}

	delete(numCount, 1)

	for num := range numCount {
		res := 0

		x := num
		for numCount[x] > 1 {
			res += 2
			x = x * x
		}

		if _, ok := numCount[x]; ok {
			ans = max(ans, res+1)
		} else {
			ans = max(ans, res-1)
		}
	}

	return ans
}
