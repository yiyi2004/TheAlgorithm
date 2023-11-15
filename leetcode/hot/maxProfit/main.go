package main

import "fmt"

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	profit := maxProfit(prices)
	fmt.Println(profit)
}

type stack struct {
	arr []int
}

func (s *stack) push(v int) {
	s.arr = append(s.arr, v)
}
func (s *stack) peek() int {
	return s.arr[len(s.arr)-1]
}
func (s *stack) pop() int {
	v := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return v
}
func (s *stack) isEmpty() bool {
	return len(s.arr) == 0
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	maxArr := make([]int, len(prices))
	maxArr[len(prices)-1] = prices[len(prices)-1]
	for i := len(prices) - 2; i >= 0; i-- {
		if prices[i] < maxArr[i+1] {
			maxArr[i] = maxArr[i+1]
			continue
		}
		maxArr[i] = prices[i]
	}
	for i := 0; i < len(prices); i++ {
		prices[i] = maxArr[i] - prices[i]
	}
	max := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] > max {
			max = prices[i]
		}
	}
	return max
}
