package leetcode

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=165 lang=golang
 *
 * [165] 比较版本号
 */

// @lc code=start
func compareVersion(version1 string, version2 string) int {
	fmtVer1 := strings.Split(version1, ".")
	fmtVer2 := strings.Split(version2, ".")
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
