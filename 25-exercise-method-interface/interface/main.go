package main

import "fmt"

type vehicle interface {
	License() string
	Name() string
}

type car struct {
	licenceNo string
	brand     string
}

func (c car) License() string {
	return c.licenceNo
}
func (c car) Name() string {
	return c.brand
}

type cube struct {
	edge float64
}

func volume(c cube) float64 {
	return c.edge * c.edge * c.edge
}

func main() {

}

func exercise1() {
	var v vehicle = car{brand: "Ford Mustang", licenceNo: "POW100ZZ"}

	fmt.Println(v.License())
	fmt.Println(v.Name())
}

func exercise2() {

	// 1.
	var empty interface{}
	fmt.Printf("%T\n", empty)

	// 2.
	empty = 5
	fmt.Printf("%T\n", empty)

	// 3.
	empty = 10.10
	fmt.Printf("%T\n", empty)

	// 4.
	empty = []int{1, 2, 3}
	fmt.Printf("%T\n", empty)

	// 5. Type Assertion
	empty = append(empty.([]int), 10)

	// 6.
	fmt.Printf("%v\n", empty)
}

func exercise3() {

	var x interface{}
	x = cube{edge: 5}

	// Type Assertion
	v := volume(x.(cube))

	fmt.Printf("Cube Volume: %v\n", v)
}
