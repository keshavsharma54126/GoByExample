package main

import (
	"errors"
	"fmt"
)

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

var ErrorOutofTea = fmt.Errorf("no more tea available")
var ErrPower = fmt.Errorf("can't boil water")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrorOutofTea
	} else if arg == 4 {
		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}
func sendsmsTocouple(msgToCustomer, msgToSpuse string) (float64, error) {
	cost, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0.0, err
	}
	constForSpuse, err := sendSMS(msgToSpuse)
	if err != nil {
		return 0.0, err
	}
	return cost + constForSpuse, nil
}
func sendSMS(message string) (float64, error) {
	const maxTextLen = 25
	const costPerChat = 0.0002
	if len(message) > maxTextLen {
		return 0.0, fmt.Errorf("can't send text over %v characters", maxTextLen)
	}
	return costPerChat * float64(len(message)), nil
}

// the error interface 
type divideError struct{
	divident float64
}
func (de divideError) Error() string {
	return fmt.Sprintf("can not divide %v by zero",de.divident)
}


func main() {
	totalcost, err := sendsmsTocouple("asdlfhoiwerhowerwer", ";lkjlhsadflkjhsadf")
	if err != nil {
		fmt.Println("error occured", err)
	} else {
		fmt.Println(totalcost)
	}
	for _, i := range []int{7, 42} {
		if r, e := f(i); e != nil {
			fmt.Println("f failed: ", e)
		} else {
			fmt.Println("f worked: ", r)
		}
	}
	for i := range 5 {
		if err := makeTea(i); err != nil {
			if errors.Is(err, ErrorOutofTea) {
				fmt.Println("We should buy new tea!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("now it is dark")
			} else {
				fmt.Println("unknown error: ", err)
			}
			continue
		}
		fmt.Println("Tea is ready")
	}
}
