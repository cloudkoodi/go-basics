package challenges

import "fmt"

func Challenge3() {
	fmtPackage1()
	fmtPackage2()
	fmtPackage3()
}

// Using fmt.Printf():
// Print each variable using a specific verb for its type
// Print the string value enclosed by double quotes ("Gophers")
// Print each variable using the same verb
// Print the type of y and score

func fmtPackage1() {
	x, y, z := 10, 15.5, "Gophers"
	score := []int{10, 20, 30}

	// 1.
	fmt.Printf("x is %d, y is %f, z is %s\n", x, y, z)
	fmt.Printf("score is %#v\n", score)

	// 2.
	fmt.Printf("z is %q\n", z)

	// 3.
	fmt.Printf("x is %v, y is %v, z is %v, score is %v\n", x, y, z, score)

	// 4.
	fmt.Printf("y type: %T, score type: %T\n", y, score)
}

// Consider the following constant declaration: const x float64 = 1.422349587101
// Write a Go program that prints the value of x with 4 decimal points.
func fmtPackage2() {
	const x float64 = 1.422349587101
	fmt.Printf("x is %.4f\n", x) // => formatting with 4 decimal points
}

// There are some errors in the following Go program.
// Try to identify the errors, change the code and run the program without errors.
func fmtPackage3() {
	shape := "circle"
	radius := 3.2
	const pi float64 = 3.14159
	circumference := 2 * pi * radius

	fmt.Printf("Shape: %q\n", shape)
	fmt.Printf("Circle's circumference with radius %f is %f\n", radius, circumference)
}
