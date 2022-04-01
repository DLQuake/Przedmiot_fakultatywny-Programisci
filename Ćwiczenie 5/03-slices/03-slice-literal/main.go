package main

import "fmt"

// ------------------------------------------------ ---------
// ĆWICZENIE:
//
// 1. Przypisz następujące dane za pomocą literałów wycinków do wycinków, które
// zadeklarowałeś w pierwszym ćwiczeniu.
//
// 1. Imiona twoich trzech najlepszych przyjaciół
//
// 2. Odległości do pięciu różnych lokalizacji
//
// 3. Pięć bajtów danych
//
// 4. Dwa kursy wymiany walut
//
// 5. Stan czterech różnych serwerów WWW
//
// 2. Wydrukuj ich typ, długość i czy są równe wartości zero, czy nie
//
//
// OCZEKIWANY WYNIK
//  names    : []string 3 false
//  distances: []int 5 false
//  data     : []uint8 5 false
//  ratios   : []float64 2 false
//  alives   : []bool 4 false
// ---------------------------------------------------------

func main() {
	var (
		names     []string  // The names of your friends
		distances []int     // The distances
		data      []byte    // A data buffer
		ratios    []float64 // Currency exchange ratios
		alives    []bool    // Up/Down status of web servers
	)

	names = []string{"Einstein", "Tesla", "Shepard"}
	distances = []int{50, 40, 75, 30, 125}
	data = []byte{72, 69, 76, 76, 79}
	ratios = []float64{3.14, 2.56}
	alives = []bool{true, false, true, false}

	fmt.Printf("names    : %T %d %t\n", names, len(names), names == nil)
	fmt.Printf("distances: %T %d %t\n", distances, len(distances), distances == nil)
	fmt.Printf("data     : %T %d %t\n", data, len(data), data == nil)
	fmt.Printf("ratios   : %T %d %t\n", ratios, len(ratios), ratios == nil)
	fmt.Printf("alives   : %T %d %t\n", alives, len(alives), alives == nil)
}
