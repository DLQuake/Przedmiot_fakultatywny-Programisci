package main

import "fmt"
import f "fmt"
import fm "fmt"
import qwerty "fmt"

func main() {
	fmt.Println("hello")
	f.Println("hey")
	fm.Println("hi")
	qwerty.Println("qwerty")

}

func init()  {
	fmt.Println("JESTEM PIERWSZY")
}