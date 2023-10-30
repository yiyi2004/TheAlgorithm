package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomArray(maxSize int, maxValue int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, int(float32(maxSize+1)*rand.Float32()))
	for i := 0; i < len(arr); i++ {
		arr[i] = (int)(float32(maxValue+1)*rand.Float32()) - (int)(float32(maxValue)*rand.Float32())
	}
	return arr
}
