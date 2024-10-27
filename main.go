package main

import "fmt"

func main() {
	fmt.Println("hello world")

	var confName = "GO CONFERENCE"
	const confTickets = 50
	var remTickets = 50

	// fmt.Println("welcome to ", confName, " booking app")
	// fmt.Println("Total: ", confTickets, " tickets, remaining: ", remTickets)
	fmt.Printf("Welcome to %v booking app\n", confName)
	fmt.Printf("Total: %v tickets, remaining: %v\n", confTickets, remTickets)

	// var userName
	// userName = "Tom"

	var bookings = [50]string{"PRabski", "gibran"}
	bookings[2] = "fufu"

	var var1 string
	// var var2 int
	// var3 := "value"
	fmt.Scan(&var1)
	fmt.Println(var1)
	fmt.Println(&var1)
}
