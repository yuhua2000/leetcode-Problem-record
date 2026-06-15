package leetcode

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */

func deleteMiddle(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}
	fast := head
	slow := head
	slowPref := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slowPref = slow
		slow = slow.Next
	}
	slowPref.Next = slow.Next
	return head
}
