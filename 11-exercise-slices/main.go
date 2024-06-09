package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

}

func exercise1() {

	countries := []string{"Romania", "Brazil", "Germany"}
	for i, v := range countries {
		fmt.Printf("Index: %d, Element: %q\n", i, v)
	}

}

func exercise2() {

	mySlice := []float64{1.2, 5.6}

	// ERROR -> index out of range [2] with length 2
	// mySlice[2] = 6
	mySlice[1] = 6

	// ERROR -> cannot use a (type int) as type float64 in assignment
	// a := 10
	a := 10.
	mySlice[0] = a

	// ERROR -> index out of range [3] with length 2
	// mySlice[3] = 10.10
	mySlice[1] = 10.10

	mySlice = append(mySlice, a)

	fmt.Println(mySlice)
}

func exercise3() {

	// 1. Declare a slice called nums with 3 float64 numbers.
	nums := []float64{1.1, 2.2, 3.3}

	// 2. Append the value 10.1 to the slice
	nums = append(nums, 10.1)

	// 3. In one statement append to the slice the values: 4.1,  5.5 and 6.6
	nums = append(nums, 4.1, 5.5, 6.6)

	// 4. Print out the slice
	fmt.Println(nums)

	// 5. Declare a slice called n with 2 float64 values
	n := []float64{10.10, 20.20}

	// 6. Append n to nums
	nums = append(nums, n...)

	// 7. Print out the slice
	fmt.Println(nums)
}

func exercise4() {

	if len(os.Args) < 2 { //if not run with arguments
		log.Fatal("Please run the program with arguments (2-10 numbers)!")

	}

	//taking the arguments in a new slice
	args := os.Args[1:]

	// declaring and initializing sum and product of type float64
	sum, product := 0., 1.

	if len(args) < 2 || len(args) > 10 {
		fmt.Println("Please enter between 2 and 10 numbers!")
	} else {

		for _, v := range args {
			num, err := strconv.ParseFloat(v, 64)
			if err != nil {
				continue //if it can't convert string to float64, continue with the next number
			}
			sum += num
			product *= num
		}
	}

	fmt.Printf("Sum: %v, Product: %v\n", sum, product)
}

func exercise5() {
	nums := []int{5, -1, 9, 10, 1100, 6, -1, 6}
	sum := 0
	for _, v := range nums[2 : len(nums)-2] {
		fmt.Println(v)
		sum += v
	}
	fmt.Println("Sum:", sum)
}

func exercise6() {

	friends := []string{"Marry", "John", "Paul", "Diana"}
	yourFriends := make([]string, len(friends))
	copy(yourFriends, friends)

	yourFriends[0] = "Dan"

	fmt.Println(friends, yourFriends)

}

func exercise7() {
	friends := []string{"Marry", "John", "Paul", "Diana"}
	yourFriends := []string{}

	yourFriends = append(yourFriends, friends...)

	yourFriends[0] = "Dan"

	fmt.Println(friends, yourFriends)

}

func exercise8() {

	years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}
	newYears := []int{}

	newYears = append(years[:3], years[len(years)-3:]...)

	fmt.Printf("%#v\n", newYears)
}
