package main

import "fmt"

type booking struct {
	firstName string
	lastName  string
	emailID   string
	tickets   uint
}

var conferenceTickets uint = 50
var remainingTickets uint = 50
var bookedTickets = make([]booking, 0)

func greetUsers() {
	fmt.Println("")
	fmt.Println("Welcome to Conference booking application")
	fmt.Printf("We have a total of %v tickets and %v tickets are still available for booking.\n", conferenceTickets, remainingTickets)
}

func bookTickets() bool {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Printf("Number of tickets remaining: %v\n", remainingTickets)

	fmt.Println("\nPlease enter the following details: ")
	fmt.Println("First name: ")
	fmt.Scan(&firstName)

	fmt.Println("Last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Email: ")
	fmt.Scan(&email)

	fmt.Println("Number of tickets you need: ")
	fmt.Scan(&userTickets)

	if userTickets > remainingTickets {
		fmt.Println("")
		fmt.Printf("Couldn't complete booking as only %v tickets are available\n", remainingTickets)
		return false
	}
	createBooking(firstName, lastName, email, userTickets)
	return true
}

func createBooking(firstName string, lastName string, email string, userTickets uint) {
	remainingTickets = remainingTickets - userTickets
	book := booking{firstName: firstName, lastName: lastName, emailID: email, tickets: userTickets}
	bookedTickets = append(bookedTickets, book)
	fmt.Printf("\nCongratulations %v %v!! You have successfully booked %v ticket(s)\nYou will receive the confirmation on %v.", firstName, lastName, userTickets, email)
}

func viewBookings() {
	fmt.Println("")
	for _, booking := range bookedTickets {
		fmt.Printf("%v %v %v %v\n", booking.firstName, booking.lastName, booking.emailID, booking.tickets)
	}
}

func main() {
	var userChoice uint
	var status bool = true
	greetUsers()
	for remainingTickets > 0 {
		fmt.Println("Enter your choice: ")
		fmt.Println("1. Book tickets")
		fmt.Println("2. View all bookings")
		fmt.Println("3. Exit")

		fmt.Scan(&userChoice)

		switch userChoice {
		case 1:
			status = bookTickets()

		case 2:
			viewBookings()

		default:
			status = false

		}
		if !status {
			break
		}
		fmt.Println("\nDo you want to continue in the application?")
		fmt.Println("1. Continue")
		fmt.Println("2. Exit")
		var continuity uint
		fmt.Scan(&continuity)
		switch continuity {
		case 1:
			status = true

		case 2:
			status = false

		default:
			status = false

		}
		if !status {
			break
		}

	}
	if remainingTickets == 0 {
		fmt.Println("Sold out!!")
	}

}
