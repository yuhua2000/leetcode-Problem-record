package leetcode

import (
	"slices"
	"sort"
	"strings"
)

func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
	categoryMap := map[string]int{
		"electronics": 0,
		"grocery":     1,
		"pharmacy":    2,
		"restaurant":  3,
	}
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

	sort.Slice(pos, func(i, j int) bool {
		iCatagory := categoryMap[businessLine[pos[i]]]
		jCatagory := categoryMap[businessLine[pos[j]]]
		if iCatagory != jCatagory {
			return iCatagory < jCatagory
		}

		return strings.Compare(code[pos[i]], code[pos[j]]) < 0
	})

	result := make([]string, 0, len(pos))
	for _, i := range pos {
		result = append(result, code[i])
	}
	return result
}
