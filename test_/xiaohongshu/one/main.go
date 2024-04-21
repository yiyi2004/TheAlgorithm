package main

import "fmt"

func main() {
	// 笔试
	var n float64
	fmt.Scan(&n)
	if n == 2 {
		fmt.Printf("%.10f", 1.0)
		return
	} else {
		c := n * (n - 1) / 2.0
		fmt.Printf("%.10f", c)
		return
	}
}
