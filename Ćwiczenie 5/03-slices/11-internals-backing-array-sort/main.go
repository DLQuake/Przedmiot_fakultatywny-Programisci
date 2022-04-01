package main

import (
	"fmt"
	"sort"
)

// ------------------------------------------------ ---------
// ĆWICZENIE:
//
// 1. Sortuj tylko 3 środkowe elementy.
// OGRANICZENIE
//
// Nie sortuj ręcznie. Sortuj przy użyciu pakietu sortowania.
//
//
// OCZEKIWANY WYNIK
//
//  Original: [pacman mario tetris doom galaga frogger asteroids simcity metroid defender rayman tempest ultima]
//
//  Sorted  : [pacman mario tetris doom galaga asteroids frogger simcity metroid defender rayman tempest ultima]
//
//
// WSKAZÓWKA:
//
// Środkowe pozycje to: [frogger asteroids simcity]
//
// Po posortowaniu stają się: [asteroids frogger simcity]
//
// ------------------------------------------------ ---------

func main() {
	items := []string{
		"pacman", "mario", "tetris", "doom", "galaga", "frogger",
		"asteroids", "simcity", "metroid", "defender", "rayman",
		"tempest", "ultima",
	}

	fmt.Println("Original:", items)
	sort.Strings(items[5:7])
	fmt.Println()
	fmt.Println("Sorted  :", items)
}
