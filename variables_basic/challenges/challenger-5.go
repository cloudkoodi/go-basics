package challenges

import "fmt"

func Challenge5() {
	alias1()
	alias2()
	alias3()
}

// Declare a new type type called duration. Have the underlying type be an int.
// Declare a variable of the new type called hour using the var keyword
// In function main:
// print out the value of the variable hour
// print out the type of the variable hour
// assign 3600 to the variable hour using the  = operator
// print out the value of hour

type duration int

func alias1() {

	var hour duration
	fmt.Printf("hour's type: %T, hour's value: %v\n", hour, hour)
	hour = 3600
	fmt.Printf("hour's value %v\n", hour)
}

// There are some errors in the following Go program.
// Try to identify the errors, change the code and run the program without errors.

func alias2() {
	var hour duration = 3600
	minute := 60
	// different names are different types
	// we convert minute which is of type int to type duration
	fmt.Println(hour != duration(minute))
}

// Declare two defined types called mile and kilometer. Have the underlying type be an float64.
// Declare a constant called m2km equals 1.609  ( 1mile=1.609km)
// In function main:
// declare a variable called mileBerlinToParis of type mile with value 655.3
// declare a variable called kmBerlinToParis of type kilometer
// calculate the distance between Berlin and Paris in km by multiplying mileBerlinToParis and m2km. Assign the result to kmBerlinToParis
// print out the distance in km between Berlin and Paris

type mile float64
type kilometer float64

const m2km = 1.609

func alias3() {

	var mileBerlinToParis mile = 655.3 //distance in miles
	var kmBerlinToParis kilometer

	kmBerlinToParis = kilometer(mileBerlinToParis * m2km)
	fmt.Printf("Distance in Km from Berlin to Paris is %f\n", kmBerlinToParis)
}
