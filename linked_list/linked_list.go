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
		list.Next = &LinkNode{
			Val:  v,
			Next: nil,
		}
		list = list.Next
	}
	return head
}
