package main

import (
	"fmt"
	"strconv"
	"strings"
)

func MyAtoi(s string) int {
	s = strings.Trim(s, " ")
	i, err := strconv.Atoi(s)
	fmt.Println(err)
	return i
}

func main() {
	fmt.Println("String to integer casting")

}
