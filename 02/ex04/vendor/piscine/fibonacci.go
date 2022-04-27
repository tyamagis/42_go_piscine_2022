package piscine

func Fibonacci(idx int) int {
	ans := 0
	if idx < 0 {
		return -1
	} else if idx == 0 {
		return 0
	} else if idx == 1 {
		return 1
	}
	ans += Fibonacci(idx - 1) + Fibonacci(idx - 2)
	return ans
}
