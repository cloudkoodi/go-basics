package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

}

func cube(n float64) float64 {
	return n * n * n
}

func f1(n uint) (int, int) {

	fact := 1
	sum := 0

	for i := 1; i <= int(n); i++ {
		fact *= i
		sum += i
	}
	return fact, sum

}

func myFunc(n string) int {

	// convert string to int
	i, err1 := strconv.Atoi(n)

	// error handling
	if err1 != nil {
		fmt.Printf("Cannot convert %q to int.", n)
		os.Exit(1) //exit the program
	}
	ii, _ := strconv.Atoi(n + n)
	iii, _ := strconv.Atoi(n + n + n)

	return i + ii + iii

}

func searchItem(mySlice []string, myStr string) bool {
	for _, v := range mySlice {
		if v == myStr {
			return true
		}

	}
	return false
}

func searchItem2(mySlice []string, myStr string) bool {
	for _, v := range mySlice {
		if strings.EqualFold(v, myStr) {
			return true
		}

	}
	return false
}

func sum(a ...int) int {
	s := 0
	for _, v := range a {
		s += v
	}
	return s
}

func sum2(a ...int) (s int) {
	for _, v := range a {
		s += v
	}
	return
}

func print(msg string) {
	fmt.Println(msg)
}

func f2(a int) {
	fmt.Println(a)
}

func exercise1() {
	x := cube(3)
	fmt.Println(x)
}

func exercise2() {

	f, s := f1(4)
	fmt.Println("f:", f, "s:", s)

	f, s = f1(10)
	fmt.Println("f:", f, "s:", s)
}

func exercise3() {

	fmt.Println(myFunc("5")) // => 615
	fmt.Println(myFunc("3")) // => 369
}

func exercise4() {

	s := sum(1, 2, 30)
	fmt.Println(s)
}

func exercise5() {
	s := sum2(1, 2, 30)
	fmt.Println(s)
}

func exercise6() {

	animals := []string{"lion", "tiger", "bear"}

	result := searchItem(animals, "bear")
	fmt.Println(result) // => true

	result = searchItem(animals, "pig")
	fmt.Println(result) // => false
}

func exercise7() {

	animals := []string{"lion", "tiger", "bear"}

	result := searchItem(animals, "BEAR")
	fmt.Println(result) // => true

	result = searchItem(animals, "lION")
	fmt.Println(result) // => true

	result = searchItem(animals, "pig")
	fmt.Println(result) // => false
}

func exercise8() {

	defer print("The Go gopher is the iconic mascot of the Go project.")
	fmt.Println("Hello, Go playground!")
}

func exercise9() {
	x := f2
	fmt.Printf("%T\n", x) // => func(int)
	x(8)
}
