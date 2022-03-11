package main

// ------------------------------------------------ ---------
// ĆWICZENIE: Multi Short Func
//
// 1. Zadeklaruj dwie zmienne przy użyciu składni krótkich deklaracji
//
// 2. Zainicjuj zmienne za pomocą funkcji `multi` poniżej
//
// 3. Odrzuć wartość pierwszej zmiennej w deklaracji
//
// 4. Wydrukuj tylko drugą zmienną
//
// OCZEKIWANY WYNIK
// 4
// ------------------------------------------------ ---------

import "fmt"

func main() {
	_, b := multi()
	fmt.Println(b)
}

// multi to funkcja zwracająca wiele wartości int
func multi() (int, int) {
	return 5, 4
}
