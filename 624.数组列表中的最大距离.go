package leetcode

func maxDistance(arrays [][]int) (res int) {
	n := len(arrays[0])
	minVal := arrays[0][0]
	maxVal := arrays[0][n-1]
	for i := 1; i < len(arrays); i++ {
		n := len(arrays[i])
		res = max(res, abs(arrays[i][n-1]-minVal), abs(maxVal-arrays[i][0]))
		minVal = min(minVal, arrays[i][0])
		maxVal = max(maxVal, arrays[i][n-1])
	}
	return res
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
