package main

import (
	"fmt"
	"strings"
)

//two main reasons we use pointers in go
// to give a reference to a value to some function
//to avoid memory copying for micro optimizations

type Message struct {
	Recipient string
	Text      string
}

func sendMessage(m Message) {
	fmt.Printf("To: %v\n", &m.Recipient)
	fmt.Printf("Message: %v\n", &m.Text)
	fmt.Printf("To: %v\n", m.Recipient)
	fmt.Printf("Message: %v\n", m.Text)
}

func removeProfanity(message *string) {
	if message==nil{
		return 
	}
	*message = strings.ReplaceAll(*message, "dang", "****")
	*message = strings.ReplaceAll(*message, "shoot", "****")
	*message = strings.ReplaceAll(*message, "heck", "****")
}

type car struct{
	color string 
}

//pointer reciever 

func (c *car) setColor(color string){
	c.color = color
}

//non pointer reciever
func (c car) setcolor(color string){
	c.color = color
}



func main() {

	//var p *int;//nil or null pointer

	mystring := "hello"
	myStringPtr := &mystring
	fmt.Println(*myStringPtr)
	m := Message{
		Recipient: "Keshav",
		Text:      "Hi buddy hoow are you",
	}
	sendMessage(Message{
		Recipient: "Kesahv",
		Text:      "hi buddy how are you doing",
	})
	sendMessage(m)
	messagestring := "hello, dang that is shoot and it is heck"
	removeProfanity(&messagestring)
	fmt.Println(messagestring)

	c:=car{
		color:"blue",
	}
	c.setColor("pink")
	fmt.Println(c.color)
	c.setcolor("black")
	fmt.Println(c.color)
}
