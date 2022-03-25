package main

import "fmt"

// ------------------------------------------------ ---------
// ĆWICZENIE: Pobierz i ustaw elementy tablicy
//
// 1. Usuń wszystko oprócz deklaracji tablicy z poprzedniego ćwiczenia
//
// 2. Przypisz imiona swoich najlepszych przyjaciół do tablicy imion
//
// 3. Przypisz odległości do najbliższych miast do tablicy odległości
//
// 4. Przypisz dowolne bajty do tablicy danych
//
// 5. Przypisz wartość do tablicy wskaźników
//
// 6. Przypisz wartości true/false do tablic statusów serwerów
//
// 8. Spróbuj przypisać do tablicy zero i obserwuj błąd
//
// 9. Teraz użyj zwykłych instrukcji pętli dla każdej tablicy i wypisz je
//
// 10. Teraz użyj instrukcji pętli dla zakresu dla każdej tablicy i wypisz je
//
// OCZEKIWANY WYNIK
//
//   names
//   ====================
//   names[0]: "Einstein"
//   names[1]: "Tesla"
//   names[2]: "Shepard"
//
//   distances
//   ====================
//   distances[0]: 50
//   distances[1]: 40
//   distances[2]: 75
//   distances[3]: 30
//   distances[4]: 125
//
//   data
//   ====================
//   data[0]: 72
//   data[1]: 69
//   data[2]: 76
//   data[3]: 76
//   data[4]: 79
//
//   ratios
//   ====================
//   ratios[0]: 3.14
//
//   alives
//   ====================
//   alives[0]: true
//   alives[1]: false
//   alives[2]: true
//   alives[3]: false
//
//   zero
//   ====================

//
//   ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//   FOR RANGES
//   ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
//
//   names
//   ====================
//   names[0]: "Einstein"
//   names[1]: "Tesla"
//   names[2]: "Shepard"
//
//   distances
//   ====================
//   distances[0]: 50
//   distances[1]: 40
//   distances[2]: 75
//   distances[3]: 30
//   distances[4]: 125
//
//   data
//   ====================
//   data[0]: 72
//   data[1]: 69
//   data[2]: 76
//   data[3]: 76
//   data[4]: 79
//
//   ratios
//   ====================
//   ratios[0]: 3.14
//
//   alives
//   ====================
//   alives[0]: true
//   alives[1]: false
//   alives[2]: true
//   alives[3]: false
//
//   zero
//   ====================
//
// ---------------------------------------------------------

func main() {
	names := [3]string{"Einstein", "Tesla", "Shepard"}
	distances := [5]int{50, 40, 75, 30, 125}
	data := [5]uint8{72, 69, 76, 76, 79}
	ratios := [1]float64{3.14}
	alives := [4]bool{true, false, true, false}
	zero := [0]uint8{}

	fmt.Println("names")
	for i := 0; i < 3; i++ {
		fmt.Println("names[", i, "]:", names[i])
	}

	fmt.Println()

	fmt.Println("distances")
	for i := 0; i < 5; i++ {
		fmt.Println("distances[", i, "]:", distances[i])
	}

	fmt.Println()

	fmt.Println("data")
	for i := 0; i < 5; i++ {
		fmt.Println("data[", i, "]:", data[i])
	}

	fmt.Println()

	fmt.Println("ratios")
	for i := 0; i < 1; i++ {
		fmt.Println("ratios[", i, "]:", ratios[i])
	}

	fmt.Println()

	fmt.Println("alives")
	for i := 0; i < 4; i++ {
		fmt.Println("alives[", i, "]:", alives[i])
	}

	fmt.Println()

	fmt.Println("zero")
	for i := 0; i < 0; i++ {
		fmt.Println("zero[", i, "]:", zero[i])
	}

	fmt.Println()
	fmt.Println("FOR RANGES")
	fmt.Println()

	fmt.Println("names")
	for i , names := range names {
		fmt.Println("names[", i, "]:", names)
	}

	fmt.Println()

	fmt.Println("distances")
	for i , distances := range distances {
		fmt.Println("distances[", i, "]:", distances)
	}

	fmt.Println()

	fmt.Println("data")
	for i , data := range data {
		fmt.Println("data[", i, "]:", data)
	}

	fmt.Println()

	fmt.Println("ratios")
	for i , ratios := range ratios {
		fmt.Println("ratios[", i, "]:", ratios)
	}

	fmt.Println()

	fmt.Println("alives")
	for i , alives := range alives {
		fmt.Println("alives[", i, "]:", alives)
	}

	fmt.Println()

	fmt.Println("zero")
	for i , zero := range zero {
		fmt.Println("zero[", i, "]:", zero)
	}
}
