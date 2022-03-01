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

func NewListNode(val ...int) *ListNode {
	list := &ListNode{Val: -1, Next: nil}
	head := list
	for _, v := range val {
		next := &ListNode{
			Val:  v,
			Next: nil,
		}
		list.Next = next
		list = list.Next
	}
	return head.Next
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

// https://leetcode-cn.com/problems/kth-node-from-end-of-list-lcci/
func kthToLast(head *ListNode, k int) int {
	p, q := head, head
	for i := 0; i < k; i++ {
		p = p.Next
	}
	for p != nil {
		p = p.Next
		q = q.Next
	}
	return q.Val
}

// https://leetcode-cn.com/problems/swap-nodes-in-pairs/
func swapPairs(head *ListNode) *ListNode {
	ret := &ListNode{Val: 0, Next: head}
	pre, cur := ret, head
	for cur != nil && cur.Next != nil {
		t := cur.Next
		cur.Next = t.Next
		t.Next = cur
		pre.Next = t
		pre = cur
		cur = cur.Next
	}
	return ret.Next
}

// https://leetcode-cn.com/problems/merge-two-sorted-lists/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	pre := &ListNode{}
	cur := pre
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			cur = cur.Next
			list1 = list1.Next
		} else {
			cur.Next = list2
			cur = cur.Next
			list2 = list2.Next
		}
	}
	for list1 != nil {
		cur.Next = list1
		cur = cur.Next
		list1 = list1.Next
	}
	for list2 != nil {
		cur.Next = list2
		cur = cur.Next
		list2 = list2.Next
	}
	return pre.Next
}

// https://leetcode-cn.com/problems/merge-k-sorted-lists/
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	if len(lists) == 2 {
		return mergeTwoLists(lists[0], lists[1])
	}
	mid := len(lists) / 2
	return mergeTwoLists(mergeKLists(lists[0:mid]), mergeKLists(lists[mid:]))
}

// https://leetcode-cn.com/problems/insertion-sort-list/
func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	ret := &ListNode{
		Next: head,
	}
	pre, cur := head, head.Next
	for cur != nil {
		if pre.Val <= cur.Val {
			pre = pre.Next
		} else {
			tmp := ret // 从头开始插入
			for tmp.Next.Val <= cur.Val {
				tmp = tmp.Next
			}
			pre.Next = cur.Next // 跳过这个结点
			cur.Next = tmp.Next
			tmp.Next = cur
		}
		cur = pre.Next
	}
	return ret.Next
}

func insertionSortList1(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	dummyHead := &ListNode{Next: head}
	lastSorted, curr := head, head.Next
	for curr != nil {
		if lastSorted.Val <= curr.Val {
			lastSorted = lastSorted.Next
		} else {
			prev := dummyHead
			for prev.Next.Val <= curr.Val {
				prev = prev.Next
			}
			lastSorted.Next = curr.Next
			curr.Next = prev.Next
			prev.Next = curr
		}
		curr = lastSorted.Next
	}
	return dummyHead.Next
}

// https://leetcode-cn.com/problems/sort-list/
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	t := slow.Next
	slow.Next = nil
	l1 := sortList(head)
	l2 := sortList(t)
	ret := &ListNode{}
	cur := ret
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return ret.Next
}

func reverseList1(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	ret := &ListNode{Val: -1, Next: head}
	pre := ret
	for i := 0; i < left-1; i++ {
		pre = pre.Next // 前置的一个节点
	}
	cur := pre.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = cur.Next.Next
		next.Next = pre.Next
		pre.Next = next
	}
	return ret.Next
}

// https://leetcode-cn.com/problems/reorder-list/
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	mid := slow.Next
	slow.Next = nil
	mid = reverseList(mid)
	pre := head
	merge(pre, mid)
}

func merge(l1, l2 *ListNode) {
	var l1Tmp, l2Tmp *ListNode
	for l1 != nil && l2 != nil {
		l1Tmp, l2Tmp = l1.Next, l2.Next
		l1.Next = l2
		l1 = l1Tmp
		l2.Next = l1
		l2 = l2Tmp
	}
}

// https://leetcode-cn.com/problems/rotate-list/
func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	n := 1
	cur := head
	for cur.Next != nil {
		n++
		cur = cur.Next
	}
	mod := k % n
	if mod == 0 {
		return head
	}
	cur.Next = head
	tmp := n - mod
	for tmp > 0 {
		cur = cur.Next
		tmp--
	}
	ret := cur.Next
	cur.Next = nil
	return ret
}

// https://leetcode-cn.com/problems/palindrome-linked-list/
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	pre := slow.Next
	var cur *ListNode
	for pre != nil {
		tmp := pre.Next
		pre.Next = cur
		cur = pre
		pre = tmp
	}
	for cur != nil {
		if cur.Val != head.Val {
			return false
		}
		cur, head = cur.Next, head.Next
	}
	return true
}

// https://leetcode-cn.com/problems/intersection-of-two-linked-lists/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p, q := headA, headB
	for p != q {
		if p == nil {
			p = headB
		} else {
			p = p.Next
		}
		if q == nil {
			q = headA
		} else {
			q = q.Next
		}
	}
	return q
}

// https://leetcode-cn.com/problems/odd-even-linked-list/
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre, end := &ListNode{}, &ListNode{}
	p := pre
	q := end
	i := 1
	for head != nil {
		if i%2 != 0 {
			p.Next = head
			p = head
		} else {
			q.Next = head
			q = head
		}
		head = head.Next
		i++
	}
	q.Next = nil
	p.Next = end.Next
	return pre.Next
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// https://leetcode-cn.com/problems/linked-list-cycle-ii/
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var (
		slow, fast = head, head
		cl         bool
	)
	for !cl && fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		cl = slow == fast
	}
	if !cl {
		return nil
	}
	p := head
	for p != slow {
		p = p.Next
		slow = slow.Next
	}
	return p
}

// https://leetcode-cn.com/problems/subsets/submissions/
func subsets(nums []int) [][]int {
	var ret [][]int
	if len(nums) == 0 {
		return ret
	}
	ret = append(ret, []int{})
	for _, v := range nums {
		for _, v2 := range ret {
			tmp := make([]int, len(v2))
			copy(tmp, v2)
			tmp = append(tmp, v)
			ret = append(ret, tmp)
		}
	}
	return ret
}
