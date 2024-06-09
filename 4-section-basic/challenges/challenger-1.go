package challenges

import "fmt"

func Challenge1() {

	declare1()
	declare2()
	declare3()
	declare4()

}

/*
	Coding Exercise #1

	Using the var keyword, declare 4 variables called a, b, c, d of type int, float64, bool and string.
	Using short declaration syntax declare and assign these values to variables x, y and z:
		- 20
		- 15.5
		- "Gopher!"
	Using fmt.Println() print out the values of a, b, c, d, x, y and z.
	Try to run the program without error.
	Do you wonder what Gopher is?  Check it here: https://blog.golang.org/gopher
*/

func declare1() {
	var a int
	var b float64
	var c bool
	var d string

	x := 20
	y := 15.5
	z := "Gopher!"

	fmt.Println(a, b, c, d, x, y, z)
}

/*
	Coding Exercise #2

	Change the code from the previous exercise in the following way:

		1. Declare a, b, c, d using a single var keyword (multiple variable declaration) for better readability.
		2. Declare x, y and z on a single line -> multiple short declarations
		3. Remove the statement that prints out the variables. See the error!
		4. Change the program to run without error using the blank identifier (_)
*/

func declare2() {

	var (
		a int
		b float64
		c bool
		d string
	)

	x, y, z := 20, 15.5, "Gopher!"

	_, _, _, _, _, _, _ = a, b, c, d, x, y, z //using blank identifier to mute the unused variable error

}

/*
	Coding Exercise #3

	There are some errors in the following Go program. Try to identify the errors, change the code and run the program without errors.

	----------------------------------------
	|func main() {							\
    |	var a float64 = 7.1					\
    |	x, y := true, 3.7					\
    |	a, x := 5.5, false					\
    |	_, _, _ = a, x, y					\
	|}										\
	-----------------------------------------
*/

func declare3() {

	var a float64 = 7.1

	x, y := true, 3.7

	// ERROR -> no new variables on left side of :=
	//a, x := 5.5, false

	_, _, _ = a, x, y
}

/*
	Coding Exercise #4

	There are some errors in the following Go program. Try to identify the errors, change the code and run the program without errors.

	package main

	version := "3.1"

	func main() {
		name := 'Golang'
		fmt.Println(name)
	}
*/

// ERROR -> := is not allowed in package scope (only local scope)
// version := "3.1"

var version = "3.1"

func declare4() {
	// ERROR -> A string is initialized only using double quotes ("")
	// name = 'Golang'

	name := "Golang"
	fmt.Println(name)
}
