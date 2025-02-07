package main

import (
	"fmt"
	"sync"
	"time"
)

type email struct {
	sender  string
	message string
	date    time.Time
}

func sendEmail(message string, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: '%s'\n", message)
	}()
	fmt.Printf("Email send:'%s'\n", message)
}

func filterOldEmails(emails []email, wg *sync.WaitGroup) {
	isOldChan := make(chan bool)

	go func() {
		defer wg.Done()
		for _, e := range emails {
			if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
				isOldChan <- true
			}
			isOldChan <- false
		}
	}()

	isOld := <-isOldChan
	fmt.Println("email 1 is old", isOld)
	isOld = <-isOldChan
	fmt.Println("email 2 is old", isOld)
}

// buffered channels
func addEmailsToQueue(emails []string) chan string {
	emailsToSend := make(chan string, len(emails))
	for _, email := range emails {
		emailsToSend <- email
	}
	return emailsToSend
}

func sendEmails(batchSize int, ch chan string) {
	for i := 0; i < batchSize; i++ {
		email := <-ch
		fmt.Println("Sending email", email)
	}
}

// closing the channels and chekcing if channels are closed or not
func countReports(numSentCh chan int) int {
	total := 0
	for {
		numSent, ok := <-numSentCh
		if !ok {
			break
		}
		total += numSent
	}

	return total
}

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
		fmt.Println("Send batch of reports", numReports)
		time.Sleep(time.Microsecond * 100)
	}
	close(ch)
}

//range keyword with channels
func concurrentFib(n int){
	chInts := make(chan int)
	go func(){
		fibonacci(29,chInts)
	}()
	for vals:= range chInts{
		fmt.Println(vals)
	}
}

func fibonacci(n int,ch chan int){
	x,y:=0,1
	for i:=0;i<n;i++{
		ch<- x 
		x,y = y,x+y 
		time.Sleep(time.Microsecond*10)
	}
	close(ch)
}

//select with channels 
func logMessages(chEmails,chSms chan string){
	for{
		select {
		case email,ok:= <-chEmails:
			if !ok{
				return 
			}
			logEmail(email)
		case sms,ok:= <- chSms:
			if !ok{
				return 
			}
			logSms(sms)
		}
	}
}

func logSms(sms string){
	fmt.Println("SMS",sms)
}
func logEmail(email string){
	fmt.Println("EMAIL",email)
}

func main() {
	fmt.Println("go routines tutorial")
	var wg sync.WaitGroup
	wg.Add(2)
	sendEmail("hello there ruchi", &wg)
	wg.Wait()
}
