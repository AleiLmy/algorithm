/**
 * @Author: xuelei02
 * @Description:
 * @File:  test
 * @Version: 1.0.0
 * @Date: 2022/3/8 19:58
 */

package test

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	var ret [][]int
	sort(nums, 0, len(nums)-1)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i >= 1 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
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
func sort(nums []int, s, e int) {
	if s < e {
		i, j := s, e
		mid := nums[(i+j)>>1]
		for i <= j {
			for nums[i] < mid {
				i++
			}
			for nums[j] > mid {
				j--
			}
			if i <= j {
				nums[i], nums[j] = nums[j], nums[i]
				i++
				j--
			}
		}
		if j > s {
			sort(nums, s, j)
		}
		if e > i {
			sort(nums, i, e)
		}
	}
}
