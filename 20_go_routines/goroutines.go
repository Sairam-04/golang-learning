package main

import (
	"fmt"
	"sync"
)
// Implementing waitGroup for synchronizing the tasks
func task(id int, wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("Running Task :", id)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go task(i, &wg)
	}
	wg.Wait()
}