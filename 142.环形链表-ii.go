/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	//获取首次相遇时候，slow的位置
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			break
		}
	}
	//如果快指针走到尽头，没环
	if fast == nil || fast.Next == nil {
		return nil
	}
	//快指针重新出发，相遇位置就是入口位置
	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

// @lc code=end

