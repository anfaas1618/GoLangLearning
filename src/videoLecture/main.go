package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName = "go conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	fmt.Printf("Welcome to our %v bookings application\n", conferenceName)
	fmt.Printf("We have total of %v out of which %v re available", conferenceTickets, remainingTickets)
	fmt.Println("book your ticket now")
	fmt.Printf("")

	var bookings []string
	var firstname string
	var lastname string
	var email string
	var userTickets int

	for remainingTickets > 0 || len(bookings) < 50 {
		//ask for user input
		fmt.Println("enter your first firstname :")
		_, err := fmt.Scan(&firstname)
		if err != nil {
			return
		}
		fmt.Println("enter your last name")
		fmt.Scan(&lastname)
		fmt.Printf("enter your email")
		fmt.Scan(&email)
		fmt.Printf("enter number of tickets: ")
		fmt.Scan(&userTickets)
		if userTickets <= remainingTickets {
			bookings = append(bookings, firstname+" "+lastname)

			fmt.Printf("the whole array %v", bookings)

			remainingTickets = remainingTickets - userTickets

			fmt.Printf("thank you, %v %v for bookimh %v tickets recipt will be sent to %v \n",
				firstname, lastname, userTickets, email)
			var firstnames []string
			for _, book := range bookings {
				var names = strings.Fields(book) // split by space

				firstnames = append(firstnames, names[0])
			}
			fmt.Println(firstnames)
			fmt.Printf("%v tickets remaining", remainingTickets)
			if remainingTickets == 0 {
				fmt.Println("houseful")
			}

		} else {
			fmt.Printf("we have only %v tickets remaing\n", remainingTickets)
		}
	}
}

/*
wow this is great
*/
// isngge @hgello % #asdas &sdas todo as
