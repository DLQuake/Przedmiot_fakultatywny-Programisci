package main

// ------------------------------------------------ ---------
// ĆWICZENIE: Zamień
//
// 1. Zmień `color` na "orange"
// i `color2` na "green" w tym samym czasie
//
// 2. Wydrukuj zmienne
//
// OCZEKIWANY WYNIK
// orange green
// ------------------------------------------------ ---------
import "fmt"

func main() {

	color, color2 := "red", "blue"

	color, color2 = "orange", "green"

	fmt.Println(color," ",color2)

}
