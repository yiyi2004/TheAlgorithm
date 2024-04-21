package main

import "fmt"

func main() {
	k, x, y := 0, 0, 0
	fmt.Scan(&k, &x, &y)
	b := (x + k + 2*y) / 3.0
	c := (x + k - y) / 3.0
	a := k - b - c
	fmt.Println(a, b, c)
}
