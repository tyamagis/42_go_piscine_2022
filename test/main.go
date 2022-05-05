package main

import "fmt"

func main() {
	fmt.Println(0)
	fmt.Println(^0)
	fmt.Println((^int(0)) >> 1)
	fmt.Println((^0) << 1)
	fmt.Println(^uint(0))
	fmt.Println(^uint(0)>>1)
}
