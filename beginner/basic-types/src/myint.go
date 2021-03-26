package myint

import "fmt"

func myint() {
	fmt.Println("Hello World")

	var myint int8
	for i := 0; i < 129; i++ {
		myint += 1
	}
	fmt.Println(myint) // prints out -127
}
