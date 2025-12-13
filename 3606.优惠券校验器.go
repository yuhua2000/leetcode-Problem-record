package leetcode

import (
	"slices"
	"strings"
)

func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
	var pos []int
	for i := range len(code) {
		if !isActive[i] || len(code[i]) == 0 {
			continue
		}

		if slices.ContainsFunc([]byte(code[i]), func(b byte) bool {
			return !('a' <= b && b <= 'z') && !('A' <= b && b <= 'Z') && !('0' <= b && b <= '9') && b != '_'
		}) {
			continue
		}

		switch businessLine[i] {
		case "electronics", "grocery", "pharmacy", "restaurant":
			pos = append(pos, i)
		}
	}

	slices.SortFunc(pos, func(i, j int) int {
		iCatagory := businessLine[i]
		jCatagory := businessLine[j]
		if iCatagory != jCatagory {
			return strings.Compare(iCatagory, jCatagory)
		}

		return strings.Compare(code[i], code[j])
	})

	result := make([]string, 0, len(pos))
	for _, i := range pos {
		result = append(result, code[i])
	}
	return result
}
