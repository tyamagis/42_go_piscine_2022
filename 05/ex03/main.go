package main

import (
	"fmt"
	p "piscine"
)

func main(){
	fmt.Printf("%#v\n", p.SplitWhiteSpaces("Hello how are you?"))
	fmt.Printf("%#v\n", p.SplitWhiteSpaces(" Hello how are you?"))
	fmt.Printf("%#v\n", p.SplitWhiteSpaces("  Hello how are you?"))
	fmt.Printf("%#v\n", p.SplitWhiteSpaces("Hello how are you? "))
	fmt.Printf("%#v\n", p.SplitWhiteSpaces("Hello how are you?  "))
	fmt.Printf("%#v\n", p.SplitWhiteSpaces("Hello \v  \f how are   \n   you?"))
	fmt.Printf("%#v\n", p.SplitWhiteSpaces("   \n Hello  \t  how are   \n   you?  \r   "))
}