package leetcode

/*
 * @lc app=leetcode.cn id=1472 lang=golang
 *
 * [1472] 设计浏览器历史记录
 */

// @lc code=start
type BrowserHistory struct {
	history []string
	index   int
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{
		index:   0,
		history: []string{homepage},
	}
}

func (b *BrowserHistory) Visit(url string) {
	b.history = append(b.history[:b.index+1], url)
	b.index++
}

func (b *BrowserHistory) Back(steps int) string {
	if b.index >= steps {
		b.index = b.index - steps
		return b.history[b.index]
	}
	b.index = 0
	return b.history[b.index]
}

func (b *BrowserHistory) Forward(steps int) string {
	if b.index+steps < len(b.history) {
		b.index += steps
		return b.history[b.index]
	}
	b.index = len(b.history) - 1
	return b.history[b.index]
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
// @lc code=end
