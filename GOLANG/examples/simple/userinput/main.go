package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Create a new reader that reads from standard input
	reader := bufio.NewReader(os.Stdin)

	// Prompt the user
	fmt.Println("Enter the rating for our pizza:")

	// Read input until a newline (\n) is encountered
	// The read returns two values: the input string and an error (or nil)
	input, _ := reader.ReadString('\n')

	// In this example we ignore the error using '_' (underscore)
	// But we'll show the proper way below

	// Print the input (includes the newline)
	fmt.Println("Thanks for rating", input)

	// Print the type of the variable (it's a string)
	fmt.Printf("Type of rating is %T\n", input)

	// --- Comma‑error syntax explained ---
	// If you want to handle errors properly:
	// input, err := reader.ReadString('\n')
	// if err != nil {
	//     fmt.Println("Error reading input:", err)
	//     return
	// }
	// fmt.Println("You entered:", input)
}