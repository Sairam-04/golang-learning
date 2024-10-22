package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func getLanguages() (string, string, string) {
	return "golang", "javascript", "c"
}

func processIt() func(a int) int{
	return func(a int) int{
		return a+5
	}
}
func process(f func(a int) int) int{
	return f(4)
}
func main() {
	fmt.Println(add(43,89))
	fmt.Println(getLanguages())
	fn := func(a int) int{
		return a + 2
	}
	fmt.Println(process(fn))
	x := processIt()
	fmt.Println(x(4))
}