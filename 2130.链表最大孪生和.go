package leetcode

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	fastHead := head
	slowHead := head
	for fastHead != nil {
		fastHead = fastHead.Next.Next
		slowHead = slowHead.Next
	}

	fmt.Println(slowHead.Val)
	midHead := reverseList(slowHead)

	result := 0
	for midHead != nil {
		fmt.Println(midHead.Val)
		result = max(result, midHead.Val+head.Val)
		head = head.Next
		midHead = midHead.Next
	}
	return result
}

// 标准反转
func reverseList(head *ListNode) *ListNode {
	var prve *ListNode
	curr := head
	for curr != nil {
		next := curr.Next

		// 反转
		curr.Next = prve

		prve = curr
		curr = next

	}

	return prve
}

// 手搓版本
func reverseList1(head *ListNode) *ListNode {
	next := head.Next // 下一个节点 nextNode
	if next == nil {
		return head
	}
	head.Next = nil // 第一个节点 头节点
	for {           // 如果存在下一个
		// 先把下下一个保存起来
		temp := next.Next // 临时变量 下下一个

		// next 的 next 指向 head | 反转的关键步骤
		next.Next = head

		if temp == nil {
			break
		}

		// 更新head
		head = next

		// 更新next
		next = temp
	}

	return next
}
