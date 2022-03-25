package main

import (
	"fmt"
)

// ------------------------------------------------ ---------
// ĆWICZENIE: Sumuj liczby
//
// Używając pętli, zsumuj liczby od 1 do 10.
//
// WSKAZÓWKA
// Musisz użyć instrukcji if, aby zapobiec
// drukowanie ostatniego znaku plusa.
//
// OCZEKIWANY WYNIK
// 1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 + 10 = 55
// ------------------------------------------------ ---------

func main() {
	suma:=0

	for i:=1; i<=10;i++ {
		suma+=i
	}

	for i := 1; i <=10; i++ {
		fmt.Print(i," + ")
		if i == 9 {
			fmt.Print(i+1)
			break
		}
	}
	fmt.Println(" =",suma)
}
