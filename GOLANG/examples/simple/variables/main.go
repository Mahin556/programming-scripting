package main

import "fmt"

var outsideVar string = "I am an outside variable"

// notworkoutsideVar := "I am an outside variable" // This will cause a compilation error because short variable declarations cannot be used outside of functions

const constantVar string = "I am a constant variable"

var PublicVar string = "I am a public variable" // Public variable, can be accessed from other packages

//Go doesn't allow unused local variables

func main() {
	var name string = "John" //string
	fmt.Println(name)
	fmt.Printf("Variable of type %T and value %v\n", name, name)
	
	var age int = 30 //integer
	fmt.Println(age)
	fmt.Printf("Variable of type %T and value %v\n", age, age)
	
	var height float64 = 5.9 //float64
	fmt.Println(height)
	fmt.Printf("Variable of type %T and value %v\n", height, height)

	var isStudent bool = true //boolean
	fmt.Println(isStudent)
	fmt.Printf("Variable of type %T and value %v\n", isStudent, isStudent)

	var demo = "string" //lexer assign a type to the variable based on the value assigned
	fmt.Println(demo)
	fmt.Printf("Variable of type %T and value %v\n", demo, demo)

	short_var := "short variable declaration" //short variable declaration
	fmt.Println(short_var)
	fmt.Printf("Variable of type %T and value %v\n", short_var, short_var)

	fmt.Println(outsideVar)
	fmt.Printf("Variable of type %T and value %v\n", outsideVar, outsideVar)

	fmt.Println(constantVar)
	fmt.Printf("Variable of type %T and value %v\n", constantVar, constantVar)

	fmt.Println(PublicVar)
	fmt.Printf("Variable of type %T and value %v\n", PublicVar, PublicVar)

	// The instructor recommends naming Boolean variables with an is prefix because it makes their meaning obvious:
	//
	// isLoggedIn
	// isVerified
	// isActive
	// isAdmin

	// Signed integers can represent negative and positive values.
	var signedInt int = -42
	fmt.Println(signedInt)
	fmt.Printf("Variable of type %T and value %v\n", signedInt, signedInt)

	// Unsigned integers can only represent positive values.
	var unsignedInt uint = 42
	// var unsignedInt uint = -42 // This will cause a compilation error because unsigned integers cannot be negative
	fmt.Println(unsignedInt)
	fmt.Printf("Variable of type %T and value %v\n", unsignedInt, unsignedInt)

	//variable wthout initialization will be assigned a zero value based on its type
	var uninitializedInt int
	fmt.Println(uninitializedInt) // Output: 0
	fmt.Printf("Variable of type %T and value %v\n", uninitializedInt, uninitializedInt)

	//Unlike languages where an uninitialized variable might contain unpredictable/garbage data, Go gives variables their appropriate zero value.
	// | Type      | Zero value |
	// | --------- | ---------- |
	// | `int`     | `0`        |
	// | `uint`    | `0`        |
	// | `float32` | `0`        |
	// | `float64` | `0`        |
	// | `bool`    | `false`    |
	// | `string`  | `""`       |
	// | pointer   | `nil`      |
	// | slice     | `nil`      |
	// | map       | `nil`      |
	// | interface | `nil`      |


	//Reassignment of variables
	name = "Doe"
	fmt.Println(name)
	fmt.Printf("Variable of type %T and value %v\n", name, name)

	short_var = "short variable declaration reassigned"
	fmt.Println(short_var)
	fmt.Printf("Variable of type %T and value %v\n", short_var, short_var)

	//println vs printf
	fmt.Println("Name:", name, "Age:", age) //It automatically adds a space between multiple arguments and a newline at the end:

	//Use Printf when you want formatted output using placeholders.
	fmt.Printf("Name: %s, Age: %d\n", name, age)
}