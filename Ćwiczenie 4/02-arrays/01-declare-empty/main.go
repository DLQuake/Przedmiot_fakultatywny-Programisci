package main

import "fmt"

// ------------------------------------------------ ---------
// ĆWICZENIE: Zadeklaruj puste tablice
//
// 1. Zadeklaruj i wydrukuj następujące tablice wraz z ich typami:
//
// 1. Imiona Twoich trzech najlepszych przyjaciół (tablica napisów)
//
// 2. Odległości do pięciu różnych lokalizacji (tablica odległości)
//
// 3. Bufor danych o pojemności pięciu bajtów (tablica danych)
//
// 4. Kursy walut tylko dla jednej waluty
//
// 5. Status Up/Down czterech różnych serwerów WWW (tablica Alives)
//
// 6. Tablica bajtów, która nie zajmuje miejsca w pamięci (tablica pojemności zero)
//
// 2. Wydrukuj tylko typy tych samych tablic.
//
// 3. Wydrukuj tylko elementy tych samych tablic.
//
// WSKAZÓWKA
// Podczas drukowania elementów tablicy możesz użyć zwykłych czasowników Printf.
//
//   Na przykład:
// Podczas drukowania tablicy ciągów możesz jak zwykle użyć czasownika "%q".
//
// OCZEKIWANY WYNIK
//  names    : [3]string{"", "", ""}
//  distances: [5]int{0, 0, 0, 0, 0}
//  data     : [5]uint8{0x0, 0x0, 0x0, 0x0, 0x0}
//  ratios   : [1]float64{0}
//  alives   : [4]bool{false, false, false, false}
//  zero     : [0]uint8{}
//
//  names    : [3]string
//  distances: [5]int
//  data     : [5]uint8
//  ratios   : [1]float64
//  alives   : [4]bool
//  zero     : [0]uint8
//
//  names    : ["" "" ""]
//  distances: [0 0 0 0 0]
//  data     : [0 0 0 0 0]
//  ratios   : [0.00]
//  alives   : [false false false false]
//  zero     : []
// ---------------------------------------------------------

func main() {
	names := [3]string{"Oliwia", "Wiktor", "Damian"}
	distances := [5]int{0, 0, 0, 0, 0}
	data := [5]uint8{0, 0, 0, 0, 0}
	ratios := [1]float64{0.00}
	alives := [4]bool{false, false, false, false}
	zero := [0]uint8{}

	fmt.Printf("names    : %v\n",names)
	fmt.Printf("distances: %v\n",distances)
	fmt.Printf("data     : %v\n",data)
	fmt.Printf("ratios   : %v\n",ratios)
	fmt.Printf("alives   : %v\n",alives)
	fmt.Printf("zero     : %v\n",zero)

	fmt.Println()

	fmt.Printf("names    : %T\n", names)
	fmt.Printf("distances: %T\n", distances)
	fmt.Printf("data     : %T\n", data)
	fmt.Printf("ratios   : %T\n", ratios)
	fmt.Printf("alives   : %T\n", alives)
	fmt.Printf("zero     : %T\n", zero)

	fmt.Println()

	fmt.Printf("names    : %q\n", names)
	fmt.Printf("distances: %d\n", distances)
	fmt.Printf("data     : %d\n", data)
	fmt.Printf("ratios   : %.2f\n", ratios)
	fmt.Printf("alives   : %t\n", alives)
	fmt.Printf("zero     : %x\n", zero)
}
