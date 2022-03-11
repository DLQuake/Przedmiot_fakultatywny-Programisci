package main

// ------------------------------------------------ ---------
// ĆWICZENIE:
// 1. Zobacz w dokumentacji jak działa funkcja z pakietu path.Split
// 2. Wydrukuj tylko katalog za pomocą `path.Split`
// 3. Odrzuć część z nazwą pliku
//
// OGRANICZENIE
// Użyj krótkiej deklaracji
//
// OCZEKIWANY WYNIK
// secret/
// ------------------------------------------------ ---------

import (
	"fmt"
	"path"
)

func main() {
	katalog, _:= path.Split("secret/file.txt")

	fmt.Println(katalog)
}
