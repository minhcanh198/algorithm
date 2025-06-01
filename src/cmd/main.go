package main

import (
	"fmt"
	"github.com/minhcanh198/algorithm/hard"
)

func main() {
	s := "aaflslflsldkalskaaa"
	t := "aaa"
	r := hard.MinWindow(s, t)

	fmt.Println(r)
}
