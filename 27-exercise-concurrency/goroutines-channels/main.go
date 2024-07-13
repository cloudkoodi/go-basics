package main

import "fmt"

func main() {

}

func exercise1() {
	// 1.
	var c1 chan float64

	// 2.
	// Declaring and initilizing a RECEIVE-ONLY channel
	c2 := make(<-chan rune)

	// Declaring and initilizing a SEND-ONLY channel
	c3 := make(chan<- rune)

	// 3.
	c4 := make(chan int, 10)

	// 4.
	fmt.Printf("%T, %T, %T, %T\n", c1, c2, c3, c4)
}

func exercise2() {
	ch := make(chan string)

	go func(n string) {
		ch <- n
	}("Gophers!")

	value := <-ch
	fmt.Println("Value received from channel:", value)
}

func exercise3() {
	c := make(chan int)

	go func(n int) {
		c <- n
	}(100)

	fmt.Println(<-c)
}

func power(x int, c chan int) {
	c <- x * x
}

func exercise4() {
	ch := make(chan int)

	for i := 1; i <= 50; i++ {
		go power(i, ch)
	}

	for i := 1; i <= 50; i++ {
		fmt.Println(<-ch)
	}
}

func exercise5() {
	ch := make(chan int)

	for i := 1; i <= 100; i++ {
		go func(x int) {
			ch <- x * x
		}(i)
	}

	for i := 1; i <= 100; i++ {
		fmt.Println(<-ch)
	}
}
