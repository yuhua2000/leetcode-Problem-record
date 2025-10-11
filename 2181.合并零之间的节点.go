package leetcode

/*
 * @lc app=leetcode.cn id=2181 lang=golang
 *
 * [2181] 合并零之间的节点
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeNodes(head *ListNode) *ListNode {
	h := &ListNode{Next: head}
	next := h.Next
	for next != nil {
		next = next.Next
		v := 0
		for next != nil && next.Val != 0 {
			v += next.Val
			next = next.Next
		}
		if v == 0 {
			h.Next = nil
			break
		}
		h = h.Next
		h.Val = v

	}

	return head
}

// @lc code=end
