package main

import (
	"fmt"
	p "piscine"
)
func main() {
	fmt.Println(p.AtoiBase("92233720368547758", "0123456789"))
	fmt.Println(p.AtoiBase("7fffffffffffffff", "0123456789abcdef"))
	// fmt.Println("1", "0123456789", p.AtoiBase("1", "0123456789"))
	// fmt.Println("2", "0123456789", p.AtoiBase("2", "0123456789"))
	// fmt.Println("345", "0123456789", p.AtoiBase("345", "0123456789"))
	// fmt.Println("6789", "0123456789", p.AtoiBase("6789", "0123456789"))
	// fmt.Println("10000001", "01", p.AtoiBase("10000001", "01"))
	// fmt.Println("00000001", "01", p.AtoiBase("00000001", "01"))
	// fmt.Println("177", "01234567", p.AtoiBase("177", "01234567"))
	// fmt.Println("0ffff", "0123456789abcdef", p.AtoiBase("0ffff", "0123456789abcdef"))
	// fmt.Println("42tokyo", "01234kotya", p.AtoiBase("42tokyo", "01234kotya"))
	// fmt.Println("uoi", "choumi", p.AtoiBase("uoi", "choumi"))
	// fmt.Println("OOOl0001", "Ol01", p.AtoiBase("OOOl0001", "Ol01"))

	// fmt.Println("uoi", "choumi+", p.AtoiBase("uoi", "choumi+"))
	// fmt.Println("uoi", "choumi-", p.AtoiBase("uoi", "choumi-"))
	// fmt.Println("aaa", "choumi", p.AtoiBase("aaa", "choumi")) // not required
	// fmt.Println("uoi", "choumic", p.AtoiBase("uoi", "choumic"))
	// fmt.Println("uuuuu", "u", p.AtoiBase("uuuuu", "u"))
	// fmt.Println("uoi", "u", p.AtoiBase("uoi", "u"))
	// fmt.Println("", "choumi", p.AtoiBase("", "choumi")) // not required
}
