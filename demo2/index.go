package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hey")
	time.Sleep(1 * time.Second) //after printing "hey", print the word "finish" with a delay time of 1 second
	fmt.Println("Finished")
}
