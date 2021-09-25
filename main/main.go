package main

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func emojiiGreeting() {
	var name, emdigit1, emdigit2 string
	var rem3, rem5, rem7, age int
	fmt.Println("Please enter your name: ")
	fmt.Scan(&name)
	emoji.Println(":wave: Hello, " + name + "!")

	fmt.Println("Let me guess your age.")
	fmt.Println("Enter remainders of dividing your age by 3, 5 and 7.")
	fmt.Scan(&rem3, &rem5, &rem7)

	age = (rem3*70 + rem5*21 + rem7*15) % 105
	// Separate age in two digits
	digit1 := age / 10
	digit2 := age % 10

	switch digit1 {
	case 0:
		emdigit1 = ":zero:"
	case 1:
		emdigit1 = ":one:"
	case 2:
		emdigit1 = ":two:"
	case 3:
		emdigit1 = ":three:"
	case 4:
		emdigit1 = ":four:"
	case 5:
		emdigit1 = ":five:"
	case 6:
		emdigit1 = ":six:"
	case 7:
		emdigit1 = ":seven:"
	case 8:
		emdigit1 = ":eight:"
	case 9:
		emdigit1 = ":nine:"
	}

	switch digit2 {
	case 0:
		emdigit2 = ":zero:"
	case 1:
		emdigit2 = ":one:"
	case 2:
		emdigit2 = ":two:"
	case 3:
		emdigit2 = ":three:"
	case 4:
		emdigit2 = ":four:"
	case 5:
		emdigit2 = ":five:"
	case 6:
		emdigit2 = ":six:"
	case 7:
		emdigit2 = ":seven:"
	case 8:
		emdigit2 = ":eight:"
	case 9:
		emdigit2 = ":nine:"
	}

	emoji.Println(":calendar: You are " + emdigit1 + emdigit2 + " years old!")
}

func main() {
	emojiiGreeting()
}
