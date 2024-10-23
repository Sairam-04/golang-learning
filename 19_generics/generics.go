package main
import "fmt"

//  generics on function parameters 
// int | string | bool to make this use any datatypes then use Comparable
func printSlice[T comparable](items []T){
	for _, val := range items{
		fmt.Println(val)
	}
}

// generics on struct
type stack[T comparable] struct{
	elements []T
}


func main(){
	nums := []int{1,2,3}
	printSlice(nums)
	extensions := []string{".go", ".js", ".ts", ".java"}
	printSlice(extensions)

	myStack := stack[int]{
		elements: []int{1,2,3},
	}
	myStackNames := stack[string]{
		elements: []string{"golang", "javascript"},
	}
	fmt.Println(myStack, myStackNames)
}