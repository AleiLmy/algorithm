/**
 * @Author: xuelei02
 * @Description:
 * @File:  common_pre
 * @Version: 1.0.0
 * @Date: 2021/10/26 20:09
 */

package main

import (
	"fmt"
	"strings"
)

// LongCommonPrefix 一个字符串数组的最长公共前缀
func LongCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for _, v := range strs {
		for strings.Index(v, prefix) != 0 {
			if len(prefix) == 0 {
				return ""
			}
			// 字符串减一
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}

func main() {
	res := LongCommonPrefix([]string{"flower", "flow", "flight"})
	fmt.Println(res)
}
