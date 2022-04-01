package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"unsafe"
)

// ------------------------------------------------ ---------
// ĆWICZENIE: Obserwuj alokacje pamięci
//
// W tym ćwiczeniu Twoim celem jest obserwowanie alokacji pamięci (różnica pomiędzy między tablicami a wycinkami)
//
// Utworzysz, przypiszesz tablice i wycinki, a następnie wydrukujesz
// użycie pamięci twojego programu w każdym kroku.
//
// Postępuj zgodnie z instrukcjami zawartymi w kodzie.
//
//
// OCZEKIWANY WYNIK
//
// Zwróć uwagę, że numery wykorzystania pamięci mogą się różnić. Jednak rozmiar
// tablice i wycinki powinny być takie same na twoim własnym systemie
// (jeśli jesteś na komputerze 64-bitowym).
//
//
//  [initial memory usage]
//          > Memory Usage: 104 KB
//  [after declaring an array]
//          > Memory Usage: 78235 KB
//  [after copying the array]
//          > Memory Usage: 156365 KB
//  [inside passArray]
//          > Memory Usage: 234495 KB
//  [after slicings]
//          > Memory Usage: 234497 KB
//  [inside passSlice]
//          > Memory Usage: 234497 KB
//
//  Array's size : 80000000 bytes.
//  Array2's size: 80000000 bytes.
//  Slice1's size: 24 bytes.
//  Slice2's size: 24 bytes.
//  Slice3's size: 24 bytes.
//
//
// WSKAZÓWKI
//
// Zadeklarowałem kilka funkcji, które mogą Ci pomóc.
//
// funkcja raportu:
// — Drukuje użycie pamięci.
// - Po prostu wywołaj to z komunikatem, który pasuje do oczekiwanego wyniku.
//
// funkcja passArray:
// - Automatycznie drukuje użycie pamięci.
// — Akceptuje tablicę [size]int, więc możesz przekazać do niej swoją tablicę.
//
// funkcja passSlice:
// - Automatycznie drukuje użycie pamięci.
// - Akceptuje wycinek int, więc możesz przekazać go jako jeden z wycinków.
//
// ------------------------------------------------ ---------

const size = 1e7

func main() {
	// nie przejmuj się tym kodem.
	// zatrzymuje garbage collector: uniemożliwia czyszczenie pamięci.
	// zobacz link, jeśli jesteś ciekawy:
	// https://en.wikipedia.org/wiki/Garbage_collection_(computer_science)
	debug.SetGCPercent(-1)

	//początkowe użycie pamięci.
	report("initial memory usage")

	// 1. przydziel tablicę z 10 milionami elementów int
	// rozmiar tablicy będzie równy ~80MB
	// wskazówka: użyj powyższej stałej `size`.
	var array [size]int

	// 2. wydrukuj użycie pamięci (użyj funkcji raportu).
	report("after declaring an array")

	// 3. skopiuj tablicę do nowej tablicy.
	var array2 [size]int
	copy(array2[:], array[:])

	// 4. wydrukuj użycie pamięci
	report("after copying the array")

	// 5. przekazać tablicę do funkcji passArray
	passArray(array)

	// 6. przekonwertuj jedną z tablic na wycinek
	passSlice(array[:])

	// 7. wycinek powinien wskazywać na pierwsze 1000 elementów tablicy
	passSlice(array[:1000])

	// 8. kolejny wycinek powinien wskazywać pomiędzy 1000 a 10000 elementów tablicy
	passSlice(array[1000:10000])

	// 9. wydrukuj użycie pamięci
	report("after slicings")

	// 10. przekaż jeden z wycinków do funkcji passSlice
	passSlice(array[:])

	// 11. wydrukuj rozmiary tablic i wycinków
	// wskazówka: użyj funkcji unsafe.Sizeof
	// zobacz szczegóły: https://golang.org/pkg/unsafe/#Sizeof
	fmt.Printf("\nArray's size : %v bytes.\n", unsafe.Sizeof(array))
	fmt.Printf("Array2's size: %v bytes.\n", unsafe.Sizeof(array2))
	fmt.Printf("Slice1's size: %v bytes.\n", unsafe.Sizeof(array[:1000]))
	fmt.Printf("Slice2's size: %v bytes.\n", unsafe.Sizeof(array[1000:10000]))
	fmt.Printf("Slice3's size: %v bytes.\n", unsafe.Sizeof(array[10000:]))
}

func passArray(items [size]int) {
	items[0] = 100
	report("inside passArray")
}

func passSlice(items []int) {
	items[0] = 100
	report("inside passSlice")
}

func report(msg string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("[%s]\n", msg)
	fmt.Printf("\t> Memory Usage: %v KB\n", m.Alloc/1024)
}
