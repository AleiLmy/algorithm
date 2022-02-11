/**
 * @Author: xuelei02
 * @Description:
 * @File:  more_haf
 * @Version: 1.0.0
 * @Date: 2021/12/17 17:03
 */

package main

import "fmt"

func MoreHaf(array []int) int {
	var (
		ret, cnt int
	)
	for i := 0; i < len(array); i++ {
		if cnt == 0 {
			ret = array[i]
		}
		if ret == array[i] {
			cnt++
		} else {
			cnt--
		}
		if cnt > len(array)/2 {
			break
		}
	}
	return ret
}

func main() {
	fmt.Println(MoreHaf([]int{1, 2, 3, 2, 2, 2, 5, 4, 2}))
}
