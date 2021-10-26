/**
 * @Author: xuelei02
 * @Description:
 * @File:  kmp
 * @Version: 1.0.0
 * @Date: 2021/10/26 20:31
 */

package main

import "fmt"

// kmp算法 实现

// 暴力求解
func bf(str, dst string) int {
	var (
		i, j = 0, 0
	)
	for i < len(str) && j < len(dst) {
		if str[i] == dst[j] {
			i++
			j++
		} else {
			i = i - j + 1
			j = 0
		}
	}
	if j == len(dst) {
		return i - j
	} else {
		return -1
	}

}

func main() {
	fmt.Println(generateNext("abcdab"))
	fmt.Println(bf("fdsfdsfs", "sf"))
}

// next数组的意思就是 next[i] = len; 长度为i的数组的前缀和后缀相等的最大长度
func generateNext(str string) []int {
	next := make([]int, len(str))
	next[0] = -1
	var (
		j, k = 0, -1
	)
	for j < len(str)-1 {
		if k == -1 || str[k] == str[j] {
			j++
			k++
			next[j] = k
		} else {
			k = next[k]
		}
	}
	return next
}

// KMP 算法实现函数
func kmpSearch(s, p string) int {
	n := len(s)             // 主串长度
	m := len(p)             // 模式串长度
	next := generateNext(p) // 生成 next 数组
	i, j := 0, 0
	for i < n && j < m {
		// 如果主串字符和模式串字符不相等，
		// 更新模式串坏字符下标位置为好前缀最长可匹配前缀子串尾字符下标
		// 然后从这个位置重新开始与主串匹配
		// 相当于前面提到的把模式串向后移动 j - k 位
		if j == -1 || s[i] == p[j] {
			i++
			j++
		} else {
			j = next[j]
		}
	}
	if j == m {
		// 完全匹配，返回下标位置
		return i - j
	} else {
		return -1
	}
}

// 基于 KMP 算法实现查找字符串子串函数
func strStrV2(haystack, needle string) int {
	// 子串长度=0
	if len(needle) == 0 {
		return 0
	}
	//主串长度=0，或者主串长度小于子串长度
	if len(haystack) == 0 || len(haystack) < len(needle) {
		return -1
	}
	// 子串长度=1时单独判断
	if len(needle) == 1 {
		i := 0
		for ; i < len(haystack); i++ {
			if haystack[i] == needle[0] {
				return i
			}
		}
		return -1
	}

	// 其他情况走 KMP 算法
	return kmpSearch(haystack, needle)
}
