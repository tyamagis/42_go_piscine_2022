package main

import (
	"fmt"
	p "piscine"
)

func main(){
	s := ".|..a.b..c.|.d..|.e..|..f.|..g..|..|.h.|."
	fmt.Printf("%#v\n", p.Split(s, ".|."))

	s = ".|a.b..c.|.d..|.e..|..f.|..g..|..|.h.|."
	fmt.Printf("%#v\n", p.Split(s, ".|."))

	s = "a.b..c.|.d..|.e..|..f.|..g..|..|.h"
	fmt.Printf("%#v\n", p.Split(s, ".|."))

	s = ".|..a.b..c.|.d..|.e..|..f.|..g..|..|.h.|"
	fmt.Printf("%#v\n", p.Split(s, ""))

	s = "."
	fmt.Printf("%#v\n", p.Split(s, "."))

	s = ""
	fmt.Printf("%#v\n", p.Split(s, "."))
}