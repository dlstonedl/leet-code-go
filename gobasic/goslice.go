package main

import "fmt"

func main() {
	var intSlice []int
	printSlice(intSlice)
	intSlice = append(intSlice, 1)
	printSlice(intSlice)

	intSlice1 := make([]int, 3)
	printSlice(intSlice1)

	intSlice2 := make([]int, 3, 5)
	printSlice(intSlice2)
	intSlice2 = append(intSlice2, 1, 2, 3)
	printSlice(intSlice2)
	intSlice2[0] = 1
	printSlice(intSlice2)

	var intSlice3 = make([]int, 3, 3)
	copy(intSlice3, intSlice2)
	printSlice(intSlice3)
}

func printSlice(x []int) {
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(x), cap(x), x)
}
