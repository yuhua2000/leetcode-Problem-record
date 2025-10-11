package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * @lc app=leetcode.cn id=143 lang=golang
 *
 * [143] 重排链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList1(head *ListNode) {
	n := 0
	for root := head; root != nil; root = root.Next {
		n++
	}

	mid := head
	for i := 0; i < (n-1)/2; i++ {
		mid = mid.Next
	}

	temp := mid.Next
	mid.Next = nil
	for ; temp != nil; temp = temp.Next {
		mid.Next = &ListNode{Val: temp.Val, Next: mid.Next}
	}

	result := &ListNode{}
	ans := result
	for headNode, midNode := head, mid.Next; midNode != nil; midNode = midNode.Next {
		ans.Next = &ListNode{Val: headNode.Val, Next: &ListNode{Val: midNode.Val}}
		ans = ans.Next.Next
		headNode = headNode.Next
	}

	if n%2 == 1 {
		ans.Next = &ListNode{Val: mid.Val}
	}

	*head = *result.Next
}

func reorderList(head *ListNode) {

	mid := head
	for fast := head; fast.Next != nil && fast.Next.Next != nil; fast = fast.Next.Next {
		mid = mid.Next
	}

	temp := mid.Next
	mid.Next = nil
	for ; temp != nil; temp = temp.Next {
		mid.Next = &ListNode{Val: temp.Val, Next: mid.Next}
	}

	l1 := head
	l2 := mid.Next
	mid.Next = nil
	for l1 != nil && l2 != nil {
		l1Temp := l1.Next
		l2Temp := l2.Next

		l1.Next = l2
		l1 = l1Temp

		l2.Next = l1
		l2 = l2Temp
	}

}

// @lc code=end
