package main

import "fmt"

func main() {
	age := 18
	if age >= 18 {
		fmt.Println("Person is an adult")
	}else if age >= 12{
		fmt.Println("Person is an teenager")
	}else{
		fmt.Println("Person is not an adult")
	}

	//we can declare variable inside if construct
	// There is no terenary opertor in go
	if age := 15; age >=18{
		fmt.Println("Person is an adult")
	}else if age>=12{
		fmt.Println("Person is an teenager")
	}
}