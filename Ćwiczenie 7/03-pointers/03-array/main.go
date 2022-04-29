package main

import "fmt"

// ------------------------------------------------ ---------
// ĆWICZENIE:
//
// Napisz funkcję reverse która jako parametr przyjmuje TABLICĘ 10 elementową typu int i ją odwraca, zmiany w tej tablicy powinny by permanentne
//
// Wypisz wynik

func main() {
	tab := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	reverse(&tab)
	fmt.Println(tab)
}

func reverse(tab *[10]int) {
	for i, j := 0, len(tab)-1; i < j; i, j = i+1, j-1 {
		tab[i], tab[j] = tab[j], tab[i]
	}
}
