/**
 * @Author: xuelei02
 * @Description:
 * @File:  str
 * @Version: 1.0.0
 * @Date: 2022/2/18 16:36
 */

package main

import "strings"

func main() {

}

// https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	var ret int
	if len(s) == 0 {
		return ret
	}
	left, right := 0, 0
	for right < len(s) {
		for i := left; i < right; i++ {
			if s[i] == s[right] {
				ret = max(ret, right-left)
				left = i + 1
				break
			}
		}
		right++
	}
	ret = max(ret, right-left)

	return ret
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// https://leetcode-cn.com/problems/longest-common-prefix/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	p := strs[0]

	for i := 0; i < len(strs); {
		if len(p) == 0 {
			return p
		}
		if t := strings.Index(strs[i], p); t != 0 {
			p = p[:len(p)-1]
			continue
		}
		i++
	}
	return p
}

//https://leetcode-cn.com/problems/reverse-vowels-of-a-string/
func reverseVowels(s string) string {
	b := []byte(s)
	i, j := 0, len(s)-1
	for i < j {
		for i < j && !isCowel(b[i]) {
			i++
		}
		for i < j && !isCowel(b[j]) {
			j--
		}
		if i < j {
			b[i], b[j] = b[j], b[i]
			i++
			j--
		}
	}
	return string(b)
}
func isCowel(i byte) bool {
	switch i {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}

// https://leetcode-cn.com/problems/longest-palindromic-substring/
func longestPalindrome(s string) string {
	length := len(s)
	var ret string
	for i := 0; i < length; i++ {
		for j := 0; j < 2; j++ {
			l, r := i, i+j
			for l >= 0 && r < length && s[l] == s[r] {
				if (r - l + 1) > len(ret) {
					ret = s[l : r+1]
				}
				l--
				r++
			}
		}
	}
	return ret
}

func countSubstrings(s string) int {
	length := len(s)
	var ret int
	if length == 0 || length == 1 {
		return length
	}
	for i := 0; i < length; i++ {
		for j := 0; j < 2; j++ {
			l, r := i, i+j
			for l >= 0 && r < length && s[l] == s[r] {
				ret++
				l--
				r++
			}
		}
	}
	return ret
}

// https://leetcode-cn.com/problems/container-with-most-water/
func maxArea(height []int) int {
	var (
		res  int
		l, r = 0, len(height) - 1
	)
	for l <= r {
		if height[l] < height[r] {
			res = max(res, height[l]*(r-l))
			l++
		} else {
			res = max(res, height[r]*(r-l))
			r--
		}
	}
	return res
}
