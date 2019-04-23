package main

import "fmt"

func practiceDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic(123)
	fmt.Println(4)
}

func main() {
	practiceDefer()
}
