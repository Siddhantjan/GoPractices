package main

import (
	"fmt"
	"sync"
)

func message(wg *sync.WaitGroup) {
	wg.Done()
	fmt.Println("hey from hello")
}
func message1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("hey from hello1")
}
func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	fmt.Println("main started")
	go message(wg)
	go message1(wg)
	fmt.Println("main end")
	wg.Wait()
}
