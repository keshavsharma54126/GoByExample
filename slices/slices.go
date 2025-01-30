package main

import (
	"fmt"
	"slices"
)

type cost struct {
	day    int
	amount float64
}

func sum(nums ...float64) float64 {
	total := 0.0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	return total

}

func getCostByDay(costs []cost) []float64 {
	costsByDay := []float64{}
	for i := 0; i < len(costs); i++ {
		costi := costs[i]
		day := costi.day
		for day >= len(costsByDay) {
			costsByDay = append(costsByDay, 0.0)
		}
		costsByDay[day] += costi.amount

	}
	return costsByDay
}

func createMatrix(rows, cols int) [][]int {
	matrix := make([][]int, 0)
	for i := 0; i < rows; i++ {
		row := make([]int, cols)
		for j := 0; j < cols; j++ {
			row[j] = i * j
		}
		matrix = append(matrix, row)
	}
	return matrix
}

func printMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Println(matrix[i][j], " ")
		}
		fmt.Println()
	}
}

func main() {
	cost := []cost{{0, 4.0},
		{1, 2.1},
		{1, 3.1},
		{5, 2.5}}
	costsbyday := getCostByDay(cost)
	fmt.Println(costsbyday)
	matrix := createMatrix(10, 10)
	printMatrix(matrix)

	var s []string
	fmt.Println("unint: ", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp: ", s, "len: ", len(s), "cap: ", cap(s))

	var d = [2]int{1, 2}
	fmt.Println(d)
	var t = make([]string, 1, 10)
	fmt.Println(t)
	fmt.Println("len of t: ", len(t), "cap of t: ", cap(t))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd: ", s)

	c := make([]string, len(s))
	copy(c, s)

	fmt.Println("cpy: ", c)

	l := s[2:5]
	fmt.Println("sl1: ", l)

	l = s[:5]
	fmt.Println("sl2", 1)

	t2 := []string{"g", "h", "i"}
	fmt.Println("dc1: ", t2)

	t3 := []string{"g", "h", "i"}
	if slices.Equal(t2, t3) {
		fmt.Println("t2==t3")
	}

	twoD := make([][]int, 3)
	for i := 0; i < len(twoD); i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}
