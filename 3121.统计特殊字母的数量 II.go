package leetcode

func numberOfSpecialChars(word string) int {
	var charPositions [2][27]int
	// charPositions[0] -> uppercase first index
	// charPositions[1] -> lowercase last index

	for i, c := range word {
		charIndex := c & 31
		caseType := c >> 5 & 1

		if caseType == 1 {
			// 小写：记录最后出现位置
			charPositions[1][charIndex] = i + 1
		} else if charPositions[0][charIndex] == 0 {
			// 大写：只记录第一次出现位置
			charPositions[0][charIndex] = i + 1
		}
	}

	specialCount := 0

	for charIndex := 1; charIndex < 27; charIndex++ {
		firstUpperPos := charPositions[0][charIndex]
		lastLowerPos := charPositions[1][charIndex]

		if firstUpperPos != 0 &&
			lastLowerPos != 0 &&
			lastLowerPos < firstUpperPos {
			specialCount++
		}
	}

	return specialCount
}
