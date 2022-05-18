package main

import (
	"fmt"
	"os"
)

func strLen(s string) int {
	len := 0
	for range s {
		len++
	}
	return len
}

func isSep(s, sep []rune, i int) bool {
	if i > strLen(string(s))-strLen(string(sep)) {
		return false
	}
	return string(s[i:i+strLen(string(sep))]) == string(sep)
}

func wordLen(s, sep []rune, i int) int {
	if strLen(string(sep)) == 0 {
		return 2
	}
	len := 0
	for i+len < strLen(string(s)) {
		if isSep(s, sep, i+len) {
			break
		}
		len++
	}
	return len
}

func Split(s, sep string) []string {
	var ret []string
	rs, rsep := []rune(s), []rune(sep)
	s_len, sep_len := strLen(s), strLen(sep)
	i, end := 0, 0
	for i < s_len {
		if isSep(rs, rsep, i) {
			i += sep_len
			if sep_len == 0 {
				ret = append(ret, string(rs[i:i+1]))
				i++
			}
		} else {
			end = wordLen(rs, rsep, i)
			ret = append(ret, string(rs[i:i+end]))
			i = i + end
		}
	}
	return ret
}

func BufTrim(buf []byte) string {
	i := 0
	for buf[i] != 0 {
		i++
	}
	return string(buf[0:i])
}

func ReadStdin() []string {
	buf := make([]byte, 1024)
	str := ""
	rsize := 1
	for rsize != 0 {
		rsize, _ = os.Stdin.Read(buf)
		str += BufTrim(buf)
	}
	return Split(str, "\n")
}

func main() {
	fmt.Println(ReadStdin())
}
