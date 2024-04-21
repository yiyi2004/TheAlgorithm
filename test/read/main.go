package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	readString, _ := reader.ReadString('\n')
	fields := strings.Fields(readString)
	for _, v := range fields {
		fmt.Println(v)
	}
}
