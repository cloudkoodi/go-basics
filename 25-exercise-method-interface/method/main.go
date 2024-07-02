package main

import "fmt"

type money float64

type book struct {
	price float64
	title string
}

func (m money) print() {
	fmt.Printf("The value is %.2f\n", m)
}

func (m money) printStr() string {
	return fmt.Sprintf("%.2f", m)
}

func (b book) vat() float64 {
	return b.price * 0.09
}

func (b *book) discount() {
	b.price = b.price * 0.9
}

func (b *book) changePrice(p float64) {
	b.price = p
}

func main() {
	exercise1()
	exercise2()
	exercise3()
	exercise4()
}

func exercise1() {

	var salary money = 10829.92817

	salary.print()

}

func exercise2() {

	var salary money = 1.5 * 5.503
	fmt.Println(salary) // => 8.2545

	salary.print() // => 8.25

	s := salary.printStr()
	fmt.Println(s)
}

func exercise3() {

	// book value
	bestBook := book{title: "The Trial", price: 9.9}

	// calling the methods
	vat := bestBook.vat()
	fmt.Printf("Vat:%v\n", vat)

	bestBook.discount()
	fmt.Printf("%#v\n", bestBook)
}

func exercise4() {
	// book value
	bestBook := book{title: "The Trial by Kafka", price: 9.9}

	// changing the price
	bestBook.changePrice(11.99)

	fmt.Printf("Book's price:%.2f\n", bestBook.price) // the price is changed
}
