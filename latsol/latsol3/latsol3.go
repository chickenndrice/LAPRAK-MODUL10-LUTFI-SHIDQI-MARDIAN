package main

import "fmt"

func main() {
	var b int
	var prima bool
	fmt.Print("bilangan : ")
	fmt.Scan(&b)

	fmt.Print("Faktor : ")
	fac := 0
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Print(" ", i)
			fac++
		}
	}

	fmt.Println()

	if fac == 2 {
		prima = true
		fmt.Println("prima:", prima)
	} else {
		fmt.Println("prima:", prima)
	}
}
