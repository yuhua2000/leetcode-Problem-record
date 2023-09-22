func distMoney(money int, children int) int {
	if money < children {
		return -1
	}
	money -= children
	result := min(money/7, children)
	money -= result * 7
	children -= result
	if (children == 0 && money > 0) || (children == 1 && money == 3) {
		result--
	}
	return result
}

func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}
