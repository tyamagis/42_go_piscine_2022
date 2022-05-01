package main

import (
	p "piscine"
)
func main() {
	nb := 127
	p.PrintNbrBase(nb, "0123456789")
	p.PrintNbrBase(nb, "01")
	p.PrintNbrBase(nb, "01234567")
	p.PrintNbrBase(nb, "0123456789abcdef")
	p.PrintNbrBase(nb, "abcdefghij")
	p.PrintNbrBase(nb, ".,/;:*")

	nb = -128
	p.PrintNbrBase(nb, "0123456789")
	p.PrintNbrBase(nb, "01")
	p.PrintNbrBase(nb, "01234567")
	p.PrintNbrBase(nb, "0123456789abcdef")
	p.PrintNbrBase(nb, "abcdefghij")
	p.PrintNbrBase(nb, ".,/;:*")

	p.PrintNbrBase(nb, "aa")
	p.PrintNbrBase(nb, "a1234567a89")
	p.PrintNbrBase(nb, "a+")
	p.PrintNbrBase(nb, "a-")
	p.PrintNbrBase(nb, "a")
 	p.PrintNbrBase(nb, "")
}
