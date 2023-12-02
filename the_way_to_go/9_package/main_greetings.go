package main

import (
	"fmt"
	"study_golang/the_way_to_go/9_package/greetings"
)

func main() {
	if greetings.IsAM() {
		fmt.Println("Good Day")
	} else if greetings.IsAfternoon() {
		fmt.Println("Good Afternoon")
	} else if greetings.IsEvening() {
		fmt.Println("Good Evening")
	} else {
		fmt.Println("Good Night")
	}

	greetings.Demo()
}
