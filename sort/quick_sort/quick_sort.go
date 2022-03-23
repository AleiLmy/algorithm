/**
 * @Author: xuelei02
 * @Description:
 * @File:  quick_sort
 * @Version: 1.0.0
 * @Date: 2021/10/26 11:25
 */

package quick_sort

import "sort"

// 快排

func QuickSort(arr []int, start, end int) {
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
			QuickSort(arr, start, j)
		}
		if end > i {
			QuickSort(arr, i, end)
		}
	}
}

// https://leetcode-cn.com/problems/group-anagrams/
func groupAnagrams(strs []string) [][]string {
	tmp := make(map[string][]string)
	for _, v := range strs {
		t := []byte(v)
		sort.Slice(t, func(i, j int) bool {
			return t[i] < t[j]
		})
		tmp[string(t)] = append(tmp[string(t)], v)
	}
	var ret [][]string
	for _, v := range tmp {
		ret = append(ret, v)
	}
	return ret
}
