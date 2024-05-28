package challenges

import "fmt"

func Challenge2() {
	constant1()
	constant2()
	constant3()
	constant4()
	constant5()
	constant6()
}

// Using the const keyword declare and initialize the following constants:
// 1. daysWeek with value 7
// 2. lightSpeed with value 299792458
// 3. pi with value 3.14159
// Run the program without errors.

func constant1() {

	const daysWeek int8 = 7
	const lightSpeed int32 = 299792458
	const pi float64 = 3.14159
}

// Change the code from the previous exercise and declare all 3 constants as grouped constants.
// Make them untyped.
func constant2() {

	const (
		daysWeek   = 7
		lightSpeed = 299792458
		pi         = 3.14159
	)
}

func constant3() {

	const (
		secPerDay   = 86400
		daysPerYear = 365
	)

	fmt.Printf("There are %v seconds in a year.", secPerDay*daysPerYear)
}

// There are an error in the following Go program.
// Try to identify the error, change the code and run the program without errors.
func constant4() {
	const x int = 10

	//** There are ONLY boolean constants, rune constants, integer constants, **//
	//** floating-point constants, complex constants, and string constants. **//

	// declaring a constant of type slice int ([]int)
	// ERROR -> const initializer []int literal is not a constant
	// const m = []int{1: 3, 4: 5, 6: 8}	-> You cannot declare a slice constant.
	// _ = m
}

// There are an error in the following Go program.
// Try to identify the error, change the code and run the program without errors.
func constant5() {

	// ERROR -> invalid operation: a * b (mismatched types int and float64)
	// const a int = 7
	const a = 7 // make `a` untyped constant
	const b float64 = 5.6
	const c = a * b

	x := 8
	_ = x
	// ERRROR ->  You cannot initiate a constant at runtime (constants belong to compile-time)
	// const xc int = x  // variables belong to runtime

	// ERRROR ->  You cannot initiate a constant at runtime (constants belong to compile-time)
	// const noIPv6 = math.Pow(2, 128) // functions calls belong to runtime

}

// Using Iota declare the following months of the year: Jun, Jul and Aug
// Jun, Jul and Aug are constant and their value is 6, 7 and 8.
func constant6() {

	const (
		//iota starts from zero
		Jun = iota + 6
		Jul //automatically incremented by one
		Aug
	)

	fmt.Println(Jun, Jul, Aug)
}
