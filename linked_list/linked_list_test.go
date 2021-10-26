/**
 * @Author: xuelei02
 * @Description:
 * @File:  linked_list_test
 * @Version: 1.0.0
 * @Date: 2021/10/26 10:33
 */

package linked_list

import (
	"fmt"
	"testing"
)

func TestNewLinkNode(t *testing.T) {
	list := NewLinkNode(1, 2, 3, 4)
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
}
