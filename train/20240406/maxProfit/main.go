package main

import "fmt"

func main() {
	input := []int{2, 1, 2, 0, 1}
	res := maxProfit(input)
	fmt.Println(res)
}

// 买卖股票3 买卖股票1 思路，超出时间限制
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	if len(prices) <= 3 {
		return maxProfitSingle(prices)
	}
	res := maxProfitSingle(prices)
	for i := 2; i <= len(prices)-2; i++ {
		res = maxV(res, maxProfitSingle(prices[0:i])+maxProfitSingle(prices[i:]))
	}
	return res
}

func maxProfitSingle(prices []int) int {
	minEle := prices[0]
	maxVal := 0
	for _, price := range prices {
		maxVal = maxV(maxVal, price-minEle)
		minEle = minV(minEle, price)
	}
	return maxVal

}

func maxV(v1, v2 int) int {
	if v1 < v2 {
		return v2
	}
	return v1
}
func minV(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}
