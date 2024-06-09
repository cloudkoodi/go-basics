package main

import "fmt"

func main() {
	//exercise1()
	//exercise2()
	// exercise3()
	// exercise4()
	// exercise5()
	exercise6()
}

func exercise1() {

	for i := 1; i < 50; i++ {

		if i%7 == 0 {
			fmt.Printf("Number %d is divisible by 7\n", i)
		}
	}
}

func exercise2() {

	for i := 1; i < 50; i++ {

		if i%7 != 0 {
			continue
		}

		fmt.Printf("Number %d is divisible by 7\n", i)
	}
}

func exercise3() {

	count := 0

	for i := 1; i < 50; i++ {

		if i%7 == 0 {
			fmt.Printf("Number %d is divisible by 7\n", i)
			count++
		}

		if count == 3 {
			break
		}
	}
}

func exercise4() {

	for i := 1; i < 500; i++ {

		if i%7 == 0 && i%5 == 0 {
			fmt.Printf("Number %d is divisible by 7 and 5\n", i)
		}

	}
}

func exercise5() {

	birthYear, currentYear := 1990, 2024

	for i := birthYear; i <= currentYear; {
		fmt.Printf("%d ", i)
		i++
	}
	fmt.Println()
}

func exercise6() {

	age := 10
	switch {
	case age < 0 || age > 100:
		fmt.Println("Invalid Age")
	case age <= 18:
		fmt.Println("You are minor!")
	case age == 18:
		fmt.Println("Congratulations! You've just become major!")
	default:
		fmt.Println("You are major!")
	}
}
