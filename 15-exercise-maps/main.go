package main

import "fmt"

func main() {

}

func exercise1() {

	var m1 map[float64]bool
	fmt.Printf("m1 type: %T, m1 value: %#v\n", m1, m1)

	// 2.
	m2 := map[int]string{1: "Sting", 2: "Queen"}

	// 3.
	m2[10] = "Abba"

	// 4
	fmt.Println(m2[2])   // existing key
	fmt.Println(m2[100]) // non-existing key
}

func exercise2() {
	var m1 map[int]bool

	// ERROR -> panic: assignment to entry in nil map
	// m1[5] = true

	m2 := map[int]int{3: 10, 4: 40}
	m3 := map[int]int{3: 10, 4: 40}

	// ERROR -> invalid operation: m2 == m3 (map can only be compared to nil)
	// fmt.Println(m2 == m3)

	_, _, _ = m1, m2, m3
}

func exercise3() {
	// 1.
	m := map[int]bool{1: true, 2: false, 3: false}

	// 2.
	delete(m, 2)

	// 3.
	for k, v := range m {
		fmt.Printf("k: %d, v: %t\n", k, v)
	}
}
