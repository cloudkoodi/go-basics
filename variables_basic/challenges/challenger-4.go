package challenges

import (
	"fmt"
	"strconv"
)

func Challenge4() {
	convertAndOperator1()
	convertAndOperator2()
	convertAndOperator3()
	convertAndOperator4()
	convertAndOperator5()
}

// Write a Go program that converts i to float64 and f to int.
// Print out the type and the value of the variables created after conversion.
func convertAndOperator1() {

	var i = 3
	var f = 3.2

	// int to float64
	f1 := float64(i)
	fmt.Printf("f1's type: %T, f1's value: %f\n", f1, f1)

	// float64 to int
	i1 := int(f)
	fmt.Printf("i1's type: %T, i1's value: %d\n", i1, i1)
}

// Write a Go program that converts:
// 1. i to string (int to string)
// 2. s2 to int (string to int)
// 3. f to string (float64 to string)
// 4. s1 to float64  (string to float64)
// 5. Print the value and the type for each variable created after conversion.
func convertAndOperator2() {
	var i = 3
	var f = 3.2
	var s1, s2 = "3.14", "5"

	// 1. int to string
	s := strconv.Itoa(i)
	fmt.Printf("s Type is %T, s value is %q\n", s, s)

	// 2. string to int
	is, err := strconv.Atoi(s2)
	if err == nil {
		fmt.Printf("i type is %T, i value is %v\n", is, is)
	} else {
		fmt.Println("Can not convert string to int.")
	}

	// 3. float64 to string
	ss1 := fmt.Sprintf("%f", f)
	fmt.Printf("ss1's type: %T, ss1's value: %s\n", ss1, ss1)

	// 4. string to float64
	f1, err1 := strconv.ParseFloat(s1, 64)
	if err1 == nil {
		fmt.Printf("f1's type: %T, f1's value: %v\n", f1, f1)
	} else {
		fmt.Println("Cannot convert string to float64.")
	}
}

// There are some errors in the following Go program.
// Try to identify the errors, change the code and run the program without errors.
func convertAndOperator3() {
	x, y := 4, 5.1

	z := float64(x) * y
	fmt.Println(z)

	a := 5.
	b := 6.2 * a
	fmt.Println(b)
}

// Create a Go program that computes how long does it take for the Sunlight to reach the Earth
// if we know that the distance from the Sun to Earth is 149.6 million km and the speed of light  is 299,792,458 m / s
// The formula used is: time = distance / speed
func convertAndOperator4() {
	const (
		distance = 149_600_000 * 1000 // distance from the Sun to Earth in m
		// (it's allowed to use underscores in numbers for a better readability)

		speed = 299_792_458 // speed of light in m / s
	)

	const time = distance / speed // time in seconds

	fmt.Printf("The Sunlight reaches Earth in %v seconds.\n", time)
}

// Write the correct logical operator (&&, || , !)
// inside the expression so that result1 will be false and result2 will be true.
func convertAndOperator5() {
	x, y := 0.1, 5
	var z float64

	// Write the correct logical operator (&&, || , !)
	// inside the expression so that result1 will be false and result2 will be true.

	// false
	result1 := x < z || int(x) != int(z)
	fmt.Println(result1)

	// true
	result2 := y == 1*5 && int(z) == 0
	fmt.Println(result2)
}
