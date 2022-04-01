package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ----------------------------------------------------------
// ĆWICZENIE:
//
// Użyj poprzedniego ćwiczenia
//
// Twoim zadaniem jest wydrukowanie średnich powierzchni, sypialni, łazienek i cen.
//
//
// OCZEKIWANY WYNIK
//
//  Location       Size           Beds           Baths          Price
//  ===========================================================================
//  New York       100            2              1              100000
//  New York       150            3              2              200000
//  Paris          200            4              3              400000
//  Istanbul       500            10             5              1000000
//
//  ===========================================================================
//                 237.50         4.75           2.75           425000.00
//
// ---------------------------------------------------------

func main() {
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`

		separator = ","
	)

	var (
		location []string
		size     []int
		beds     []int
		baths    []int
		price    []int
	)


	for _, line := range strings.Split(data, "\n") {
		columns := strings.Split(line, separator)

		location = append(location, columns[0])
		size = append(size, atoi(columns[1]))
		beds = append(beds, atoi(columns[2]))
		baths = append(baths, atoi(columns[3]))
		price = append(price, atoi(columns[4]))
	}

	headerColumns := strings.Split(header, separator)
	fmt.Printf("%-15v", headerColumns[0])
	fmt.Printf("%-15v", headerColumns[1])
	fmt.Printf("%-15v", headerColumns[2])
	fmt.Printf("%-15v", headerColumns[3])
	fmt.Printf("%-15v\n", headerColumns[4])

	fmt.Printf("===========================================================================\n")


	for i := 0; i < len(location); i++ {
		fmt.Printf("%-15v", location[i])
		fmt.Printf("%-15v", size[i])
		fmt.Printf("%-15v", beds[i])
		fmt.Printf("%-15v", baths[i])
		fmt.Printf("%-15v\n", price[i])
	}

	fmt.Printf("\n===========================================================================\n")

	fmt.Printf("%-15v", "")
	fmt.Printf("%-15.2f", average(size))
	fmt.Printf("%-15.2f", average(beds))
	fmt.Printf("%-15.2f", average(baths))
	fmt.Printf("%-15.2f\n", average(price))
}

func atoi (s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func average (numbers []int) float64 {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return float64(sum) / float64(len(numbers))
}