package main 

import(
	"fmt"
	"time"
)

func main(){

	i:=2
	fmt.Println("Write ",i," as ")
	switch i{
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday(){
	case time.Saturday , time.Sunday:
		fmt.Println("it is a weekend yay")
	default:
		fmt.Println("nahh it is a weekday")
	}

	//swithc without an expression is like an if,else flow 
	t:= time.Now()
	switch{
	case t.Hour()<12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it is after noon")
	}

	whatAmI:= func(i interface{}){
		switch t:=i.(type){
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm a integer")
		default:
			fmt.Println("Don't know type %T\n",t)
		}
	
	}
	whatAmI(true)
	whatAmI(34)
	whatAmI("hey")


}