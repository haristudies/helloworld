package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// Generate a random number between 1 and 10
	randomNumber := rand.Intn(10) + 1
	fmt.Println("random number is ", randomNumber)

	// Example 1: For loop to iterate through numbers 1 to 5
	fmt.Println("Example 1: For Loop")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
	for i := 0; i <= randomNumber; i++ {
		fmt.Println("using random number to print", i)
	}
	// Example 2: If/else statement to check if the random number is even or odd
	fmt.Println("\nExample 2: If/Else Statement")
	if randomNumber%2 == 0 {
		fmt.Printf("%d is even\n", randomNumber)
	} else {
		fmt.Printf("%d is odd\n", randomNumber)
	}

	// Example 3: Switch statement to categorize the random number
	fmt.Println("\nExample 3: Switch Statement")
	switch {
	case randomNumber < 5:
		fmt.Printf("%d is less than 5\n", randomNumber)
	case randomNumber >= 5 && randomNumber <= 7:
		fmt.Printf("%d is between 5 and 7\n", randomNumber)
	default:
		fmt.Printf("%d is greater than 7\n", randomNumber)
	}

}
