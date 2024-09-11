package main

import "fmt"

/*func main() {
	//someValue := 15

	/* calling a function that prints a value under this condition (simple if-else)
	if someValue == 10 {
		fmt.Println(someValue, "== 10")
	} else {
		fmt.Println(someValue, "!= 10")
	}*/

	/* calling a function that has more condition
	if someValue > 10 && someValue < 20 { 
		fmt.Println(someValue, "is between 10 and 20")
	} else {
		fmt.Println(someValue, "is NOT between 10 and 20")
	}*/

	/* ??????
	if result := doSomething(); result == "ok" {
		fmt.Println("ok")
	}else{
		fmt.Println("Not ok")
	}
}

func doSomething() string {
	return "tada"	
}*/

//Switch-case is a control flow statement in Go that allows you to execute 
//different blocks of code based on the value of an expression.

func main()  {
	fnSwitchCase()
	
}

func fnSwitchCase()  {
	index := 3
	switch index {
	case 0:
		fmt.Println("0")
		//break (There's no need for a break statement at the end of each case; 
		//Go automatically falls through to the next case)
	case 1:
		fmt.Println("1")
		//break
	case 2:
		fmt.Println("2")
		//break
	default:
		fmt.Println("something else")
	}
	
}