/**
 * @Author: xuelei02
 * @Description:
 * @File:  jiao_ji
 * @Version: 1.0.0
 * @Date: 2021/10/26 11:18
 */

package main

import (
	"algorithm/sort/quick_sort"
	"fmt"
	"sort"
)

// Intersection 求两个数组的交集
func Intersection(a, b []int) []int {
	if len(a) == 0 || len(b) == 0 {
		return nil
	}
	tmp := make(map[int]int, len(a))
	for _, v := range a {
		tmp[v] += 1
	}
	var k int
	for _, v := range b {
		if tmp[v] > 0 {
			tmp[v] -= 1
			b[k] = v
			k++
		}
	}
	return b[0:k]
}

// IntersectionSort 有序的两个数组
// 双指针
func IntersectionSort(a, b []int) []int {
	i, j, k := 0, 0, 0
	quick_sort.QuickSort(a, 0, len(a)-1)
	sort.Ints(b)
	fmt.Println(a)
	fmt.Println(b)

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		} else {
			a[k] = a[i]
			i++
			j++
			k++
		}
	}
	return a[0:k]
}

func main() {
	res := IntersectionSort([]int{4, 9, 5}, []int{9, 4, 9, 8, 4})
	fmt.Println(res)
}
