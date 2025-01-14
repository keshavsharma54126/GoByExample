package main

import (
	"fmt"
)

func sum(nums []int) {
	len := len(nums)
	sum:=0
	for i:=0;i<len;i++{
		sum+=nums[i]
	}
	fmt.Println(sum)
}

func sum2(nums...int){

	sum:=0
	for i:=0;i<len(nums);i++{
		sum+=nums[i]
	}
	fmt.Println(sum)
}

func main() {
	fmt.Println("hello world", 2, 3, 4, 5, 6, 7, 8, 9, 10)//this is a variadic funciton
	nums:= []int{1,2,3}
	var nums2 = make([]int,10)
	for i:=0;i<3;i++{
		nums2[i]=i
	}
	fmt.Println(nums,nums2)
	sum(nums)
	sum2(nums2...)
}
