package main

import (
	"fmt"
	p "piscine"
)

func main() {
	fmt.Println("0 [0]", p.ActiveBits(0)) // 0
	fmt.Println("1 [1]", p.ActiveBits(1)) // 01
	fmt.Println("2 [10]", p.ActiveBits(2)) // 10
	fmt.Println("3 [11]", p.ActiveBits(3)) // 11
	fmt.Println("4 [100]", p.ActiveBits(4)) // 100
	fmt.Println("5 [101]", p.ActiveBits(5)) // 101
	fmt.Println("6 [110]", p.ActiveBits(6)) // 110
	fmt.Println("7 [111]", p.ActiveBits(7)) // 111
	fmt.Println("max [0111...1111]", p.ActiveBits(9223372036854775807)) // 0111 ... 1111(63)
	fmt.Println("min [1000...0000]", p.ActiveBits(-9223372036854775808)) // 1000 ... 0000(1)
	fmt.Println("-1 [1111...1111]", p.ActiveBits(-1)) // 1111...1111 (64)
	fmt.Println("-2 [1111...1110]", p.ActiveBits(-2)) // 1111...1110 (63)
	fmt.Println("-3 [1111...1101]", p.ActiveBits(-3)) // 1111...1101 (63)
}
