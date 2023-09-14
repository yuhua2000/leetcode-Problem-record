/*
 * @lc app=leetcode.cn id=630 lang=golang
 *
 * [630] 课程表 III
 */

// @lc code=start
func scheduleCourse(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})
	h := &Heap{}
	total := 0
	for _, course := range courses {
		if t := course[0]; t+total <= course[1] { // 可以完成 直接加入
			total += t
			heap.Push(h, t)
		} else if h.Len() > 0 && t < h.IntSlice[0] { // 结束时间晚 持续时间短 所以替换
			total += t - h.IntSlice[0]
			h.IntSlice[0] = t
			heap.Fix(h, 0)
		}
	}
	return h.Len()
}

type Heap struct {
	sort.IntSlice
}

func (h Heap) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}

func (h *Heap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *Heap) Pop() interface{} {
	a := h.IntSlice
	x := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return x
}

// @lc code=end
