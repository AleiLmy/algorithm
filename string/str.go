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
