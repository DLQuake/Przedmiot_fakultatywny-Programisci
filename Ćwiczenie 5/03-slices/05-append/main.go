package main

import (
	"bytes"
	"fmt"
)

// ------------------------------------------------ ---------
// ĆWICZENIE:
//
// Postępuj zgodnie z instrukcjami zawartymi w poniższym kodzie.
//
// OCZEKIWANY WYNIK
// They are equal
//
// WSKAZÓWKI
// bytes.Equal funkcja pozwala porównać dwa bajty
// Sprawdź jego dokumentację: go doc bytes.Equal
// ------------------------------------------------ ---------

func main() {
	// 1. odkomentuj poniższy kod
	// png, header := []byte{'P', 'N', 'G'}, []byte{}
	png, header := []byte{'P', 'N', 'G'}, []byte{}

	// 2. dołącz elementy do header, aby był równy z wycinkiem png
	header = append(header, png...)
	// 3. porównaj wycinki za pomocą funkcji bytes.Equal
	pom := bytes.Equal(png, header)

	// 4. Wydrukuj informację czy wycinki są równe
	if pom {
		fmt.Println("They are equal")
	} else {
		fmt.Println("They are not equal")
	}
}
