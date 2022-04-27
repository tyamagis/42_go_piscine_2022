package main

import (
	"fmt"
	"strconv"
	p "piscine"
)

func main(){
	var ret_i int
	var ret_e error

	fmt.Println("mine : ", p.BasicAtoi("12345"))
	ret_i, ret_e = strconv.Atoi("12345")
	fmt.Println("Atoi : ", ret_i, ret_e, "\n")

	fmt.Println("mine : ", p.BasicAtoi("0000000012345"))
	ret_i, ret_e = strconv.Atoi("0000000012345")
	fmt.Println("Atoi : ", ret_i, ret_e, "\n")

	fmt.Println("mine : ", p.BasicAtoi("00 00000"))
	ret_i, ret_e = strconv.Atoi("00 00000")
	fmt.Println("Atoi : ", ret_i, ret_e, "\n")

	fmt.Println("=== sign ====\n")

	fmt.Println("mine : ", p.BasicAtoi("+12345"))
	ret_i, ret_e = strconv.Atoi("+12345")
	fmt.Println("Atoi : ", ret_i, ret_e, "\n")

	fmt.Println("mine : ", p.BasicAtoi("++12345"))
	ret_i, ret_e = strconv.Atoi("++12345")
	fmt.Println("Atoi : ", ret_i, ret_e, "\n")

	fmt.Println("mine : ", p.BasicAtoi("-12345"))
	ret_i, ret_e = strconv.Atoi("-12345")
	fmt.Println("Atoi : ", ret_i, ret_e, "\n")

	fmt.Println("mine : ", p.BasicAtoi("--12345"))
	ret_i, ret_e = strconv.Atoi("--12345")
	fmt.Println("Atoi : ", ret_i, ret_e, "\n")

	fmt.Println("=== int32 max << ====\n")

	fmt.Println("mine : ", p.BasicAtoi("2147483647"))
	ret_i, ret_e = strconv.Atoi("2147483647")
	fmt.Println("Atoi : ", ret_i, ret_e, "\n")

	fmt.Println("mine : ", p.BasicAtoi("2147483648"))
	ret_i, ret_e = strconv.Atoi("2147483648")
	fmt.Println("Atoi : ", ret_i, ret_e, "\n")

	fmt.Println("=== int32 min >> ====\n")

	fmt.Println("mine : ", p.BasicAtoi("-2147483648"))
	ret_i, ret_e = strconv.Atoi("-2147483648")
	fmt.Println("Atoi : ", ret_i, ret_e, "\n")

	fmt.Println("mine : ", p.BasicAtoi("-2147483649"))
	ret_i, ret_e = strconv.Atoi("-2147483649")
	fmt.Println("Atoi : ", ret_i, ret_e, "\n")

	fmt.Println("=== long max << ====\n")

	fmt.Println("mine : ", p.BasicAtoi("9223372036854775807"))
	ret_i, ret_e = strconv.Atoi("9223372036854775807")
	fmt.Println("Atoi : ", ret_i, ret_e, "\n")

	fmt.Println("mine : ", p.BasicAtoi("9223372036854775808"))
	ret_i, ret_e = strconv.Atoi("9223372036854775808")
	fmt.Println("Atoi : ", ret_i, ret_e, "\n")

	fmt.Println("mine : ", p.BasicAtoi("9223372036854775809"))
	ret_i, ret_e = strconv.Atoi("9223372036854775809")
	fmt.Println("Atoi : ", ret_i, ret_e, "\n")

	fmt.Println("=== long min >> ====\n")

	fmt.Println("mine : ", p.BasicAtoi("-9223372036854775808"))
	ret_i, ret_e = strconv.Atoi("-9223372036854775808")
	fmt.Println("Atoi : ", ret_i, ret_e, "\n")

	fmt.Println("mine : ", p.BasicAtoi("-9223372036854775809"))
	ret_i, ret_e = strconv.Atoi("-9223372036854775809")
	fmt.Println("Atoi : ", ret_i, ret_e, "\n")

	fmt.Println("mine : ", p.BasicAtoi("-9223372036854775810"))
	ret_i, ret_e = strconv.Atoi("-9223372036854775810")
	fmt.Println("Atoi : ", ret_i, ret_e, "\n")
}