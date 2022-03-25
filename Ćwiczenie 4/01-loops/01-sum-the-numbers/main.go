package main

import "fmt"
// ------------------------------------------------ ---------
// ĆWICZENIE: Sumuj liczby
//
// 1. Używając pętli, zsumuj liczby od 1 do 10.
// 2. Wydrukuj sumę.
//
// OCZEKIWANY WYNIK
// Suma: 55
// ------------------------------------------------ ---------

func main() {
	suma:=0

	for i:=1; i<=10;i++ {
		suma+=i
	}
	fmt.Println("Suma: ",suma)
}
