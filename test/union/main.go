package main

import "fmt"

func main() {
	u := Init(7)
	u.Uni(1, 2)
	u.Uni(2, 3)
	u.Uni(3, 4)
	fmt.Println(u.Find(1))
	fmt.Println(u.Find(2))
	fmt.Println(u.Find(3))
	fmt.Println(u.Find(4))

	fmt.Println(u.Find(5))
	u.Uni(6, 7)
	fmt.Println(u.Find(6))
	fmt.Println(u.Find(7))
	u.Uni(6, 5)
	fmt.Println(u.Find(5))
	fmt.Println(u.Find(6))
	fmt.Println(u.Find(7))

}

type Union struct {
	Father []int
}

func Init(n int) Union {
	var u Union
	u.Father = make([]int, n+1)
	for i := 1; i <= n; i++ {
		u.Father[i] = i
	}
	return u
}

func (u Union) Uni(i, j int) {
	iFather := u.Find(i)
	jFather := u.Find(j)
	u.Father[iFather] = jFather
}

func (u Union) Find(i int) int {
	if i == u.Father[i] {
		return i
	} else {
		u.Father[i] = u.Find(u.Father[i])
		return u.Father[i]
	}
}
