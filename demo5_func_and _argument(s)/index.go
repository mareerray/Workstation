package main

import "fmt"

func main() {
	fn1() //how to call function-fn1 without passing arguments
	fn2("Good Maire!")
	fn3("Good Morning!", 22)

	fmt.Println(sum(2, 3))

}

func fn1() {
	fmt.Println("CodeM")
}

func fn2(msg string) { // msg=name of variable, string=type of variable (one argument)
	fmt.Println(msg)
}

func fn3(title string, version int) { // passing more than one argument
	fmt.Print(title) //do not print on the next line
	fmt.Println(version)
}

func sum(a int, b int) int {
	return a + b
}
