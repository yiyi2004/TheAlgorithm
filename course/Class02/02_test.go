package Class02

import (
	"fmt"
	"testing"
)

func Test_km(t *testing.T) {
	arr := []int{10, 10, 3, 3, 3, 4, 4, 4, 100, 100, 100}
	ans := km(arr, 3, 2)
	fmt.Println(ans)
}
