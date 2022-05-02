package piscine

func ConcatParams(args []string) string {
	s := ""
	for i, ai := range args {
		if i != 0 {
			s += "\n"
		}
		s += ai
	}
	return s
}
