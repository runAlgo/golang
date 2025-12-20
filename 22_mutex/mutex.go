package main

import (
	"fmt"
	"sync"
)

type Post struct {
	views int
	mu sync.Mutex
}

func (p *Post) inc(w *sync.WaitGroup) {
	defer func(){
		p.mu.Unlock()
		w.Done()
	}()

	p.mu.Lock()
	p.views += 1
}

func main() {
	var wg sync.WaitGroup
	myPost := Post{
		views: 0,
	}
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg)
	}
	wg.Wait()
	fmt.Println("Total views:",myPost.views)

}