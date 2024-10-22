package main

import "fmt"

func main() {
	// const can not be reassigned and it can be declared outside func
	const name string = "Sairam"

	// grouping const
	const (
		port = 5000
		host = "localhost"
	)
	fmt.Println(name)
	fmt.Println("host:", host, "port:", port)
}