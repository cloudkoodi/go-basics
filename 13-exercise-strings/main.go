package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	exercise2()
}

func exercise1() {
	var name string = "Gabriel"

	country := "Brazil"

	fmt.Printf("Your Name: %s\n", name)
	fmt.Printf("Your Country: %s\n", country)

	fmt.Println("He say: \"Hello!\"")
	fmt.Println(`C:\Users\a.txt`)
}

func exercise2() {

	r := "ă"

	fmt.Printf("Type: %T\n", r)

	s1, s2 := "ma", "m"

	s3 := strings.Join([]string{s1, s2, string(r)}, "")

	fmt.Printf("s3 is %s\n", s3)

}

func exercise3() {
	s1 := "țară means country in Romanian"

	// iterating over the string and print out byte by byte
	fmt.Printf("Bytes in string: ")
	for i := 0; i < len(s1); i++ {

		fmt.Printf("%v ", s1[i])
	}

	fmt.Println()

	// iterating over the string and print out rune by rune
	// and the byte index where the rune starts in the string
	for i, r := range s1 {
		fmt.Printf("byte index: %d -> rune: %c\n", i, r)
	}

	fmt.Println()
}

func exercise4() {
	s1 := "Go is cool!"

	x := s1[0]
	fmt.Println(x)

	// ERROR -> cannot assign to s1[0]
	// s1[0] = 'x'

	// printing the number of runes in the string
	// fmt.Println(len(s1))
	fmt.Println(utf8.RuneCountInString(s1))
}

func exercise5() {
	s := "你好 Go!"

	// converting string to byte slice
	b := []byte(s)

	// printing out the byte slice
	fmt.Printf("%#v\n", b)

	// iterating over the byte slice and printing out each index and byte in the byte slice
	for i, v := range b {
		fmt.Printf("%d -> %d\n", i, v)
	}
}

func exercise6() {
	s := "你好 Go!"

	// converting string to rune slice
	r := []rune(s)

	// printing out the rune slice
	fmt.Printf("%#v\n", r)

	// iterating over the rune slice and printing out each index and rune in the rune slice
	for i, v := range r {
		fmt.Printf("%d -> %c\n", i, v)
	}
}
