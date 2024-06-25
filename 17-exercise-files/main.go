package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	exercise1()
}

func exercise1() {

	file, err := os.OpenFile(
		"info.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0644,
	)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
}

func exercise2() {
	// checking if file exists
	_, err := os.Stat("info.txt")
	// error handling
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("The file does not exist!")
		}
	}

	// renaming the file
	err = os.Rename("info.txt", "data.txt")
	// error handling
	if err != nil {
		log.Fatal(err)
	}
}

func exercise3() {
	// removing the file
	err := os.Remove("data.txt")
	// error handling
	if err != nil {
		log.Fatal(err)
	}
}

func exercise4() {
	// opening the file in read-only mode. The file must exist (in the current working directory)
	// use a valid path!
	file, err := os.Open("info.txt")
	if err != nil {
		log.Fatal(err)
	}

	// defer closing the file
	defer file.Close()

	// ioutil.ReadAll() reads from the file until an error or EOF and returns the data it read
	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Data as string: %s\n", data)
}

func exercise5() {
	// opening the file in read-only mode. The file must exist (in the current working directory)
	// use a valid path!
	file, err := os.Open("info.txt")
	// error handling
	if err != nil {
		log.Fatal(err)
	}
	// defer closing the file
	defer file.Close()

	// the file value returned by os.Open() is wrapped in a bufio.Scanner just like a buffered reader.
	scanner := bufio.NewScanner(file)

	// reading the whole file line by line:
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// checking for any possible errors:
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func exercise6() {
	// ioutil.WriteFile() handles creating, opening, writing a slice of bytes and closing the file.
	// if the file doesn't exist WriteFile() creates it
	// and if it already exists the function will truncate it before writing to file.

	bs := []byte("The Go gopher is an iconic mascot!")
	err := os.WriteFile("info.txt", bs, 0644)
	// error handling
	if err != nil {
		log.Fatal(err)
	}
}
