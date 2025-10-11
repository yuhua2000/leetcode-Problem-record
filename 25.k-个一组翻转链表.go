package leetcode

/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
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
func reverseKGroup(head *ListNode, k int) *ListNode {
	hair := &ListNode{Next: head}
	pre := hair
	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}

		next := tail.Next
		head, tail = reveres(head, tail)

		pre.Next = head  // pre 的下一个 是反转后的头部 pre 是前一个
		tail.Next = next // 接上原来的 next

		head = tail.Next // head是这会 k 个的下一个
		pre = tail       // 这会的 pre 是刚才的最后一个 也就是下一个 k个的前一个
	}

	return hair.Next
}

func reveres(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	p := head
	for prev != tail {
		next := p.Next
		p.Next = prev
		prev = p
		p = next
	}
	return tail, head
}

// @lc code=end
