package piscine

import "fmt"

func ToInt(a, b string) (c, d int, stat bool) {
	n, err := fmt.Sscanf(a, "%d", &c)
	if err != nil || n != 1 {
		return 0, 0, false
	}
	n, err = fmt.Sscanf(b, "%d", &d)
	if err != nil || n != 1 {
		return 0, 0, false
	}
	return c, d, true
}

func add(a, b int){
	if a>0 && b>0 && a+b<0{
		return
	} else if a<0 && b<0 && a+b>0{
		return
	} else {
		fmt.Println(a+b)
	}
}

func sub(a, b int) {
	if a>0 && b<0 && a+b<0{
		return
	} else if a<0 && b>0 && a+b>0{
		return
	} else {
		fmt.Println(a-b)
	}
}

func multi(a, b int) {
	if a == 0 || b == 0 {
		fmt.Println(0)
	}
	max := int((^uint(0))>>1)
	min := max + 1
	if ((a>0 && b>0) || (a<0 && b<0)) && a > max/b {
		return
	} else if (a<0 && b>0) && a < min/b {
		return
	} else if (a>0 && b<0) && b < min/a {
		return
	} else {
		fmt.Println(a*b)
	}
}

func div(a, b int) {
	if b == 0 {
		fmt.Println("No division by 0")
		return
	}
	fmt.Println(a/b)
}

func mod(a, b int) {
	if b == 0 {
		fmt.Println("No modulo by 0")
		return
	}
	fmt.Println(a%b)
}

func Doop(a []string) {
	if ValidArgs(a) == false {
		return
	}
	n1, n2, st := ToInt(a[1], a[3])
	if st == false {
		return
	}
	if a[2] == "+" {
		add(n1, n2)
	} else if a[2] == "-" {
		sub(n1, n2)
	} else if a[2] == "*" {
		multi(n1, n2)
	} else if a[2] == "/" {
		div(n1, n2)
	} else if a[2] == "%" {
		mod(n1, n2)
	}
	return
}
