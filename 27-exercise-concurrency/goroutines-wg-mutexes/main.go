package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	// exercise1()
	exercise2()
}

func sayHello(n string, wg *sync.WaitGroup) {
	fmt.Printf("Hello, %s!\n", n)

	wg.Done()
}

func exercise1() {

	var wg sync.WaitGroup
	wg.Add(1)

	go sayHello("Gabriel", &wg)

	wg.Wait()
}

func sum(w *sync.WaitGroup, n ...float64) {
	total := 0.0

	for _, num := range n {
		total += num
	}

	fmt.Printf("The total is: %.2f\n", total)
	w.Done()
}

func exercise2() {

	var wg sync.WaitGroup

	wg.Add(3)

	go sum(&wg, 4.1, 4.6)
	go sum(&wg, 5.1, 4.6)
	go sum(&wg, 6.1, 4.6)

	wg.Wait()

}

func exercise3() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func(f float64, wg *sync.WaitGroup) {
		x := math.Sqrt(f)
		fmt.Printf("Square root of %.2f is %.4f\n", f, x)
		wg.Done()
	}(16.1, &wg)

	wg.Wait()
}

func exercise4() {
	var wg sync.WaitGroup

	wg.Add(50)

	for i := 100.; i < 150.; i++ {
		go func(f float64, wg *sync.WaitGroup) {
			x := math.Sqrt(f)
			fmt.Printf("Square root of %.2f is %.6f\n", f, x)
			wg.Done()
		}(i, &wg)
	}

	wg.Wait()
}

func deposit(b *int, n int, wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	*b += n
	m.Unlock()
	wg.Done()
}

func withdraw(b *int, n int, wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	*b -= n
	m.Unlock()
	wg.Done()
}

func exercise5() {
	var wg sync.WaitGroup

	var m sync.Mutex

	wg.Add(200)

	balance := 100

	for i := 0; i < 100; i++ {
		go deposit(&balance, i, &wg, &m)
		go withdraw(&balance, i, &wg, &m)
	}
	wg.Wait()

	// the final balance value be always 100
	fmt.Println("Final balance value:", balance)
}
