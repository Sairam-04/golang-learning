package main

import (
	"fmt"
	"time"
)

func main() {
	i := 3
	switch i {
		case 1:
			fmt.Println("one")
		case 2:
			fmt.Println("two")
		case 3:
			fmt.Println("Three")
		default:
			fmt.Println("Other than 3")
	}

	// Multiple Condition Switch
	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("Weekend")
	default:
		fmt.Println("WeekDay")
	}

	// type switch
	whoAmI := func (i interface{}){
		switch t := i.(type){
		case int:
			fmt.Println("It is an Integer")
		case string:
			fmt.Println("It is an String")
		case bool:
			fmt.Println("It is an boolean")
		default:
			fmt.Println("other", t)
		}
	}

	whoAmI("golang")
	whoAmI(1)
	whoAmI(11.11)
}