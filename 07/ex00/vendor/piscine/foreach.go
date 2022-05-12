package piscine

import (
	"ft"
)

func printStr(s string) {
	for _, c :=range s {
		ft.PrintRune(c)
	}
}

func swap(a, b *rune) {
	*a, *b = *b, *a
}

func strLen(s string) int {
	len := 0
	for range s {
		len++
	}
	return len
}

func strRev(s string) string {
	l := strLen(s)
	rev := []rune(s)
	for i := 0; i < l / 2; i++ {
		swap(&rev[i], &rev[l - 1 - i])
	}
	return string(rev)
}

func PrintNbr(n int) {
	str := ""
	sign := 1
	if (n < 0) {
		sign = -1
	}
	for {
		str += string(rune('0' + (n % 10) * sign))
		n /= 10
		if (n == 0) {
			break
		}
	}
	if sign == -1 {
		str += string('-')
	}
	printStr(strRev(str))
}

func ForEach(f func(int), a []int){
	if a != nil && f != nil {
		for i := range a {
			f(a[i])
		}
	}
	ft.PrintRune('\n')
}