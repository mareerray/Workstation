package majn

import "fmt"

func main() {

	//General Formatting Verbs
	var i = 15.5
	var txt = "Hello World!"

	fmt.Printf("%v\n", i)   //Prints the value in the default format
	fmt.Printf("%#v\n", i)  //Prints the value in Go-syntax format
	fmt.Printf("%v%%\n", i) //Prints the % sign
	fmt.Printf("%T\n", i)   //Prints the type of the value
	fmt.Print("\n")

	fmt.Printf("%v\n", txt)  //Prints the value in the default format
	fmt.Printf("%#v\n", txt) //Prints the value in Go-syntax format
	fmt.Printf("%T\n", txt)  //Prints the type of the value

	//Integer Formatting Verbs
	var j = 15

	fmt.Printf("%b\n", j)
	fmt.Printf("%d\n", j)
	fmt.Printf("%+d\n", j)
	fmt.Printf("%o\n", j)
	fmt.Printf("%O\n", j)
	fmt.Printf("%x\n", j)
	fmt.Printf("%X\n", j)
	fmt.Printf("%#x\n", j)
	fmt.Printf("%4d\n", j)
	fmt.Printf("%-4d\n", j)
	fmt.Printf("%04d\n", j)

}
