package main

import (
	"fmt"
	"sync"
)

type posts struct {
	views int
	mu sync.Mutex
}

func (p *posts) inc(wg *sync.WaitGroup){
	defer func(){
		wg.Done()
		p.mu.Unlock()
	}()

	p.mu.Lock()
	p.views+=1
}
func main() {
	newPosts := posts{views: 0}
	var wg sync.WaitGroup
	for i:=0;i<100;i++{
		wg.Add(1)
		go newPosts.inc(&wg)
	}
	wg.Wait()
	fmt.Println("Views:", newPosts.views)

}