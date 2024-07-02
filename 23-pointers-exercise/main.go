package main

import "fmt"

func main() {
	exercise1()
	exercise2()
}

func exercise1() {

	x := 10.10

	fmt.Printf("The address of x is %p\n", &x)

	ptr := &x

	fmt.Printf("ptr is of type %T with value %v\n", ptr, ptr)

	fmt.Printf("The address of ptr: %p\n", &ptr)
	fmt.Printf("The value of x through the pointer:%f\n", *ptr)
}

func exercise2() {

	x, y := 10, 2
	ptrx, ptry := &x, &y

	z := *ptrx / *ptry

	fmt.Printf("The value of z is %v\n", z)
}

func swap(a, b *float64) {
	*a, *b = *b, *a
}

func exercise3() {

	x, y := 5.5, 8.8
	swap(&x, &y)

	fmt.Printf("x is %v and y is %v\n", x, y)
}
