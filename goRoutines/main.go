package main

import (
	"fmt"
	"time"
)

func sendEmail(message string) {
	go func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: '%s'\n", message)
	}()
	fmt.Printf("Email send:'%s'\n", message)
}

func main() {
	fmt.Println("go routines tutorial")
	sendEmail("hello there ruchi")
}
