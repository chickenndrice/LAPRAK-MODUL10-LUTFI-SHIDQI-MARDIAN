package main

import (
	"fmt"
)

func main() {
	var bilangan, a, b, c, d int
	var teks string
	fmt.Print("Bilangan: ")	
	fmt.Scan(&bilangan)
	a = bilangan / 1000
	b = (bilangan % 1000) / 100
	c = (bilangan % 100) / 10
	d = bilangan % 10

	if a < b && b < c && c < d{
		teks = "Turut membesar"
	} else if a > b && b > c && c > d{
		teks = "Turur mengecil"
	} else {
		teks = "Tidak terurut"
	}
	fmt.Println("Digit pada bilangan", bilangan, teks)
}