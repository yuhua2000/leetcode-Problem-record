package leetcode

/*
 * @lc app=leetcode.cn id=722 lang=golang
 *
 * [722] 删除注释
 */

// @lc code=start
func removeComments(source []string) (res []string) {
	var inBlock bool
	var newLine []byte
	for _, line := range source {
		n := len(line)
		for i := 0; i < n; i++ {
			if inBlock {
				if i+1 < n && line[i] == '*' && line[i+1] == '/' {
					inBlock = false
					i++
				}
			} else {
				if i+1 < n && line[i] == '/' && line[i+1] == '*' {
					inBlock = true
					i++
				} else if i+1 < n && line[i] == '/' && line[i+1] == '/' {
					break
				} else {
					newLine = append(newLine, line[i])
				}
			}
		}
		if !inBlock && len(newLine) > 0 {
			res = append(res, string(newLine))
			newLine = nil
		}
	}
	return
}

// @lc code=end
