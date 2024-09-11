package main

import "fmt"

/*/calling a fucntion that prints 0-10
func main() {
	fnFor()

}

func fnFor(){ /
    for i := 0; i <= 10; i++ {
		fmt.Printf("Index %d\n", i)
		
	}
}*/

/*calling a fucntion that prints 1-5
func main()  {
	fnWhile()
}
func fnWhile()  {
	i := 0
	for i < 5 {
		i++
		fmt.Printf("Index %d\n", i)
	}
}*/

/*calling a fucntion that prints 0-4
func main()  {
	fnWhile()
}
func fnWhile()  {
	i := 0
	for i < 5 {
		fmt.Printf("Index %d\n", i)
		i++
	}
}*/

//calling a fucntion that prints digit starts from 0 until it reaches the condition, 
//then break out of the loop
func main()  {
	fnWhile()
}
func fnWhile()  {
	i := 0
	for true {
		if i > 7 {
			break
		}
		fmt.Printf("Index %d\n", i)
		i++
	}
}

