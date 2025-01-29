package main

import "fmt"

func bulkSend(numMessages int) float64 {
	cost := 0.0
	for i := 0; i < numMessages; i++ {
		cost += 1.0 + (0.01 * float64(i))
	}
	return cost
}

func maxMesages(thresh float64) int{
	totalCost:=0.0
	for i:=0; ;i++{
		totalCost +=1.0+ (0.01+float64(i))
		if(totalCost > thresh){
			return i
		}
	}
}

func test(numMessages int) {
	fmt.Printf("Sending %v messages\n", numMessages)
	cost := bulkSend(numMessages)
	fmt.Printf("Bulk send complete! Cost = %2f\n ", cost)
}

//go does not have while loop for loop is used as a while loop 

func main() {
	test(10)
	i := 1
	for i <= 6 {
		fmt.Println(i)
		i++
	}

	for j := 0; j < 34; j++ {
		fmt.Println(j)
	}

	for i := range 3 {
		fmt.Println("range", i)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := range 43 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
