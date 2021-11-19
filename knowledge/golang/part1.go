/**
 * @Author: xuelei02
 * @Description:
 * @File:  part1
 * @Version: 1.0.0
 * @Date: 2021/11/19 14:14
 */

package main

import "fmt"

// 值传递还是引用传递
// 只有值传递 对于map slice chan实际传的是一个引用类型
func main() {
	s := "xuelei"
	fmt.Printf("地址: %p\n", &s)
	hello(&s)
	fmt.Println(s)
	p := make(map[string]interface{})
	p["lei"] = 2
	fmt.Printf("地址: %p\n", p)
	mao(p)
	fmt.Println(p)
}

func hello(s *string) {
	fmt.Printf("地址: %p\n", &s)
	*s = "lei"
}

func mao(p map[string]interface{}) {
	fmt.Printf("%p\n", &p)
	p["lei"] = 1
}
