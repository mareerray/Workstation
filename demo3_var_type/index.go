package main

import "fmt"

var count int = 0

func main() {
	fmt.Println("Begin")

	//Explicit Declaration
	var tmp1 int = 0
	var tmp2 string = "Hello"
	var tmp3 bool = true

	//Implicit Declaration
	tmp4 := 10
	tmp5 := "codeMaire"

	fmt.Println(tmp1)
	fmt.Println(tmp2)
	fmt.Println(tmp3)
	fmt.Println(tmp4)
	fmt.Println(tmp5)

	count++
	fmt.Println(count)
	run() //replace the repetition of count++,fmt.println(count) with run-function
	/*count++
	fmt.Println(count)
	count++
	fmt.Println(count)
	count++
	fmt.Println(count)*/
}

func run() {

	count++
	fmt.Println(count)
}
