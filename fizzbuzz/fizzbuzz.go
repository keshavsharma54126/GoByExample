package main

import (
	"fmt"
)

func fizzbuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("\n", "fizzbuzz")

		} else if i%3 == 0 {
			fmt.Println("\n", "fizz")
		} else if i%5 == 0 {
			fmt.Println("\n", "buzz")
		} else {
			fmt.Println("\n", i)
		}

	}
}

func printPrimes(max int) {
	for n := 2; n < max+1; n++ {
		if n == 2 {
			fmt.Println(n)
			continue
		}
		if n%2 == 0 {
			continue
		}
		isPrime := true
		for i := 3; i*i < n+1; i++ {
			if n%i == 0 {
				isPrime = false
				break
			}
		}
		if !isPrime {
			continue
		}
		fmt.Println(n)
	}
}

func test(max int) {
	printPrimes(max)
}

func main() {
	fizzbuzz()
	test(100)
}
