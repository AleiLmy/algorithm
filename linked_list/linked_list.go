/**
 * @Author: xuelei02
 * @Description:
 * @File:  linked_list
 * @Version: 1.0.0
 * @Date: 2021/10/26 10:29
 */

package linked_list

import (
	"math"
	"strconv"
)

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
func subsets1(nums []int) [][]int {
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
func subsets(nums []int) [][]int {
	var (
		tmp = make([]int, 0)
		ret [][]int
	)
	dfsSet(0, nums, tmp, &ret)
	return ret
}

func dfsSet(index int, nums []int, tmp []int, ret *[][]int) {
	t := make([]int, len(tmp))
	copy(t, tmp)
	*ret = append(*ret, t)
	for i := index; i < len(nums); i++ {
		tmp = append(tmp, nums[i])
		dfsSet(i+1, nums, tmp, ret)
		tmp = tmp[:len(tmp)-1]
	}
}

// https://leetcode-cn.com/problems/single-number-ii/
func singleNumber1(nums []int) int {
	ret := int32(0)
	for i := 0; i < 32; i++ {
		total := int32(0)
		for _, v := range nums {
			total = total + int32(v)>>i&1
		}
		if total%3 > 0 {
			ret = ret | (total%3)<<i
		}
	}
	return int(ret)
}

// https://leetcode-cn.com/problems/single-number-iii/
func singleNumber(nums []int) []int {
	var total int
	for _, v := range nums {
		total ^= v
	}
	// 求最低位的1
	low := total & -total
	var left int
	for _, v := range nums {
		if low&v > 0 {
			left ^= v
		}
	}
	return []int{left, left ^ total}

}

// https://leetcode-cn.com/problems/set-mismatch/
func findErrorNums(nums []int) []int {
	// 位运算
	var total int
	// 将数组+ 1~n的  重复就3次 丢失 1次  其余两次
	for k, v := range nums {
		total ^= v
		total ^= k + 1
	}
	// 最低位的1
	low := total & -total
	var left int
	for k, v := range nums {
		if low&v > 0 {
			left ^= v
		}
		if low&(k+1) > 0 {
			left ^= k + 1
		}
	}
	for _, v := range nums {
		if v == left {
			return []int{left, left ^ total}
		}
	}
	return []int{left ^ total, left}
}

// https://leetcode-cn.com/problems/number-of-1-bits/
func hammingWeight(num uint32) int {
	var ret int
	//for num > 0 {
	//	ret += int(num & 1)
	//	num >>= 1
	//}

	for num > 0 {
		num = num & (num - 1)
		ret++
	}

	return ret
}

// https://leetcode-cn.com/problems/count-primes/
func countPrimes(n int) int {
	var (
		ret int
		tmp = make([]bool, n)
	)
	for i := range tmp {
		tmp[i] = true
	}
	for i := 2; i < n; i++ {
		if tmp[i] {
			ret++
			for j := i * 2; j < n; j += i {
				tmp[j] = false
			}
		}
	}
	return ret
}

func missingNumber(nums []int) int {
	var ret int
	for k, v := range nums {
		ret ^= v
		ret ^= k
	}
	return ret ^ len(nums)
}

// https://leetcode-cn.com/problems/valid-parentheses/
func isValid(s string) bool {
	var stack []byte
	tmp := map[byte]byte{
		'}': '{',
		')': '(',
		']': '[',
	}
	for i := range s {
		if s[i] == '{' || s[i] == '(' || s[i] == '[' {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 || tmp[s[i]] != stack[len(stack)-1] {
				return false
			}
			// 出栈
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

// https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/
func evalRPN(tokens []string) int {
	var stack []int // 栈
	for _, v := range tokens {
		if !isCal(v) {
			v1, _ := strconv.Atoi(v)
			stack = append(stack, v1)
		} else {
			// pop量个元素
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, cal(a, b, v))
		}
	}
	return stack[0]
}

func cal(a, b int, s string) int {
	if s == "+" {
		return a + b
	}
	if s == "-" {
		return a - b
	}
	if s == "*" {
		return a * b
	}
	return a / b

}

func isCal(s string) bool {
	switch s {
	case "+":
		return true
	case "-":
		return true
	case "*":
		return true
	case "/":
		return true
	}
	return false
}

func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	// 左边界
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if nums[left] != target {
		return []int{-1, -1}
	}
	l := left
	right = len(nums) - 1
	for left < right {
		mid := (left + right + 1) >> 1
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid
		}
	}
	return []int{l, left}
}

func mySqrt(x int) int {
	l, r := 0, x
	for l < r {
		mid := l + (r-l+1)>>1
		if mid <= x/mid {
			l = mid
		} else {
			r = mid - 1
		}
	}
	return l
}

// https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	list := &ListNode{Next: head}
	pre, cur := list, head
	i := 0
	for cur != nil {
		i++
		cur = cur.Next
		if i > n {
			pre = pre.Next
		}
	}
	pre.Next = pre.Next.Next
	return list.Next
}

// https://leetcode-cn.com/problems/generate-parentheses/
func generateParenthesis(n int) []string {
	var ret []string
	back(n, n, "", &ret)
	return ret
}
func back(l, r int, tmp string, ret *[]string) {
	if r == 0 {
		*ret = append(*ret, tmp)
		return
	}
	if l > 0 {
		back(l-1, r, tmp+"(", ret)
	}
	if r > l {
		back(l, r-1, tmp+")", ret)
	}
}

// https://leetcode-cn.com/problems/next-permutation/
func nextPermutation(nums []int) {
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := len(nums) - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	// reverse
	reverse(nums[i+1:])
}

func reverse(a []int) {
	for s, e := 0, len(a); s < e/2; s++ {
		a[s], a[e-s-1] = a[e-s-1], a[s]
	}
}

// https://leetcode-cn.com/problems/min-stack/
type MinStack struct {
	stk  []int
	min  int
	size int
}

func Constructor() MinStack {
	return MinStack{stk: make([]int, 0), min: math.MaxInt64}
}

// 4.2.1.3
// min=4
// 0,
func (this *MinStack) Push(val int) {
	if this.size == 0 {
		this.min = val
	}
	this.stk = append(this.stk, val-this.min)
	if this.min > val {
		this.min = val
	}
	this.size++
}

func (this *MinStack) Pop() {
	if this.stk[this.size-1] < 0 {
		this.min -= this.stk[this.size-1]
	}
	this.stk = this.stk[:this.size-1]
	this.size--
}

func (this *MinStack) Top() int {
	if this.stk[this.size-1] >= 0 {
		return this.min + this.stk[this.size-1]
	}
	return this.min
}

func (this *MinStack) GetMin() int {
	return this.min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// https://leetcode-cn.com/problems/combination-sum/
func combinationSum(candidates []int, target int) [][]int {
	var (
		tmp []int
		ret [][]int
	)
	dfs(0, target, candidates, tmp, &ret)
	return ret
}

func dfs(index, target int, candidates, tmp []int, ret *[][]int) {
	if target == 0 {
		t := make([]int, len(tmp))
		copy(t, tmp)
		*ret = append(*ret, t)
		return
	}
	if target < 0 {
		return
	}
	for i := index; i < len(candidates); i++ {
		tmp = append(tmp, candidates[i])
		dfs(i, target-candidates[i], candidates, tmp, ret)
		tmp = tmp[:len(tmp)-1]
	}
}
