/**
 * @Author: xuelei02
 * @Description:
 * @File:  linked_list
 * @Version: 1.0.0
 * @Date: 2021/10/26 10:29
 */

package linked_list

type LinkNode struct {
	Val  int
	Next *LinkNode
}

func NewLinkNode(val ...int) *LinkNode {
	list := &LinkNode{Val: -1, Next: nil}
	head := list
	for _, v := range val {
		next := &LinkNode{
			Val:  v,
			Next: nil,
		}
		list.Next = next
		list = list.Next
	}
	return head
}

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode-cn.com/problems/add-two-numbers/
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := &ListNode{}
	head := ret
	var tmp int
	for l1 != nil || l2 != nil {
		s := tmp
		if l1 != nil {
			s += l1.Val
		}
		if l2 != nil {
			s += l2.Val
		}
		head.Next = &ListNode{Val: s % 10, Next: nil}
		head = head.Next
		tmp = s / 10
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if tmp != 0 {
		head.Next = &ListNode{Val: tmp, Next: nil}
	}
	return ret.Next
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		ll1  = make([]*ListNode, 0)
		ll2  = make([]*ListNode, 0)
		tmp  int
		head = &ListNode{}
	)
	for l1 != nil {
		ll1 = append(ll1, l1)
		l1 = l1.Next
	}
	for l2 != nil {
		ll2 = append(ll2, l2)
		l2 = l2.Next
	}
	// 遍历处理了
	i, j := len(ll1)-1, len(ll2)-1
	for i >= 0 || j >= 0 || tmp == 1 {
		s := tmp
		if i >= 0 {
			s += ll1[i].Val
			i--
		}
		if j >= 0 {
			s += ll2[j].Val
			j--
		}
		node := &ListNode{
			Val:  s % 10,
			Next: head.Next,
		}
		head.Next = node
		tmp = s / 10
	}
	return head.Next
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

// https://leetcode-cn.com/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/
func reversePrint(head *ListNode) []int {
	pre := reverseList(head)
	var res []int
	for pre != nil {
		res = append(res, pre.Val)
		pre = pre.Next
	}
	return res
}

// https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/
func deleteDuplicates(head *ListNode) *ListNode {
	pre := &ListNode{}
	cur := pre
	tmp := make(map[int]struct{})
	for head != nil {
		if _, ok := tmp[head.Val]; !ok {
			cur.Next = &ListNode{
				Val:  head.Val,
				Next: nil,
			}
			cur = cur.Next
			tmp[head.Val] = struct{}{}
		}
		head = head.Next
	}
	return pre.Next
}
