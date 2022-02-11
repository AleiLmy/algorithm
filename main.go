/**
 * @Author: xuelei02
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2021/10/26 10:28
 */

package main

import "fmt"

func main() {
	fff := []int{2, 3, -1, 6, 1}
	quick(fff, 0, len(fff)-1)
	fmt.Println(fff)
}

func quick(arr []int, start, end int) {
	if start < end {
		i, j := start, end
		key := arr[(start+end)/2]
		for i <= j {
			for arr[i] < key {
				i++
			}
			for arr[j] > key {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}
		if start < j {
			quick(arr, start, j)
		}
		if end > i {
			quick(arr, i, end)
		}
	}
}
