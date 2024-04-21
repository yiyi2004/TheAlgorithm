package main

import (
	"fmt"
	"testing"
)

func Test_slice1(t *testing.T) {
	s := make([]int, 10)
	s = append(s, 10)
	t.Logf("s: %v, len of s: %d, cap of s: %d", s, len(s), cap(s))
}

func Test_slice2(t *testing.T) {
	s := make([]int, 0, 10)
	s = append(s, 10)
	t.Logf("s: %v, len of s: %d, cap of s: %d", s, len(s), cap(s))
}

func Test_slice3(t *testing.T) {
	s := make([]int, 10, 11)
	s = append(s, 10)
	t.Logf("s: %v, len of s: %d, cap of s: %d", s, len(s), cap(s))
}

func Test_slice4(t *testing.T) {
	s := make([]int, 10, 12)
	s1 := s[8:]
	// len 2 cap 4
	t.Logf("s1: %v, len of s1: %d, cap of s1: %d", s1, len(s1), cap(s1))
}

func Test_slice5(t *testing.T) {
	s := make([]int, 10, 12)
	s1 := s[8:9]
	// 1 4
	t.Logf("s1: %v, len of s1: %d, cap of s1: %d", s1, len(s1), cap(s1))
}

func Test_slice6(t *testing.T) {
	s := make([]int, 10, 12)
	s1 := s[8:]
	s1[0] = -1
	t.Logf("s: %v", s)
}

func Test_slice7(t *testing.T) {
	s := make([]int, 10, 12)
	v := s[10]
	_ = v
	// 越界
	// 求问，此时数组访问是否会越界
}

func Test_slice8(t *testing.T) {
	s := make([]int, 10, 12)
	s1 := s[8:]
	s1 = append(s1, []int{10, 11, 12}...)
	fmt.Println(s)
	v := s[10]
	_ = v
	// ...
	// 求问，此时数组访问是否会越界
}

func Test_slice9(t *testing.T) {
	s := make([]int, 10, 12)
	s1 := s[8:]
	changeSlice9(s1)
	t.Logf("s: %v", s)
}

func changeSlice9(s1 []int) {
	s1[0] = -1
}

// 为什么
func Test_slice10(t *testing.T) {
	s := make([]int, 10, 12)
	s1 := s[8:]
	changeSlice10(s1)
	// 10 12
	t.Logf("s: %v, len of s: %d, cap of s: %d", s, len(s), cap(s))
	// 2 4
	// 其实底层数据结构变了，因为 change 之后的 s1 没有赋值回来，所以外面的 s1 没有变化
	// 但是如果改变了原来的值，那么就会体现出来，因为那再 s1 的访问范围内
	t.Logf("s1: %v, len of s1: %d, cap of s1: %d", s1, len(s1), cap(s1))
}

func changeSlice10(s1 []int) {
	s1 = append(s1, 10)
}

func Test_slice11(t *testing.T) {
	s := []int{0, 1, 2, 3, 4}
	s = append(s[:2], s[3:]...)
	// len 变化，但是 cap 是不变的
	t.Logf("s: %v, len: %d, cap: %d", s, len(s), cap(s))
	v := s[4]
	_ = v
	// 是否会数组访问越界
}

func Test_slice12(t *testing.T) {
	s := make([]int, 512)
	s = append(s, 1)
	t.Logf("len of s: %d, cap of s: %d", len(s), cap(s))
}
