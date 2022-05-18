package main

// add import package
import (
	"ft"
	"os"
)

func printStr(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
	ft.PrintRune('\n')
}

// boolean -> bool
// even(nbr) -> nbr % 2
// yes -> true
// no -> false
func isEven(nbr int) bool {
	if nbr % 2 == 1 {
		return true
	} else {
		return false
	}
}

// lengyhOfArg -> len(os.Args) --> count lengthOfArg
// EvenMsg, OddMsg -> "~~"
func main() {
	lengthOfArg := 0
	for range os.Args {
		lengthOfArg++
	}
	if isEven(lengthOfArg) == true {
		printStr("I have an even number of arguments")
	} else {
		printStr("I have an odd number of arguments")
	}
}