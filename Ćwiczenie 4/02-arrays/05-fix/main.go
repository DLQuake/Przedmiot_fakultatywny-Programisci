package main

import "fmt"

// ------------------------------------------------ ---------
// ĆWICZENIE: Napraw
//
// 1. Odkomentuj kod
//
// 2. I napraw problemy
//
// 3. BONUS: Uprość kod
// ------------------------------------------------ ---------

func main() {
	// var names [3]string = [3]string{
	// 	"Einstein" "Shepard"
	// 	"Tesla"
	// }

	// var books [5]string = [5]string{
	// 	"Kafka's Revenge",
	// 	"Stay Golden",
	// 	"",
	// 	"",
	// 	""
	// }

	// fmt.Printf("%q\n", names)
	// fmt.Printf("%q\n", books)

	names := [3]string{"Einstein", "Shepard", "Tesla"}
	books := [5]string{"Kafka's Revenge", "Stay Golden", "", "", ""}
	fmt.Printf("%q\n %q\n", names, books)
}
