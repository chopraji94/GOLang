package main

import "fmt"

func main() {
	var confrenceName = "Go confrence"
	const confrenceTickets = 50
	var remainingTickets = 50
	fmt.Println("Welcome to ", confrenceName, "booking application")
	fmt.Println("We have total of", confrenceTickets, "and", remainingTickets, "are still available")

	//Updateed Code to printf
	fmt.Printf("welcome to %v booking application\n", confrenceName)
	fmt.Printf("We have tota of %v and %v are still remaining\n", confrenceTickets, remainingTickets)

	//More updated code
	var userName string
	var userTickets int

	userName = "Pranav"
	userTickets = 20

	fmt.Printf("User %v booked %v tickets", userName, userTickets)
}
