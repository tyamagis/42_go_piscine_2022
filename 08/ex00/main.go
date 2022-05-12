package main

import (
	"fmt"
	p "piscine"
)

// abcdefghijklmnopqrstuvwxyz
// opqrstuvwxyzabcdefghijklmn

func main() {
	fmt.Println(p.Rot14("abcde...vwxyz"))
	fmt.Println(p.Rot14("ABCDE,,,VWXYZ"))
	fmt.Println(p.Rot14("Hello! How are You?"))
	fmt.Println(p.Rot14("Ftue we Daf14."))
}
