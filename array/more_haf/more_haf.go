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
	b := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	removeDuplicates(b)
}

func removeDuplicates(nums []int) {
	slow := 1
	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	fmt.Println(slow)
	fmt.Println(nums[:slow])
}
