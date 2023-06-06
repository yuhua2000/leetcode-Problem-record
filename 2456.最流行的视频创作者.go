/*
 * @lc app=leetcode.cn id=2456 lang=golang
 *
 * [2456] 最流行的视频创作者
 */

// @lc code=start
func mostPopularCreator(creators []string, ids []string, views []int) [][]string {
	type creator struct {
		viewSum int
		maxView int
		id      string
	}
	var creatorMap = make(map[string]*creator)
	var maxView = 0
	for i := range creators {
		name, id, view := creators[i], ids[i], views[i]
		if c, ok := creatorMap[name]; ok {
			c.viewSum += view
			if view > c.maxView || (view == c.maxView && id < c.id) {
				c.maxView = view
				c.id = id
			}
			maxView = max(maxView, c.viewSum)
		} else {
			creatorMap[name] = &creator{
				viewSum: view,
				maxView: view,
				id:      id,
			}
			maxView = max(maxView, view)
		}
	}
	var ans [][]string
	for name, c := range creatorMap {
		if c.viewSum == maxView {
			ans = append(ans, []string{name, c.id})
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
