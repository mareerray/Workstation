package main

import (
	"fmt"
)

func main() {

	//fmt.Println(sum(2, 3))

	const a, b = 2, 3
	//fmt.Println(sum(a, b))
	fmt.Printf("%d+%d = %d \n", a, b, sum(a, b))

	var x, y int = swap(a, b)
	//or x,y := swap(a,b)
	fmt.Printf("%d,%d => %d,%d", a, b, x, y)
	fmt.Print("\n")

}

func sum(a int, b int) int { //a function that returns one result
	return a + b
}

func swap(a int, b int) (int, int) { //a function that returns more than one results
	return b, a
}
