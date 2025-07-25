package main

import "fmt"

func ConditionalStatement () {

	age := 21
	if age > 21 {
		fmt.Println("you are eligible to merriage")
	} else if age < 21 {
		fmt.Println("you are not eligible to merriage. But you can love someone!")
	} else {
		fmt.Println("Enjoy your life without pain....🙂")
	}

}