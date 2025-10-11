package leetcode

func giveGem(gem []int, operations [][]int) int {
	for _, operations := range operations {
		temp := gem[operations[0]] / 2
		gem[operations[1]] += temp
		gem[operations[0]] -= temp
	}
	max, min := gem[0], gem[1]
	for _, g := range gem {
		if g > max {
			max = g
		}
		if g < min {
			min = g
		}
	}
	return max - min
}
