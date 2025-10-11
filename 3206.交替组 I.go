package leetcode

func numberOfAlternatingGroups(colors []int) (result int) {
	for i := 0; i < len(colors); i++ {
		start := ((i - 1) + len(colors)) % len(colors)
		end := (i + 1) % len(colors)

		if colors[i] != colors[start] && colors[i] != colors[end] {
			result++
		}
	}
	return result
}
