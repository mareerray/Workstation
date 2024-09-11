package main

import (
	"fmt"
)

func main() {

	//fmt.Println(sum(2, 3))

	/*const a, b = 2, 3
	//fmt.Println(sum(a, b))
	fmt.Printf("%d+%d = %d \n", a, b, sum(a, b))*/

	/*var x, y int = swap(a, b)
	//or x,y := swap(a,b)
	fmt.Printf("%d,%d => %d,%d", a, b, x, y)
	fmt.Print("\n")*/

	x, y := swap2(10, 20)
	fmt.Printf("%d,%d => %d,%d", 10, 20, x, y)
	fmt.Print("\n")

}

func swap2(a, b int) (x, y int) {
	y = a
	x = b
	return
}
