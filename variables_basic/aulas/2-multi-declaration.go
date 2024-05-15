package aulas

import "fmt"

func MultiDeclaration() {

	car, cost := "Audi", 50000
	fmt.Println(car, cost)

	// car, cost := "BMW", 6000 Exception
	car, year := "BMW", 2020 // If you introduce another variable is totally fine

	_ = year

	var opened = false

	opened, file := true, "test.txt"

	_, _ = opened, file

	var (
		salary    float64
		firstName string
		gender    bool
	)

	fmt.Println(salary, firstName, gender)

	var a, b, c int
	fmt.Println(a, b, c)

	var i, j int
	i, j = 5, 8

	// _,_ = i,j

	j, i = i, j // swap variables

	fmt.Println(i, j)

	sum := 5 + 2.3

	fmt.Println(sum)

}
