package main

// ------------------------------------------------ ---------
// ĆWICZENIE:
//
// Wydrukuj oczekiwane wyniki,
// w zależności od zmiennej age.
//
// OCZEKIWANY WYNIK
// Jeśli wiek jest większy niż 60 lat, wypisz:
//    	Getting older
// Jeśli wiek jest większy niż 30, wypisz:
// 		Getting wiser
// Jeśli wiek jest większy niż 20, wypisz:
// 		Adulthood
// Jeśli wiek jest większy niż 10, wypisz:
//    Young blood
// W przeciwnym razie wypisz:
// 	  Booting up
// ------------------------------------------------ ---------

import "fmt"

func main() {
	// Zmień wartość zmiennej age, aby uzyskać oczekiwane wyniki
	age := 20

	if age > 60 {
		fmt.Println("Getting older")
	} else if age > 30 {
		fmt.Println("Getting wiser")
	} else if age > 20{
		fmt.Println("Adulthood")
	} else if age > 10{
		fmt.Println("Young blood")
	} else {
		fmt.Println("Booting up")
	}
}
