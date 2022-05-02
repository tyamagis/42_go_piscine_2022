package piscine

func isSeps(r rune) bool {
	return (r == ' ' || r == '\t' || r == '\n' || r == '\v' || r == '\f' || r == '\r')
}

func strLen(s string) int {
	len := 0
	for range s {
		len ++
	}
	return len
}

func SplitWhiteSpaces(s string) []string {
	rs := []rune(s)
	max_idx := strLen(s)
	var res []string
	for i, r := range rs {
		if (i != 0 && isSeps(rs[i-1]) && !isSeps(r)) || (i == 0 && !isSeps(r)) {
			var rw []rune
			for j := 0; i+j < max_idx && !isSeps(rs[i+j]); j++ {
				rw = append(rw, rs[i+j])
			}
			res = append(res, string(rw))
		}
	}
	return res
}
