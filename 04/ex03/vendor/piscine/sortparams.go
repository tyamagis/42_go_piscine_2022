package piscine

import (
	"ft"
)

func printStr(s string){
	for _, c := range s {
		ft.PrintRune(c)
	}
	ft.PrintRune('\n')
}

func RuneLen(r []rune) int {
	len := 0
	for range r {
		len++
	}
	return len
}

func strCmp(a, b string) int {
	ra,	rb := []rune(a), []rune(b)
	len_ra, len_rb := RuneLen(ra), RuneLen(rb)
	var shorter []rune
	if len_ra >= len_rb {
		shorter = rb
	} else {
		shorter = ra
	}
	i := 0
	for range shorter {
		if ra[i] != rb[i] {
			break
		}
		i++
	}
	if i >= RuneLen(shorter) {
		i--
		if len_ra > len_rb {
			return 1
		} else if len_ra < len_rb {
			return -1
		}
	}
	if ra[i] > rb[i] {
		return 1
	} else if ra[i] < rb[i] {
		return -1
	}
	return 0
}

func swapElem(s1, s2 *string){
	tmp := ""
	for _, r := range *s1 {
		tmp += string(r)
	}
	*s1 = ""
	for _, r := range *s2 {
		*s1 += string(r)
	}
	*s2 = ""
	for _, r := range tmp {
		*s2 += string(r)
	}
}

func SortParams(s []string){
	argc := 0
	for i := range s {
		for j := range s {
			if i == j {
				continue
			}
			if (i < j && strCmp(s[i], s[j]) > 0) || (i > j && strCmp(s[i], s[j]) < 0) {
				swapElem(&s[i], &s[j])
			}
		}
		argc++
	}
	for i := 1; i < argc; i++ {
		printStr(s[i])
	}
}