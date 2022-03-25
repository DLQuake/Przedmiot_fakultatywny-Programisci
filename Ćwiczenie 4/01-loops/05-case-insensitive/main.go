package main

import (
	"fmt"
	"os"
	"strings"
)

// ------------------------------------------------ ---------
// ĆWICZENIE: Wyszukiwanie bez uwzględniania wielkości liter
//
// Zezwól na wyszukiwanie bez rozróżniania wielkości liter
//
// PRZYKŁAD
// Załóżmy, że użytkownik uruchamia program w następujący sposób:
//  go run main.go lAzY lub
//  go run main.go lazy
//
// We wszystkich powyższych przypadkach program powinien znaleźć
// słowo kluczowe "lazy".
// ------------------------------------------------ ---------
const corpus = "lazy cat jumps again and again and again"

func main() {
	words := strings.Fields(corpus)
	query := os.Args[1:]

queries :
	for _, q := range query {
		q = strings.ToLower((q))

	search:
		for i, w := range words {
			switch q {
			case "and", "or", "the":
				break search
			}

			if q == w {
				fmt.Printf("##{",i+1,"}: #{w}\n")
				continue queries
			}
		}

	}
}