/**
 * @Author: evnxo
 * @Description:
 * @File:  removeNodes.go
 * @Version: 1.0.0
 * @Date: 2024/1/3 21:00
 */

package leetcode

type listNode struct {
	Val  int
	Next *listNode
}

// 反转链表
// 时间 On
// 空间 O1
func reverser(head *listNode) *listNode {
	dummy := &listNode{}
	if head != nil {
		p := head
		head = head.Next
		p.Next = dummy.Next
		dummy.Next = p
	}
	return dummy.Next
}

func removeNodes3(head *listNode) *listNode {
	head = reverser(head)
	for p := head; p.Next != nil; {
		if p.Val > p.Next.Val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return reverser(head)
}

// 栈
// 时间 On
// 空间 On
func removeNodes2(head *listNode) *listNode {
	var st []*listNode
	for ; head != nil; head = head.Next {
		st = append(st, head)
	}
	for ; len(st) > 0; st = st[:len(st)-1] {
		if head == nil || st[len(st)-1].Val >= head.Val {
			st[len(st)-1].Next = head
			head = st[len(st)-1]
		}
	}
	return head
}

// 递归
// 时间 On
// 空间 On
func removeNodes(head *listNode) *listNode {
	if head == nil {
		return nil
	}
	head.Next = removeNodes(head.Next)
	if head.Next != nil && head.Val < head.Next.Val {
		return head.Next
	} else {
		return head
	}
}
