/**
 * @Author: xuelei02
 * @Description: 数组中的两个数的和等于 target
 * @File:  two_num
 * @Version: 1.0.0
 * @Date: 2022/2/11 16:35
 */

package main

import (
	"fmt"
	"sort"
)

func main() {
	//a := []int{1, 2, 7, 10}
	//fmt.Println(twoNum(a, 8))

	b := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(b))

	productExceptSelf([]int{1, 2, 3, 4})
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

// 三数之和
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	var ret [][]int
	// 排序
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i >= 1 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[left] + nums[right] + nums[i]
			if sum == 0 {
				ret = append(ret, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return ret
}

func threeSumSmall(nums []int, target int) int {
	var count int
	if len(nums) < 3 {
		return count
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		count += cal(nums, i+1, len(nums)-1, target-nums[i])
	}
	return count
}

func cal(nums []int, s, e int, tar int) int {
	var c int
	for s < e {
		if nums[s]+nums[e] < tar {
			c += e - s
			s++
		} else {
			e--
		}
	}
	return c
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	right := m + n - 1
	for i, j := m-1, n-1; i >= 0 || j >= 0; right-- {
		if i == -1 {
			nums1[right] = nums2[j]
			j--
			continue
		}
		if j == -1 {
			nums1[right] = nums1[i]
			i--
			continue
		}
		if nums1[i] >= nums2[j] {
			nums1[right] = nums1[i]
			i--
		} else {
			nums1[right] = nums2[j]
			j--
		}
	}
}

func findMin(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	l, r := 0, len(nums)-1
	for l < r {
		mid := (l + r) / 2
		if nums[mid] < nums[r] {
			r = mid
		} else if nums[mid] > nums[r] {
			l = mid + 1
		} else {
			r--
		}
	}
	return nums[l]
}

func productExceptSelf(nums []int) []int {
	l := make([]int, len(nums))
	l[0] = 1
	for i := 1; i < len(nums); i++ {
		l[i] = l[i-1] * nums[i-1]
	}
	r := make([]int, len(nums))
	r[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		r[i] = r[i+1] * nums[i+1]
	}
	for i := 0; i < len(nums); i++ {
		l[i] = l[i] * r[i]
	}
	return l
}
