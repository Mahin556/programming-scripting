package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Create a new reader that reads from standard input
	reader := bufio.NewReader(os.Stdin)
	// Creates a buffered reader that reads from os.Stdin – the standard input.
	// Buffered reading is efficient and offers methods like ReadString, ReadBytes, ReadLine etc.

	// Prompt the user
	fmt.Println("Enter the rating for our pizza:")

	// Read input until a newline (\n) is encountered
	// The read returns two values: the input string and an error (or nil)
	// _ , err := reader.ReadString('\n')
	input, _ := reader.ReadString('\n')
	// Reads from the input until the first occurrence of the delimiter (\n here).
	// Returns the read string (including the delimiter) and an error.
	// Common delimiters: '\n' (newline), '\r' (carriage return), or custom characters.

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

	// ---- Second prompt: age ----
	fmt.Print("Enter your age: ")

	ageInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading age:", err)
		return
	}

	ageStr := strings.TrimSpace(ageInput) // clean way to remove newline

	fmt.Printf("You entered: %s (type: %T)\n", ageStr, ageStr)

	// ---- Next step (preview) ----
	// To convert ageStr to an integer:
	// age, err := strconv.Atoi(ageStr)
	// if err != nil { ... }
	// fmt.Println("Your age next year:", age+1)
}