package main

import "fmt"

func main() {

	// exercise1()
	// exercise2()
	exercise3()
}

func exercise1() {

	// 1.
	var cities [2]string
	fmt.Printf("%#v\n", cities)

	// 2.
	grades := [3]float64{4.5, 9.7}
	fmt.Printf("%#v\n", grades)

	// 3.
	taskDone := [...]bool{true, false, false, true}
	fmt.Printf("%#v\n", taskDone)

	// 4.
	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}

	// 5.
	for index, value := range grades {
		fmt.Println(index, value)
	}

}

func exercise2() {

	nums := []int{30, -1, -6, 90, -6}

	var count int

	for _, v := range nums {
		if v%2 == 0 && v > 0 {
			count++
		}
	}

	fmt.Println("No. of positive even numbers in nums: ", count)
}

func exercise3() {

	myArray := [3]float64{1.2, 5.6}

	myArray[2] = 6

	a := 10
	// ERROR -> invalid array index 3 (out of bounds for 3-element array)
	// myArray[0] = a
	myArray[0] = float64(a)

	// ERROR -> invalid array index 3 (out of bounds for 3-element array)
	// myArray[3] = 10.10

	myArray[2] = 10.10

	fmt.Println(myArray)

}
