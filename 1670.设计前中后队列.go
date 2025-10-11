package leetcode

/*
 * @lc app=leetcode.cn id=1670 lang=golang
 *
 * [1670] 设计前中后队列
 */

// @lc code=start

type FrontMiddleBackQueue struct {
	queue []int
}

func Constructor() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{}
}

func (this *FrontMiddleBackQueue) PushFront(val int) {
	this.queue = append([]int{val}, this.queue...)
}

func (this *FrontMiddleBackQueue) PushMiddle(val int) {
	index := len(this.queue) / 2
	this.queue = append(this.queue[0:index], append([]int{val}, this.queue[index:]...)...)
}

func (this *FrontMiddleBackQueue) PushBack(val int) {
	this.queue = append(this.queue, val)
}

func (this *FrontMiddleBackQueue) PopFront() int {
	if len(this.queue) == 0 {
		return -1
	}
	result := this.queue[0]
	this.queue = this.queue[1:]
	return result
}

func (this *FrontMiddleBackQueue) PopMiddle() int {
	if len(this.queue) == 0 {
		return -1
	}
	mid := (len(this.queue) - 1) / 2

	result := this.queue[mid]
	this.queue = append(this.queue[0:mid], this.queue[mid+1:]...)
	return result
}

func (this *FrontMiddleBackQueue) PopBack() int {
	if len(this.queue) == 0 {
		return -1
	}
	result := this.queue[len(this.queue)-1]
	this.queue = this.queue[:len(this.queue)-1]
	return result
}

/**
 * Your FrontMiddleBackQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PushFront(val);
 * obj.PushMiddle(val);
 * obj.PushBack(val);
 * param_4 := obj.PopFront();
 * param_5 := obj.PopMiddle();
 * param_6 := obj.PopBack();
 */
// @lc code=end
