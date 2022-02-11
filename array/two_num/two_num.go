/**
 * @Author: xuelei02
 * @Description: 数组中的两个数的和等于 target
 * @File:  two_num
 * @Version: 1.0.0
 * @Date: 2022/2/11 16:35
 */

package main

import "fmt"

func main() {
	a := []int{1, 2, 7, 10}
	fmt.Println(twoNum(a, 8))
}
func twoNum(arr []int, target int) []int {
	var ret []int
	tmp := make(map[int]int)

	for k, v := range arr {
		if value, ok := tmp[target-v]; ok {
			ret = append(ret, value)
			ret = append(ret, k)
		} else {
			tmp[v] = k
		}
	}
	return ret
}
