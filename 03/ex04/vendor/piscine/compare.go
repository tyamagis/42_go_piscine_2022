package piscine

func RuneLen(r []rune) int {
	len := 0
	for range r {
		len++
	}
	return len
}

func Compare(a, b string) int {
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
