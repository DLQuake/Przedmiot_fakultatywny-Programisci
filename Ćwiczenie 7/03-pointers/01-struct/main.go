package main

import "fmt"

// ------------------------------------------------ ---------
// ĆWICZENIE:
//
// Utwórz dowolną sktrukturę z jednym lub kilkoma polami
//
// 1. Napisz funkcję która przyjmuje tę strukturę jako parametr oraz modyfikuje jej wartości. Zmiany powinny by permanentne
// Wypisz wynik

type structName struct {
	field1 string
	field2 int
	field3 float64
}

func main() {
	structName := &structName{field1: "value1", field2: 2, field3: 3.14}
	fmt.Printf("1) %+v\n", structName)

	function(structName)

	fmt.Printf("3) %+v\n", structName)
}

func function(structName *structName) {
	structName.field1 = "value2"
	structName.field2 = 3
	structName.field3 = 6.28
}
