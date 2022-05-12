package main

import (
	"fmt"
	p "piscine"
)

func main() {
	link := &p.List{}
	p.ListPushBack(link, "Hello")
	p.ListPushBack(link, "there")
	p.ListPushBack(link, "how are you")
	for link.Head != nil {
		fmt.Println(link.Head.Data)
		link.Head = link.Head.Next
	}
}
