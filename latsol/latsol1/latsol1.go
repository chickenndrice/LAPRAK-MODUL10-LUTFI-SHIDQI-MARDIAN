package main

import "fmt"

func main() {
	var berat int
	fmt.Scan(&berat)
	kg := berat / 1000
	gr := berat % 1000
	tambahan := gr * 5
	tambahanr := gr * 15
	biaya := kg * 10000
	total := biaya + tambahan
	totalr := biaya + tambahanr
	if kg > 10 && gr < 500 {
		fmt.Println("Berat parsel (gram): ", berat)
		fmt.Println("Detail berat: ", kg, "kg +", gr, "gr")
		fmt.Println("Detail biaya: Rp.", biaya, "+ Rp.", tambahanr)
		fmt.Println("Rp.", biaya)
	} else if kg > 10 && gr >= 500 {
		fmt.Println("Berat parsel (gram): ", berat)
		fmt.Println("Detail berat: ", kg, "kg +", gr, "gr")
		fmt.Println("Detail biaya: Rp.", biaya, "+ Rp.", tambahan)
		fmt.Println("Rp.", biaya)
	} else if gr < 500 {
		fmt.Println("Berat parsel (gram): ", berat)
		fmt.Println("Detail berat: ", kg, "kg +", gr, "gr")
		fmt.Println("Detail biaya: Rp.", biaya, "+ Rp.", tambahanr)
		fmt.Println("Rp.", totalr)
	} else if gr >= 500 {
		fmt.Println("Berat parsel (gram): ", berat)
		fmt.Println("Detail berat: ", kg, "kg +", gr, "gr")
		fmt.Println("Detail biaya: Rp.", biaya, "+ Rp.", tambahan)
		fmt.Println("Rp.", total)
	}

}
