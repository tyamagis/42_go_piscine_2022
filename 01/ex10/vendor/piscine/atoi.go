package piscine

func IsDigit(c rune) bool {
	return (c >= '0' && c <= '9')
}

func IsSign(c rune) int {
	if c == '-' {
		return -1
	}
	return 1
}

func IsOverflow(tmp, ret, idx int) bool {
	if idx == 20 || tmp > ret {
		return (true)
	}
	return (false)
}

func Atoi(s string) int {
	ret := 0
	sign := 1
	digit_idx := 0
	tmp := ret
	for i, c := range s {
		if (c == '+' || c == '-') && i == 0 {
			sign = IsSign(c)
			continue
		} else if IsDigit(c) == true {
			tmp = ret;
			ret = ret*10 + int(c) - int('0')
			digit_idx++;
			if IsOverflow(tmp, ret, digit_idx) {
				if sign < 0 {
					return (-9223372036854775808)
				} else {
					return (9223372036854775807)
				}
			}
		} else {
			return (0)
		}
	}
	return ret * sign
}
