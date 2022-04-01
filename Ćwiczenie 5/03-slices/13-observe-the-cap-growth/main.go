package main

import "fmt"

// ------------------------------------------------ ---------
// ĆWICZENIE: Obserwuj wzrost wydajności
//
// Napisz program, który dołącza elementy do wycinka
// 10 milionów razy w pętli. Obserwuj, jak pojemność
// wycinka się zmienia.
//
//
// KROKI
//
// 1. Utwórz wycinek zerowy
//
// 2. Zapętl 1e7 razy
//
// 3. W każdej iteracji: Dołącz element do wycinka
//
// 4. Wydrukuj długość i pojemność wycinka TYLKO kiedy zmienia się jego pojemność.
//
// BONUS: Wydrukuj również tempo wzrostu pojemności.
//
//
// OCZEKIWANY WYNIK
//
//  len:0               cap:0               growth:NaN
//  len:1               cap:1               growth:+Inf
//  len:2               cap:2               growth:2.00
//  ...
//
// ---------------------------------------------------------

func main() {
	s := make([]int, 0)

	for i := 0; i < 1e7; i++ {
		s = append(s, i)

		if len(s) > 0 && len(s) != cap(s) {
			fmt.Printf("len:%d\tcap:%d\tgrowth:%.2f\n", len(s), cap(s), float64(len(s))/float64(cap(s)))
		}
	}
}
