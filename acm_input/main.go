package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), " ")
		_ = data
	}
}
