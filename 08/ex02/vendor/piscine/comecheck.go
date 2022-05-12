package piscine

import (
	"ft"
)

func isFtStr(s string) bool {
	if s == "42" || s == "piscine" || s == "piscine 42" {
		return true
	}
	return false
}

func printStr(s string) {
	for _, r := range s {
		ft.PrintRune(r)
	}
}

func ComCheck(a []string){
	for i := range a {
		if isFtStr(a[i]) {
			printStr("Alert!!!\n")
		}
	}
}