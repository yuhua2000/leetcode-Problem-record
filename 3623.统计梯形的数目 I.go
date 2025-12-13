package leetcode

func countTrapezoids(points [][]int) int {
	pointNum := make(map[int]int)
	mod := 1000000007

	for _, point := range points {
		pointNum[point[1]]++
	}

	ans := 0
	sum := 0

	for _, pNum := range pointNum {
		edge := pNum * (pNum - 1) / 2
		ans = (ans + edge*sum) % mod
		sum = (sum + edge) % mod
	}

	return ans
}
