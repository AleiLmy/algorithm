/**
 * @Author: xuelei02
 * @Description:
 * @File:  test
 * @Version: 1.0.0
 * @Date: 2022/3/8 19:58
 */

package test

import (
	"math"
	"sync"
	"sync/atomic"
)

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

func threeSumClosest(nums []int, target int) int {
	sort(nums, 0, len(nums)-1)
	m := math.MaxInt32
	update := func(cur int) {
		if abs(cur-target) < abs(m-target) {
			m = cur
		}
	}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == target {
				return target
			}
			update(sum)
			if sum > target {
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				r--
			} else {
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				l++
			}
		}
	}
	return m
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

type instance struct {
	elm  int32
	lock sync.Mutex
}

type sing struct{}

var (
	ins *instance
	s   *sing
)

func initSingle() *sing {
	if s != nil {
		return s
	}
	if atomic.LoadInt32(&ins.elm) == 0 {
		ins.lock.Lock()
		defer ins.lock.Unlock()
		if ins.elm == 0 {
			s = new(sing)
			atomic.StoreInt32(&ins.elm, 1)
		}
	}
	return s
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	tmp := make([]int, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		n1 := int(num1[i] - '0')
		for j := len(num2) - 1; j >= 0; j-- {
			n2 := int(num2[j] - '0')
			sum := n1*n2 + tmp[i+j+1]
			tmp[i+j+1] = sum % 10
			tmp[i+j] += sum / 10
		}
	}
	var ret string
	for k, v := range tmp {
		if k == 0 && v == 0 {
			continue
		}
		ret += string(rune(v + '0'))
	}
	return ret
}

func swapNumbers(numbers []int) []int {
	numbers[0] = numbers[0] ^ numbers[1]
	numbers[1] = numbers[0] ^ numbers[1]
	numbers[0] = numbers[0] ^ numbers[1]
	return numbers
}

/**
 * Definition for a binary tree node.

 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func sortedArrayToBST(nums []int) *TreeNode {
	return midSort(nums, 0, len(nums)-1)
}

func midSort(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}
	mid := (left + right + 1) >> 1
	t := &TreeNode{Val: nums[mid]}
	t.Left = midSort(nums, left, mid-1)
	t.Right = midSort(nums, mid+1, right)
	return t
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	mid := math.MaxInt32
	if root.Left != nil {
		mid = min(mid, minDepth(root.Left))
	}
	if root.Right != nil {
		mid = min(mid, minDepth(root.Right))
	}
	return mid + 1
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	var (
		ret [][]int
		tmp []int
	)
	dfsPathSum(root, tmp, &ret, targetSum)
	return ret
}

func dfsPathSum(root *TreeNode, tmp []int, ret *[][]int, target int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil && root.Val == target {
		t := make([]int, len(tmp)+1)
		copy(t, append(tmp, root.Val))
		*ret = append(*ret, t)
		return
	}
	tmp = append(tmp, root.Val)
	dfsPathSum(root.Left, tmp, ret, target-root.Val)
	dfsPathSum(root.Right, tmp, ret, target-root.Val)
}

func reverseWords(s string) string {
	var (
		left, right, i int
	)
	s1 := []rune(s)

	for i < len(s1) {
		left = i
		for i < len(s1) && s1[i] != ' ' {
			i++
		}
		right = i - 1
		for left < right {
			s1[left], s1[right] = s1[right], s1[left]
			left++
			right--
		}
		i++
	}
	return string(s1)
}

func kthSmallest(root *TreeNode, k int) int {
	n := count(root.Left)
	if k <= n {
		return kthSmallest(root.Left, k)
	} else if k == n+1 {
		return root.Val
	} else {
		return kthSmallest(root.Right, k-n-1)
	}
}
func count(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return count(root.Left) + count(root.Right) + 1
}
