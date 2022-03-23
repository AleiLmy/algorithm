/**
 * @Author: xuelei02
 * @Description:
 * @File:  tire
 * @Version: 1.0.0
 * @Date: 2022/3/3 20:41
 */

package tree

import "fmt"

type TrieNode struct {
	Children [26]*TrieNode
	isEnd    bool
}

func NewTrie() *TrieNode {
	return &TrieNode{}
}

func (t *TrieNode) insert(word string) {
	node := t
	for _, v := range word {
		idx := v - 'a'
		if node.Children[idx] == nil {
			node.Children[idx] = &TrieNode{}
		}
		node = node.Children[idx]
	}
	node.isEnd = true
}
func (t *TrieNode) SearchPrefix(s string) *TrieNode {
	node := t
	for _, v := range s {
		idx := v - 'a'
		if node.Children[idx] == nil {
			return nil
		}
		node = node.Children[idx]
	}
	return node
}

// https://leetcode-cn.com/problems/trapping-rain-water/
func trap(height []int) int {
	n := len(height)
	if n <= 2 {
		return 0
	}
	var ret int
	lmax, rmax := make([]int, n), make([]int, n)
	lmax[0] = height[0]
	rmax[n-1] = height[n-1]
	for i := 1; i < n; i++ {
		lmax[i] = max(lmax[i-1], height[i])
		rmax[n-1-i] = max(rmax[n-i], height[n-1-i])
	}
	for i := 0; i < len(height); i++ {
		ret += min(lmax[i], rmax[i]) - height[i]
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// https://leetcode-cn.com/problems/median-of-two-sorted-arrays/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var (
		n1, n2      = len(nums1), len(nums2)
		left, right int
		l1, l2      int
	)
	length := n1 + n2
	for i := 0; i <= length/2; i++ {
		left = right
		if l1 < n1 && (l2 >= n2 || nums2[l2] > nums1[l1]) {
			right = nums1[l1]
			l1++
		} else {
			right = nums2[l2]
			l2++
		}
	}
	if length%2 != 0 {
		return float64(right)
	}
	fmt.Println(left, right)
	return float64(right+left) / 2.0
}

// https://leetcode-cn.com/problems/top-k-frequent-elements/
func topKFrequent(nums []int, k int) []int {
	var (
		tmp = make(map[int]int, 0)
		s   []int
	)

	for _, v := range nums {
		v1, ok := tmp[v]
		if ok {
			tmp[v] = v1 + 1
		} else {
			tmp[v] = 1
			s = append(s, v)
		}
	}
	sort(s, 0, len(s)-1, tmp)
	return s[:k]
}

func sort(nums []int, s, e int, tmp map[int]int) {
	if s < e {
		i, j := s, e
		mid := tmp[nums[(s+e)>>1]]
		for i <= j {
			for tmp[nums[i]] > mid {
				i++
			}
			for tmp[nums[j]] < mid {
				j--
			}
			if i <= j {
				nums[i], nums[j] = nums[j], nums[i]
				i++
				j--
			}
		}
		if s < j {
			sort(nums, s, j, tmp)
		}
		if e > i {
			sort(nums, i, e, tmp)
		}
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	if k <= 1 {
		return nums
	}
	var ret []int
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			nums[i] = nums[i-1]
		}
		if (i+1)%k == 0 {
			ret = append(ret, nums[i])
		}
	}
	return ret
}
