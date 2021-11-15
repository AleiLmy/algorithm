/**
 * @Author: xuelei02
 * @Description:
 * @File:  reverse
 * @Version: 1.0.0
 * @Date: 2021/11/15 11:12
 */

package main

import (
	"algorithm/linked_list"
	"fmt"
)

func main() {
	list := linked_list.NewLinkNode([]int{1, 2, 3, 4}...)
	list2 := Reverse(list)
	for list2 != nil {
		fmt.Println(list2.Val)
		list2 = list2.Next
	}
}

// Reverse 双指针法
func Reverse(l *linked_list.LinkNode) *linked_list.LinkNode {
	var res *linked_list.LinkNode
	head := l
	for head != nil {
		next := head.Next
		head.Next = res
		res = head
		head = next
	}
	return res
}
