package main

// ------------------------------------------------ ---------
// ĆWICZENIE:  długość stałej
// KROKI:
// 1. Zadeklaruj stałą o nazwie `home`
// 2. Zainicjuj ją jako literał ciągu „Milky Way Galaxy”
//
// 3. Zadeklaruj inną stałą o nazwie `length`
// 4. Zainicjuj go za pomocą wbudowanej funkcji `len`.
//
// 5. Wydrukuj poniższy komunikat, używając stałych, które
// zadeklarowałeś.
//
// OGRANICZENIE:
// Użyj Printf.
// Wydrukuj stałą `home` używając czasownika w cudzysłowie.
//
// OCZEKIWANY WYNIK
// There are 16 characters inside "Milky Way Galaxy"
import "fmt"

func main() {
	const home string = "\"Milky Way Galaxy\""
	const length = len(home)

	fmt.Printf("There are %d characters inside %s",length, home)
}
