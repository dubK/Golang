package main

import (
	"fmt"
)

func main() {
	var val int
	for val != 5 {
		fmt.Print("[1]Add [2]Subtract [3]Multiply [4]Divide [5]Quit\n \n")
		fmt.Print("Your choice: ")
		fmt.Scan(&val)
		switch {
		case val == 1:
			tot := plus()
			fmt.Print("Result: ", tot, "\n\n")
		case val == 2:
			tot := minus()
			fmt.Print("Result: ", tot, "\n\n")
		case val == 3:
			tot := ganger()
			fmt.Print("Result: ", tot, "\n\n")
		case val == 4:
			tot := delat()
			fmt.Print("Result: ", tot, "\n\n")
		case val == 5:
			fmt.Print("You've chosen to quit, ba ba!\n \n")
		default:
			fmt.Println("You have entered an invalid number, try again!")
			fmt.Print("---------------------------------------------\n\n")
		}
	}

}
func plus() int {
	var a int
	var b int
	fmt.Print("You chose addition\n \n")
	fmt.Print("Number 1: ")
	fmt.Scan(&a)
	fmt.Print("Number 2: ")
	fmt.Scan(&b)
	return a + b
}
func minus() int {
	var a int
	var b int
	fmt.Print("You chose subtraction\n \n")
	fmt.Print("Number 1: ")
	fmt.Scan(&a)
	fmt.Print("Number 2: ")
	fmt.Scan(&b)
	return a - b
}
func ganger() int {
	var a int
	var b int
	fmt.Print("You chose multiplication\n \n")
	fmt.Print("Number 1: ")
	fmt.Scan(&a)
	fmt.Print("Number 2: ")
	fmt.Scan(&b)
	return a * b
}
func delat() float32 {
	var a int
	var b int
	fmt.Print("You chose division\n \n")
	fmt.Print("Number 1: ")
	fmt.Scan(&a)
	fmt.Print("Number 2: ")
	fmt.Scan(&b)
	return float32(a) / float32(b)
}

// fmt.Print("Number 1: ")
// fmt.Scan(&num1)
// fmt.Print("Number 2: ")
// fmt.Scan(&num2)
