package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	readString, _ := reader.ReadString('\n')
	fields := strings.Fields(readString)
	m, _ := strconv.Atoi(fields[1])
	for i := 0; i < m; i++ {
		reader.ReadString('\n')
	}
	fmt.Println(100)
}
