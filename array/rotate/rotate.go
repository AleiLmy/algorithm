/**
 * @Author: xuelei02
 * @Description: 旋转一个数组的前k的数字
 * @File:  rotate
 * @Version: 1.0.0
 * @Date: 2022/2/11 14:21
 */

package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(a, 2)
	fmt.Println(a)
}

func rotate(arr []int, k int) {
	lenght := len(arr)
	if k > lenght {
		return
	}
	mid := lenght / 2
	for i := 0; i < mid; i++ {
		arr[i], arr[lenght-1-i] = arr[lenght-1-i], arr[i]
	}
	reverse(arr[:k])
	reverse(arr[k:])
}

func reverse(arr []int) {
	lenght := len(arr)
	mid := lenght / 2
	for i := 0; i < mid; i++ {
		arr[i], arr[lenght-1-i] = arr[lenght-1-i], arr[i]
	}
}
