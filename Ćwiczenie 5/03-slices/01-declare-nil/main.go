package main

import "fmt"

// ------------------------------------------------ ---------
// ĆWICZENIE:
//
// 1. Zadeklaruj następujące wycinki jako zerowe:
//
// 1. Imiona znajomych (wycinek names)
//
// 2. Odległości do lokalizacji (wycinek distances)
//
// 3. Bufor danych (wycinek data)
//
// 4. Kursy walut (wycinek ratios)
//
// 5. Stan serwerów WWW up/down (wycinek alives)
//
//
// 2. Wydrukuj ich typ, długość i czy są równe wartości zero, czy nie.
//
//
// OCZEKIWANY WYNIK
//  names    : []string 0 true
//  distances: []int 0 true
//  data     : []uint8 0 true
//  ratios   : []float64 0 true
//  alives   : []bool 0 true
// ---------------------------------------------------------

func main() {
	var names []string
	var distances []int
	var data []uint8
	var ratios []float64
	var alives []bool

	fmt.Printf("names    : %T %d %t\n", names, len(names), names == nil)
	fmt.Printf("distances: %T %d %t\n", distances, len(distances), distances == nil)
	fmt.Printf("data     : %T %d %t\n", data, len(data), data == nil)
	fmt.Printf("ratios   : %T %d %t\n", ratios, len(ratios), ratios == nil)
	fmt.Printf("alives   : %T %d %t\n", alives, len(alives), alives == nil)
}
