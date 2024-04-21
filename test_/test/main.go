package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type city struct {
	x int
	y int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	readString, _ := reader.ReadString('\n')
	fields := strings.Fields(readString)
	n, _ := strconv.Atoi(fields[0])
	k, _ := strconv.Atoi(fields[1])
	inputCities := make([]city, 0)
	inputCities = append(inputCities, city{x: 0, y: 0})

	for i := 0; i < n; i++ {
		readString, _ = reader.ReadString('\n')
		fields = strings.Fields(readString)
		x, _ := strconv.Atoi(fields[0])
		y, _ := strconv.Atoi(fields[1])
		inputCities = append(inputCities, city{x: x, y: y})
	}
	sort.Slice(inputCities, func(i, j int) bool {
		return inputCities[i].x < inputCities[j].x
	})

	l, r := 1, 1
	val := 0
	maxHappy := 0
	for r <= n {
		for inputCities[r].x-inputCities[l].x < k {
			val += inputCities[r].y
			r++
			if r > n {
				break
			}
		}
		if val > maxHappy {
			maxHappy = val
		}
		val -= inputCities[l].y
		l++
	}
	fmt.Println(maxHappy)
}
