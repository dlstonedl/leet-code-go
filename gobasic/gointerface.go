package main

import "fmt"

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
	fmt.Println("nokiaPhone call")
}

type Iphone struct {
}

func (iPhone Iphone) call() {
	fmt.Println("iPhone call")
}

func main() {
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()

	phone = new(Iphone)
	phone.call()
}
