package main

import (
	"fmt"
	p "piscine"
)

func main(){
	s := []string{"Hello", "how", "are", "you?"}
	fmt.Println(p.ConcatParams(s))
}