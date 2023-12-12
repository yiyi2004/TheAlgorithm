package main

import (
	"fmt"
	"testing"
)

func Test_changeString(t *testing.T) {
	s := []byte{'1', '2', '3'}
	fmt.Println(string(s))
	s[1] = 'c'
	fmt.Println(string(s))
}
