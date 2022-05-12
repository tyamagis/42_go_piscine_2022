package main

import (
	"fmt"
	"piscine"
)

func main() {
	toConcat := []string{"Hello!", " How", " are", " you?"}
	fmt.Println(piscine.Join(toConcat, ":"))

	empty := []string{}
	fmt.Println(piscine.Join(empty, ":"))

	empty2 := []string{"", "", ""}
	fmt.Println(piscine.Join(empty2, ":"))
}
