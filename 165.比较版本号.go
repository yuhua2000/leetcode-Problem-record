package leetcode

import (
	"strconv"
)

/*
 * @lc app=leetcode.cn id=165 lang=golang
 *
 * [165] 比较版本号
 */

// @lc code=start
func compareVersion(version1 string, version2 string) int {
	fmtVersion := func(version string) (result []string) {
		b := true
		var node []byte
		for i := 0; i < len(version); i++ {
			switch version[i] {
			case '.':
				if len(node) > 0 {
					result = append(result, string(node))
				} else {
					result = append(result, "0")
				}
				node = node[:0]
				b = true
			case '0':
				if !b {
					node = append(node, version[i])
				}
			default:
				b = false
				node = append(node, version[i])
			}
		}
		if len(node) > 0 {
			result = append(result, string(node))
		}
		return
	}

	fmtVer1 := fmtVersion(version1)
	fmtVer2 := fmtVersion(version2)
	for i := 0; i < max(len(fmtVer1), len(fmtVer2)); i++ {
		var ver1, ver2 int
		if i < len(fmtVer1) {
			ver1, _ = strconv.Atoi(fmtVer1[i])
		}
		if i < len(fmtVer2) {
			ver2, _ = strconv.Atoi(fmtVer2[i])
		}

		if ver1 < ver2 {
			return -1
		} else if ver1 > ver2 {
			return 1
		}
	}

	return 0
}

// @lc code=end
