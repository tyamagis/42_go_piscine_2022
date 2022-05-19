package main

import (
	"fmt"
)

func main() {
	n := int((^uint(0))>>1)
	n += 1
	fmt.Println(n)
}
