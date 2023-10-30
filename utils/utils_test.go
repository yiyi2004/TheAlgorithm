package utils

import (
	"fmt"
	"testing"
)

func TestGenerateRandomArray(t *testing.T) {
	arr := GenerateRandomArray(10, 20)
	fmt.Println(arr)
}
