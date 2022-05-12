package piscine

func countArgs(args []string) int {
	len := 0
	for range args {
		len++
	}
	return len
}

func Cat(args []string) {
	argc := countArgs(args)
	if argc == 1 {
		// open and read from stdin
	} else {
		for i := 1; i < argc; i++ {
			DisplayFile(args[i])
		}
	}
}
