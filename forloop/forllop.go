package main 

import "fmt"

func main(){

	i :=1
	for i<=6{
		fmt.Println(i)
		i++
	}

	for j := 0; j<34;j++ {
		fmt.Println(j)
	}

	for i:= range 3{
		fmt.Println("range",i)
	}

	for{
		fmt.Println("loop")
		break
	}

	for n:= range 43{
		if n%2 ==0{
			continue
		}
		fmt.Println(n)
	}
}