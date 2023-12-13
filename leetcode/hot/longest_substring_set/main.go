package main

import (
	"fmt"
)

// 滑动窗口解决问题
// 滑动窗口的思想是：
// 1. 右侧指针一直往右移动，知道不满足某个条件，在该题中值得是集合中已经存在当前右指针
// 指向的元素
// 2. 移动左侧指针，知道右侧指针满足可以右移的条件。
// 该题中指的是弹出 set 中的元素。

func main() {
	arr := "abcdefghizklmnacabcdefg"
	res := lengthOfLongestSubstring(arr)
	fmt.Println(res)
}

func lengthOfLongestSubstring(s string) int {
	set := initSet()
	left := 0
	right := 0
	length := 0
	maxLen := 0
	for right < len(s) {
		for right < len(s) && !set.Exists(s[right]) {
			set.Set(s[right])
			right++
			length++
		}
		if length > maxLen {
			maxLen = length
		}
		for left < len(s) && right < len(s) && set.Exists(s[right]) {
			set.Delete(s[left])
			left++
			length--
		}
	}

	return maxLen
}

func initSet() set {
	return make(map[uint8]struct{})
}

type set map[uint8]struct{}

func (s *set) Exists(key uint8) bool {
	_, ok := (*s)[key]
	return ok
}
func (s *set) Set(key uint8) {
	(*s)[key] = struct{}{}
}
func (s *set) Delete(key uint8) {
	if s.Exists(key) {
		delete(*s, key)
	} else {
		panic("key 不存在")
	}
}

func (s *set) Length() int {
	return len(*s)
}
