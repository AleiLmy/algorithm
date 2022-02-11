/**
 * @Author: xuelei02
 * @Description: 数组加+1
 * @File:  add_one
 * @Version: 1.0.0
 * @Date: 2022/2/11 15:47
 */

package main

import "fmt"

func main() {
	a := []int{9, 9, 9}
	a = addOne(a)
	fmt.Println(a)
}

func addOne(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	var (
		tmp = 0
	)
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i] = arr[i] + tmp
		tmp = 0
		if i == len(arr)-1 {
			arr[i] = arr[i] + 1
		}
		if arr[i] == 10 {
			arr[i] = arr[i] % 10
			tmp = 1
		}
	}
	if tmp == 1 {
		ret := make([]int, 0, len(arr)+1)
		ret = append(ret, 1)
		ret = append(ret, arr...)
		return ret
	}
	return arr
}
