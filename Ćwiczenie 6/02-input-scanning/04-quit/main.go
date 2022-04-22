package main

import (
	"bufio"
	"fmt"
	"os"
)

// ------------------------------------------------ ---------
// ĆWICZENIE:
//
// Utwórz program, który kończy działanie, gdy użytkownik wpisze
// to samo słowo dwa razy.
//
//
// OGRANICZENIE
//
// Program powinien działać bez uwzględniania wielkości liter.
//
//
// OCZEKIWANY WYNIK
//
//  go run main.go
//
//   hey
//   HEY
//   TWICE!
//
//  go run main.go
//
//   hey
//   hi
//   HEY
//   TWICE!
// ---------------------------------------------------------

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Split(bufio.ScanWords)

	words := make(map[string]bool)

	for scanner.Scan() {

		word := scanner.Text()

		if words[word] {
			fmt.Println("TWICE!")
			break
		}

		words[word] = true
	}

}
