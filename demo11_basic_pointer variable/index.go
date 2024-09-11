package main

import "fmt"

func main() {
	msg := "some message"
	var msgPointer *string = &msg

	/*fmt.Println(msg) //shows the value of msg (some message)
	fmt.Println(msgPointer) //shows address of the msg (0x14000010050)
	fmt.Println(*msgPointer) //shows the value of msg (some message)*/

	changeMessage(&msg, "new message 1") //calling a function to change the value of variable msg
	fmt.Println(msg)

	changeMessage(msgPointer, "new message 2") //calling a function to change the value of variable msg
	fmt.Println(msg)

	changeMessage(msgPointer, "new message 3") //calling a function to change the value of variable msg
	fmt.Println(*msgPointer)
}


func changeMessage(aPointer *string, newMessage string){
	*aPointer = newMessage
}