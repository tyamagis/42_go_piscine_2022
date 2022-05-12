package main

import (
	"fmt"
	p "piscine"
)

func main() {
	link := &p.List{}
	p.ListPushFront(link, "Hello")
	p.ListPushFront(link, "there")
	p.ListPushFront(link, "how are you")
	it := link.Head
	for it != nil {
		fmt.Print(it.Data, " ")
		it = it.Next
	}
	fmt.Println()
}
