package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	mt := sync.Mutex{}

	wg.Add(10)
	var a = []int{}

	go ekle(&a, &wg, &mt)
	go ekle(&a, &wg, &mt)
	go ekle(&a, &wg, &mt)
	go ekle(&a, &wg, &mt)
	go ekle(&a, &wg, &mt)
	go ekle(&a, &wg, &mt)
	go ekle(&a, &wg, &mt)
	go ekle(&a, &wg, &mt)
	go ekle(&a, &wg, &mt)
	go ekle(&a, &wg, &mt)

	wg.Wait()

	fmt.Println(a)

}

func ekle(a *[]int, wg *sync.WaitGroup, mt *sync.Mutex) {
	mt.Lock()
	*a = append(*a, 0)
	mt.Unlock()
	wg.Done()
}
