package main

import "fmt"

type person struct {
	name   string
	age    int
	colors []string
}

type person2 struct {
	name   string
	age    int
	colors []string
	gr     grades
}

type grades struct {
	grade  int
	course string
}

func main() {

}

func exercise1() {
	// 2.
	me := person{name: "Joao", age: 30, colors: []string{"red", "yellow"}}
	you := person{name: "Maria", age: 22, colors: []string{"pink", "blue"}}

	// 3.
	fmt.Printf("%v\n", me)
	fmt.Printf("%+v\n", you)
}

func exercise2() {

	me := person{name: "Joao", age: 30, colors: []string{"red", "yellow"}}
	you := person{name: "Maria", age: 22, colors: []string{"pink", "blue"}}

	// 1.
	me.name = "Gabriel"

	// 2.
	var colors []string = you.colors
	fmt.Printf("Type: %T, Value: %v\n", colors, colors)

	// 3.
	colors = append(colors, "green")
	you.colors = colors

	// 4.
	fmt.Printf("%v\n", me)
	fmt.Printf("%+v\n", you)

}

func exercise3() {
	me := person{name: "Joao", age: 30, colors: []string{"red", "yellow"}}
	you := person{name: "Maria", age: 22, colors: []string{"pink", "blue"}}

	for index, color := range me.colors {
		fmt.Printf("%d -> %q\n", index, color)
	}

	_ = you

}

func exercise4() {

	me := person2{
		name:   "Joao",
		age:    30,
		colors: []string{"red", "yellow"},
		gr:     grades{grade: 85, course: "Python"},
	}
	you := person2{
		name:   "Maria",
		age:    22,
		colors: []string{"pink", "blue"},
		gr:     grades{grade: 100, course: "History"},
	}

	// 4.
	me.gr.grade = 98
	me.gr.course = "Golang"

	// 5.
	fmt.Printf("%v\n", me)
	fmt.Printf("%+v\n", you)
}
