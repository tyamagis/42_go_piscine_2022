package main

import "fmt"

func main() {
	src := [][]int {
		{0,0,0},
		{1,1,1},
		{2,2,2},
	}
	dst := make([][]int, len(src))
	for i, _ := range src {
		dst[i] = append(dst[i], src[i]...)
	}
	fmt.Println(src)
	fmt.Println(dst)
	src[1][1] = 5
	fmt.Println(src)
	fmt.Println(dst)
}
