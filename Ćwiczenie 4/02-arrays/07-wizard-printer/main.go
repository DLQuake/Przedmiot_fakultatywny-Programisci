package main

import "fmt"

// ------------------------------------------------ ---------
// ĆWICZENIE:
//
// W tym ćwiczeniu Twoim celem jest pokazanie kilku znanych naukowców
// w ładnym wydruku.
//
// 1. Utwórz tablicę wielowymiarową
// 2. W każdej tablicy wewnętrznej zapisz imię i nazwisko naukowca oraz jego/jej
//      nick
// 3. Wydrukuj  informacje w ładnej tabeli za pomocą pętli.
//
// EXPECTED OUTPUT
//   First Name      Last Name       Nickname
//   ==================================================
//   Albert          Einstein        time
//   Isaac           Newton          apple
//   Stephen         Hawking         blackhole
//   Marie           Curie           radium
//   Charles         Darwin          fittest
// ---------------------------------------------------------

func main() {
	tablica := [5][3]string{
		{"Albert", "Einstein", "time"},
		{"Isaac", "Newton", "apple"},
		{"Stephen", "Hawking", "blackhole"},
		{"Marie", "Curie", "radium"},
		{"Charles", "Darwin", "fittest"}}

	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("	%-10s\t", tablica[i][j])
		}
		fmt.Printf("\n")
	}
}

