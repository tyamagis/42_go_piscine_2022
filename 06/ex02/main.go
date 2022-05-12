package main

import (
	"os"
	p "piscine"
)

func main(){
	arg_len := 0
	for range os.Args {
		arg_len++
	}
	if arg_len == 1 {
		p.PrintStr("File name missing")
		return
	} else if arg_len >= 3 {
		p.PrintStr("Too many arguments")
		return
	}
	p.DisplayFile(os.Args[1])
}