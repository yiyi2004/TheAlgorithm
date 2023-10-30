package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{1, 3, 1, 2, 0, 5}
	k := 3

	res := maxSlidingWindow(nums, k)
	fmt.Println(res)
}

// maxSlidingWindow 我暂时有点想不清楚了。

func maxSlidingWindow(nums []int, k int) []int {
	//q := NewMyQueue()

	// 流程是怎样的呢？
	for i := 0; i < len(nums)-k+1; i++ {
		for j := i; j < i+k; j++ {
		}
	}

	return nil
}

// 其实有正确的思路之后 Coding 是比较容易的
// 你需要将经常使用的数据结构都写的熟悉一点 | 常见的算法应该都会写

// MyQueue 实现单调队列
type MyQueue struct { //创建结构体
	queue []int
}

func NewMyQueue() *MyQueue { //初始化函数
	return &MyQueue{
		queue: make([]int, 0),
	}
}

func (m *MyQueue) Push(k int) {
	//将单调队列中比k小的全都弹出
	if len(m.queue) > 0 && k > m.queue[0] { //队列中所有的元素都比k小
		m.queue = []int{}
	}
	for len(m.queue) > 0 && m.queue[len(m.queue)-1] < k { // 队列中所可能存在元素比k小
		m.queue = m.queue[:len(m.queue)-1]
	}
	m.queue = append(m.queue, k)
}

func (m *MyQueue) Pop(t int) {
	if m.queue[0] == t { //判断滑动窗口最左边元素t是不是单调队列最大值，若是弹出，否则什么都不用做
		m.queue = m.queue[1:]
	}
}

func (m *MyQueue) GetMax() int {
	if len(m.queue) != 0 {
		return m.queue[0]
	}
	return math.MinInt
}

//// 二分查找进行优化
//func maxSlidingWindow(nums []int, k int) []int {
//	maxVal := nums[0]
//	index := 0
//	var arr []int
//	arr = append(arr, nums[:k]...)
//	for i := 1; i < len(arr); i++ {
//		if arr[i] > maxVal {
//			maxVal = arr[i]
//		}
//	}
//
//	// max second max
//
//	var res []int
//	res = append(res, maxVal)
//	for i := k; i < len(nums); i++ {
//		if nums[i] == arr[index] {
//			res = append(res, maxVal)
//			arr[index] = nums[i]
//		} else if nums[i] < arr[index] {
//			arr[index] = nums[i]
//			maxVal = arr[0]
//			for j := 0; j < k; j++ {
//				if arr[j] > maxVal {
//					maxVal = arr[j]
//				}
//			}
//			res = append(res, maxVal)
//		} else if nums[i] > arr[index] {
//			if nums[i] > maxVal {
//				res = append(res, nums[i])
//				maxVal = nums[i]
//			} else {
//				res = append(res, maxVal)
//			}
//			arr[index] = nums[i]
//		}
//
//		index++
//		if index >= k {
//			index = 0
//		}
//
//	}
//	return res
//}

// 题理解错了
//func maxSlidingWindow(nums []int, k int) []int {
//	var res []int
//	pre := make([]int, len(nums)+1)
//	pre[0] = 0
//	preVal := 0
//	maxVal := nums[0]
//	for i := 0; i < len(nums); i++ {
//		preVal += nums[i]
//		pre[i+1] = preVal
//		if i >= k-1 && i <= len(nums)-1 {
//			if pre[i+1]-pre[i+1-k] > maxVal {
//				res = append(res, pre[i+1]-pre[i+1-k])
//				maxVal = pre[i+1] - pre[i+1-k]
//			} else {
//				res = append(res, maxVal)
//			}
//		}
//	}
//
//	return res
//}
