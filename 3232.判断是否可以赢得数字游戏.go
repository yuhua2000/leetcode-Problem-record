package leetcode

func canAliceWin(nums []int) bool {
	sum := 0
	unitsSum := 0
	for _, num := range nums {

		if num < 10 {
			unitsSum += num
		} else {
			sum += num
		}
	}

	return unitsSum != sum
}
