package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	fmt.Printf("working with first repo %s", "Hari prasad")
	//product := mul(2, 3)
	//fmt.Printf("product of two numbers %v", product)
	//sum := addition(3, 4)
	//fmt.Printf("sum of two numbers %v", sum)
	var a nums
	a.firstNum = 10
	a.secondNum = 20
	fmt.Println(a)
	fmt.Println(a.addition())
	fmt.Println(a.mul())
	//region with functions
	//var a = &nums{firstNum: 10, secondNum: 20} //result := a.addition()
	//fmt.Println(result)
	//find the sum as we have receiver on the nums struct
	//fmt.Printf("product of the two numbers is %v and sum is %v", a.mul(), a.addition())
}
