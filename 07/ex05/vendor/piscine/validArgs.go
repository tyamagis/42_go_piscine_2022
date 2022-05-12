package piscine

import "fmt"

func ArgLen(a []string) (l int) {
	for range a {
		l++
	}
	return l
}

func StrLen(s string) (l int) {
	for range s {
		l++
	}
	return l
}

func IsInt(s string) bool {
	num := 0
	n, err := fmt.Sscanf(s, "%d", &num)
	if err != nil || n != 1 {
		return false
	}
	return true
}

func IsOperator(s string) bool {
	if StrLen(s) != 1 {
		return false
	}
	if s != "+" && s != "-" && s != "*" && s != "/" && s != "%" {
		return false
	}
	return true
}

func ValidArgs(a []string) bool {
	if ArgLen(a) != 4 {
		return false
	}
	if !IsInt(a[1]) || !IsInt(a[3]) {
		return false
	}
	if !IsOperator(a[2]) {
		return false
	}
	return true
}
