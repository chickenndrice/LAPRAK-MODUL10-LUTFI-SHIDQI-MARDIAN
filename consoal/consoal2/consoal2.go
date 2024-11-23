package main

import "fmt"

func main() {
	var v rune
	var huruf, vkecil, vbesar bool
	fmt.Scanf("%c", &v)
	huruf = (v >= 'a' && v <= 'z')||(v >= 'A' && v<= 'Z')
	vkecil = v== 'a' || v== 'i'|| v== 'u' || v== 'e'|| v== 'o'
	vbesar = v== 'A' || v== 'I'|| v== 'U' || v== 'E'|| v== 'O'  
	if huruf && (vkecil || vbesar) {
		fmt.Println("Vokal")
	} else if huruf && !(vkecil || vbesar) {
		fmt.Println("Konsonan")
	} else {
		fmt.Println("bukan huruf")
	}
}