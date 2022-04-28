package piscine

import "ft"

func PrintComb(){
	for n1 := '0'; n1 <= '7'; n1++ {
		for n2 := n1 + 1; n2 <= '8'; n2++ {
			for n3 := n2 + 1; n3 <= '9'; n3++ {
				ft.PrintRune(n1)
				ft.PrintRune(n2)
				ft.PrintRune(n3)
				if n1 != '7' {
					ft.PrintRune(',')
					ft.PrintRune(' ')
				}
			}
		}
	}
	ft.PrintRune('\n')
}