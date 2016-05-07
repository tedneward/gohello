package gohello

import "fmt"

var original = "Hello, CoDe world"
var Message = original

func PigLatinizer(msg string) (success bool, returnVal string) {
	success = true
	returnVal = msg + "way"
	return
}

func SayHello() {
	fmt.Printf("%v\n", Message)
}

func Reset() {
	Message = original
}

func DemoDefer() {
	fmt.Println("Entering DemoDefer")

	defer fmt.Println("In defer the first time")

	defer func() {
		fmt.Println("In defer the second time")
	}()

	fmt.Println("Exiting DemoDefer")
}
