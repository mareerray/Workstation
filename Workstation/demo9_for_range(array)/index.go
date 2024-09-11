package main

import "fmt"

func main() {
	courses := [] string {"Android", "iOS","React"}
	
	for index, item := range courses {
		//fmt.Printf("%d. %s\n", index, item) //this will print index 0,1,2,..
		fmt.Printf("%d. %s\n", index+1, item) //this will print index 1,2,3,..
	}

	for _, item := range courses {
		fmt.Printf("%s\n", item)
	}
}
