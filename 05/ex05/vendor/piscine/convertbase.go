package piscine

func ConvertBase(nb, b_from, b_to string) string {
	return ItoaBase(AtoiBase(nb, b_from), b_to)
}
