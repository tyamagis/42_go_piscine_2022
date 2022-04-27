package piscine

import "ft"

func PrintComb2(){
	for n0 := '0'; n0 <= '9'; n0++ {
		for n1 := '0'; n1 <= '9'; n1++ {
			for n2 := n0; n2 <= '9'; n2++ {
				for n3 :='0'; n3 <= '9'; n3++ {
					if (n0 == n2) && (n1 >= n3) {
						continue
					}
					ft.PrintRune(rune(n0))
					ft.PrintRune(rune(n1))
					ft.PrintRune(rune(' '))
					ft.PrintRune(rune(n2))
					ft.PrintRune(rune(n3))
					if (n0 != '9') || (n1 != '8') || (n2 != '9') || (n3 != '9') {
						ft.PrintRune(rune(','))
						ft.PrintRune(rune(' '))
					}
				}
			}
		}
	}
}