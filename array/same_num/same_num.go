/**
 * @Author: xuelei02
 * @Description:
 * @File:  same_num
 * @Version: 1.0.0
 * @Date: 2022/2/14 10:46
 */

package main

func main() {
	//a := []int{1, 1, 1, 2, 2, 3}
	//fmt.Println(SameNum(a, 2))
	//
	//fmt.Println(removeDuplicates([]int{1, 1, 1, 2, 2, 3}))
	moveZeroes([]int{0, 1, 0, 3, 12})
}

func SameNum(arr []int, key int) int {
	if len(arr) <= key {
		return len(arr)
	}
	num := key
	for i := key; i < len(arr); i++ {
		if arr[i] == arr[num-key+1] {
			continue
		}
		num++
		arr[num] = arr[i]
	}
	return num
}

func removeDuplicates(arr []int) int {
	if len(arr) <= 2 {
		return len(arr)
	}
	num := 1
	for i := 2; i < len(arr); i++ {
		if arr[i] == arr[num-1] {
			continue
		}
		num++
		arr[num] = arr[i]
	}
	return num + 1
}

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] == val {
			continue
		}
		nums[i] = nums[j]
		i++
	}
	return i
}

func moveZeroes(nums []int) {
	var (
		index = 0
	)
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[index] = nums[i]
			index++
		}
	}
	for index < len(nums) {
		nums[index] = 0
		index++
	}
}

func findRepeatNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		}
	}
	return -1
}
